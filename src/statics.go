/*  A personal web server intended for hosting dwbrite.com
 *  Copyright (C) 2020 Devin Brite
 */

package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func (s *server) serveStatic(internalPath string, metadata *metadata) {
	pathDirectory, reqPath := getPaths(internalPath)

	content, _ := ioutil.ReadFile(filepath.Join(pathDirectory, "index.html"))

	data := metadata.Data.(map[string]interface{})

	prototype := struct {
		Title         string
		StaticContent template.HTML
	}{
		data["title"].(string),
		template.HTML(string(content)),
	}

	cache := cache{
		new(bytes.Buffer),
	}

	template.Must(template.New("base").
		Funcs(template.FuncMap{"fieldExists": fieldExists}).
		ParseFiles("tmpl/static.gohtml", "tmpl/base.gohtml")).
		Execute(cache, prototype)

	s.HandleFunc(reqPath, cache.serve)
}

func (s *server) serveResources(internalPath string, m *metadata) {
	pathDirectory, reqPath := getPaths(internalPath)

	fs := http.FileServer(http.Dir(pathDirectory))
	s.PathPrefix(reqPath + "/").Handler(http.StripPrefix(reqPath, fs))
	// force / at the end of resource directories - permanent redirect
	s.Handle(reqPath, http.RedirectHandler(reqPath+"/", 308))
}
