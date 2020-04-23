package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ProjectData struct {
	Title   string
	Summary template.HTML
	Body    template.HTML
	Path    string
}

type FolioRegister struct {
	s        *server
	Projects []*ProjectData
	Title    string
}

func (s *server) servePortfolio(internalPath string, metadata *metadata) {
	pathDirectory, reqPath := getPaths(internalPath)

	p := FolioRegister{
		s:        s,
		Projects: []*ProjectData{},
		Title:    "portfolio",
	}

	filepath.Walk(pathDirectory, p.walkPortfolio)

	//content, _ := ioutil.ReadFile(filepath.Join(dir, "index.html"))
	//data := metadata.Data.(map[string]idnterface{})

	cache := cache{
		new(bytes.Buffer),
	}

	template.Must(template.New("base").
		Funcs(template.FuncMap{"fieldExists": fieldExists}).
		ParseFiles("tmpl/portfolio.gohtml", "tmpl/base.gohtml")).
		Execute(cache, p)

	s.HandleFunc(reqPath, cache.serve)
}

func (p *FolioRegister) walkPortfolio(internalPath string, file os.FileInfo, err error) error {
	pathDirectory, reqPath := getPaths(internalPath)

	// early failure clauses
	if "metadata.json" != file.Name() {
		return nil
	}
	metadata := p.s.readMetadata(internalPath)
	if metadata.Archetype != "project" {
		return nil
	}

	// define project data
	data := metadata.Data.(map[string]interface{})

	projectBody, _ := ioutil.ReadFile(filepath.Join(pathDirectory, "index.html"))

	project := ProjectData{
		Title:   data["title"].(string),
		Summary: template.HTML(data["summary"].(string)),
		Path:    reqPath,
		Body:    template.HTML(projectBody),
	}

	// add project to register
	p.Projects = append(p.Projects, &project)

	cache := cache{
		new(bytes.Buffer),
	}

	_ = template.Must(template.New("base").
		Funcs(template.FuncMap{"formatDate": formatDate, "fieldExists": fieldExists}).
		ParseFiles("tmpl/project.gohtml", "tmpl/base.gohtml")).
		Execute(cache, project)

	p.s.HandleFunc(reqPath, cache.serve)

	return nil
}
