/*  A personal web server intended for hosting dwbrite.com
 *  Copyright (C) 2018 Devin Brite
 *  Licensing information can be found in `README.md` and `LICENSE.md`
 */

package main

import (
	_ "github.com/lib/pq"
	"regexp"
	"html/template"
	"net/http"
	"strconv"
	"database/sql"
	"log"
)

const (
	DB_INFO = "user=devin dbname=dwbrite_com sslmode=disable"
	CERT = "certs/dwbrite.com.cert"
	KEY = "certs/dwbrite.com.key"
	PORT = ":41234"
)

var pageTemplate = template.Must(template.ParseFiles("tmpl/page.html", "tmpl/base.html"))
var fileWhitelist = regexp.MustCompile("^/.*\\.(html|css|js|png|jpg|gif|webm)$")

// TODO: Create page-something interface.

type page struct {
	Title string
	Body  template.HTML
}

/*type post struct {
	Uid  int32
	PostDate time.Time
	Content page
}

type project struct {
	Summary string
	Content page
}*/

func main() {
	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/certs/", forbidRequest)
	http.HandleFunc("/tmpl/", forbidRequest)

	log.Fatal(http.ListenAndServeTLS(PORT, CERT, KEY, nil))
}

func renderPage(w http.ResponseWriter, title string) {
	p, err := selectPage(title)
	if checkError(w, err) {
		println(err)
		return
	}

	err = pageTemplate.ExecuteTemplate(w, "base", p)
	if checkError(w, err) { return }
}

func checkError(w http.ResponseWriter, err error) bool {
	if err != nil { http.Error(w, err.Error(), http.StatusInternalServerError) }
	return err != nil
}

func handleRequest(writer http.ResponseWriter, request *http.Request) {
	//TODO: switch/case here.
	//TODO: "/blog/", "/blog/article/", "/work/"
	if request.URL.Path == "/" {
		renderPage(writer, "About Me")
	} else {
		serveFile(writer, request)
	}
}

func forbidRequest(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusForbidden)
	renderPage(writer, strconv.Itoa(http.StatusForbidden))
}

func serveFile(writer http.ResponseWriter, request *http.Request) {
	if fileWhitelist.MatchString(request.URL.Path) {
		http.ServeFile(writer, request, request.URL.Path[1:])
	} else {
		forbidRequest(writer, request)
	}
}

func selectPage(title string) (*page, error) {
	db, _ := sql.Open("postgres", DB_INFO)
	var body template.HTML
	err := db.QueryRow("SELECT * FROM pages WHERE title = $1;", title).Scan(&title, &body)
	if err != nil { return nil, err }
	return &page{Title: title, Body: body}, nil
}