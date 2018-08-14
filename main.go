/*  A personal web server intended for hosting dwbrite.com
 *  Copyright (C) 2018 Devin Brite
 *  Licensing information can be found in `README.md` and `LICENSE.md`
 */

package main
/*
import (
	_ "github.com/lib/pq"
	"regexp"
					)

const (
	DB_INFO = "user=devin dbname=dwbrite_com sslmode=disable"
	CERT = "certs/dwbrite.com.cert"
	KEY = "certs/dwbrite.com.key"
	PORT = ":41234"
	POSTS_PER_PAGE = 5
)

var fileWhitelist = regexp.MustCompile("^/.*\\.(html|css|js|png|jpg|gif|webm)$")
var dirBlacklist = regexp.MustCompile("^/(certs|examples|tmpl)/?.*$")

var rPost = regexp.MustCompile("^/blog/post/?.*?$")
var rBlog = regexp.MustCompile("^/blog/?.*$")
var rPortfolio = regexp.MustCompile("^/work/?.*$")

/*
type project struct {
	Summary string
	Content page
}*/

func main() {
	//http.HandleFunc("/", matchRequest)
	//log.Fatal(http.ListenAndServeTLS(PORT, CERT, KEY, nil))
}

/*

func checkError(w http.ResponseWriter, err error) bool {
	if err != nil { http.Error(w, err.Error(), http.StatusInternalServerError) }
	return err != nil
}

func matchRequest(writer http.ResponseWriter, request *http.Request) {
	//TODO: "/blog/", "/blog/article/", "/work/"
	path := request.URL.Path
	switch {
	case dirBlacklist.MatchString(path):
		serveError(writer, http.StatusForbidden)
	case path == "/":
		servePage(writer, "About Me")
	case rPost.MatchString(path):
		servePost(writer, request)
	case rBlog.MatchString(path):
		serveBlog(writer, request)
	case rPortfolio.MatchString(path):
	default:
		serveFile(writer, request)
	}
}

func serveError(writer http.ResponseWriter, httpStatus int) {
	writer.WriteHeader(httpStatus)
	servePage(writer, strconv.Itoa(httpStatus))
}

func serveFile(writer http.ResponseWriter, request *http.Request) {
	if fileWhitelist.MatchString(request.URL.Path) {
		http.ServeFile(writer, request, request.URL.Path[1:])
	} else {
		serveError(writer, http.StatusForbidden)
	}
}*