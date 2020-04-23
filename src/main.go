/*  A personal web server intended for hosting dwbrite.com
 *  Copyright (C) 2020 Devin Brite
 */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/qri-io/jsonschema"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type metadata struct {
	Archetype string      `json:"archetype"`
	Data      interface{} `json:"data"`
}

type server struct {
	*mux.Router
	schema *jsonschema.RootSchema
}

func initServer() *server {

	schema, err := os.Open("metadata.schema.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Reading json schema")
	// defer the closing of our jsonFile so that we can parse it later on
	defer schema.Close()

	schemaData, err := ioutil.ReadAll(schema)

	rs := &jsonschema.RootSchema{}
	if err := json.Unmarshal(schemaData, rs); err != nil {
		panic("unmarshal schema: " + err.Error())
	}

	return &server{
		mux.NewRouter(),
		rs,
	}
}

func (s *server) readMetadata(filename string) *metadata {
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	if errors, _ := s.schema.ValidateBytes(byteValue); len(errors) > 0 {
		panic(errors)
	}

	metadata := &metadata{}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &metadata)

	return metadata
}

func (s *server) handlePaths(internalPath string, file os.FileInfo, err error) error {
	if "metadata.json" != file.Name() {
		return nil
	}

	metadata := s.readMetadata(internalPath)
	switch metadata.Archetype {
	case "basic":
		s.serveStatic(internalPath, metadata)
	case "resource":
		s.serveResources(internalPath, metadata)
		//case "blog":
		//	s.serveBlog(internalPath, metadata)
		//case "portfolio":
		//	s.servePortfolio(internalPath, metadata)
	}

	return nil
}

type cache struct {
	*bytes.Buffer
}

func (c cache) serve(writer http.ResponseWriter, request *http.Request) {
	writer.Write(c.Bytes())
}

func main() {
	s := initServer()

	_ = filepath.Walk("root", s.handlePaths)

	s.
		Host("localhost").
		Host("dwbrite.com").
		Host("127.0.0.1")

	server := http.Server{
		Addr:    ":41234",
		Handler: s,
	}

	_ = server.ListenAndServe()
}
