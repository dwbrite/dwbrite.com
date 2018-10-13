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
	strconv "strconv"
)

type blogPage struct {
	BlogPosts	[]*PageContents
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
	var limit uint32 = 5
	query := request.URL.Query()
	postUid := query.Get("post")

	blogPosts := blogPage{
		[]*PageContents{},
		"Blog",
		0,
	}

	// TODO: check if pageNum query > number of pages.
	page, err := strconv.Atoi(request.URL.Query().Get("page"))
	if err != nil || page < 1 {
		blogPosts.PageNum = 0
	} else {
		blogPosts.PageNum = uint32(page)
	}
	if postUid != "" {
		post, _ := SelectContentByUid(postUid, "posts", dbInfo)

		blogPosts.BlogPosts = append(blogPosts.BlogPosts, post)
		return blogPosts
	} else {
		posts, err := SelectMultipleContents(limit, limit*blogPosts.PageNum, Post_date, DESC, "posts", dbInfo)

		if err != nil {
			posts, _ = SelectMultipleContents(limit, 0, Post_date, DESC, "posts", dbInfo)
		}

		blogPosts.BlogPosts = append(blogPosts.BlogPosts, posts...)
		return blogPosts
	}
}

/*func postPath(writer http.ResponseWriter) interface{} {
	return writer
}*/
