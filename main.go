/*  A personal web server intended for hosting dwbrite.com
 *  Copyright (C) 2018 Devin Brite
 *  Licensing information can be found in `README.md` and `LICENSE.md`
 */

package main

import (
	. "./teacup"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

type blogPage struct {
	BlogPosts	[]PageContents
	Title		string
	PageNum		uint32
}

func main() {
	file, err := os.Create("dwbrite.com.log")
	if err != nil {
		log.Fatal(err)
	}

	t := NewTeacup(
		41234,
		"user=devin dbname=dwbrite_com sslmode=disable",
		&TlsKeyPair{
			"certs/dwbrite.com.cert",
			"certs/dwbrite.com.key",
		},

		*regexp.MustCompile("^/.*\\.(html|css|js|png|jpg|gif|webm|ico|md|mp3|mp4)$"),
		*regexp.MustCompile("^/(certs|examples|tmpl)/?.*$"),

		log.New(file, "", log.LstdFlags|log.Lshortfile),
		)

	t.CreateTable("posts", false)
	t.CreateTable("projects", true)
	//t.CreateTable("pages", true)

	//"^...?.*?$"
	t.AddDynamicPage("^/blog/?.*?$", "posts", template.FuncMap{"formatDate" : formatDate}, blogQuery, "tmpl/blogpost.gohtml", "tmpl/blogpage.gohtml", "tmpl/base.gohtml")

	t.StartServer()
}

func formatDate(t time.Time) string {
	{ return t.Format("2006-01-02") }
}


func blogQuery(request http.Request, dbInfo string) interface{} {
	query := request.URL.Query()
	postUid := query.Get("post")
	pageNum := query.Get("page")

	blogPosts := blogPage{
		[]PageContents{},
		"Blog",
		0,
	}

	/*pageNum, err := strconv.Atoi(request.URL.Query().Get("page"))
	if err != nil {
		pageNum = -1
	}*/
	if postUid != "" {
		post, _ := SelectContentByUid(postUid, "posts", dbInfo)
		if post != nil {
			blogPosts.BlogPosts = append(blogPosts.BlogPosts, *post)
			return blogPosts
		}
	} /* TODO("Create Query for multiple ")
	else if pageNum == -1 {
		posts, err := t.selectPosts("SELECT * FROM " +
			"(SELECT * FROM posts ORDER BY post_date DESC " +
			"LIMIT 5) AS derivedTable " +
			"ORDER BY post_date DESC;")
	} else {
		posts, err := t.selectPosts("SELECT * FROM posts ORDER BY post_date ASC " +
			"LIMIT " + strconv.Itoa(t.PostsPerPage) +
			"OFFSET " + strconv.Itoa(t.PostsPerPage*(pageNum)) + ";")
	}

	blogPosts.BlogPosts = posts
	err = postsTemplate.ExecuteTemplate(writer, "base", blogPosts)
	if t.checkError(writer, err) {
		return
	}*/

	return nil //TODO(Return blog page 0)
}

func postPath(writer http.ResponseWriter) interface{} {
	return writer
}
