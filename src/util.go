/*  A personal web server intended for hosting dwbrite.com
 *  Copyright (C) 2020 Devin Brite
 */

package main

import (
	"path/filepath"
	"reflect"
	"time"
)

func getPaths(internalPath string) (dir string, reqPath string) {
	// dir of path (path may point to file)
	// e.g. /root/portfolio/metadata.json -> /root/portfolio
	dir = filepath.Dir(internalPath)
	relPath, _ := filepath.Rel("root", dir)
	reqPath = filepath.Clean("/" + relPath)
	return
}

func formatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func fieldExists(name string, obj interface{}) bool {
	s := reflect.TypeOf(obj)
	_, b := s.FieldByName(name)
	return b
}
