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
	"reflect"
	"regexp"
	"strconv"
	"time"
)

type blogPage struct {
	BlogPosts []*PageContents
	Title     string
	PageNum   uint32
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

		*regexp.MustCompile("^/.*\\.(html|css|js|png|jpg|gif|webm|ico|md|mp3|mp4|ttf)$"),
		*regexp.MustCompile("^/(certs|examples|tmpl)/?.*$"),

		log.New(file, "", log.LstdFlags|log.Lshortfile),
	)

	t.CreateTable("posts", false)
	t.CreateTable("projects", true)
	//t.CreateTable("pages", true)

	// error template
	errTmpl := template.Must(template.New("base").
		Funcs(template.FuncMap{"fieldExists": fieldExists}).
		ParseFiles("tmpl/page.gohtml", "tmpl/base.gohtml"))
	t.SetErrorTemplate(errTmpl)

	//"^...?.*?$"
	t.AddDynamicPage("^/blog/?.*?$", "posts", blogQuery)

	t.StartServer()
}

func formatDate(t time.Time) string {
	{
		return t.Format("2006-01-02")
	}
}

func blogQuery(request http.Request, dbInfo string) (*template.Template, interface{}) {
	// If a specific post is requested, return that post.
	//   If no such post exists, return (nil, nil) | (404)
	// Otherwise
	//   If a page is requested, return that page.
	//   If the page doesn't exist, just return the first page.

	var limit uint32 = 5
	blogPosts := blogPage{
		[]*PageContents{},
		"Blog",
		0,
	}

	query := request.URL.Query()
	blogPosts.PageNum = queryToUint32(query.Get("page"), 0)

	if query.Get("post") != "" {
		postUid := queryToUint32(query.Get("post"), 1)
		post, _ := SelectContentByUid(postUid, "posts", dbInfo)
		blogPosts.BlogPosts = append(blogPosts.BlogPosts, post)

		if post == nil {
			return nil, nil
		}
	} else {
		posts, err := SelectMultipleContents(limit, limit*blogPosts.PageNum, Post_date, DESC, "posts", dbInfo)

		if err != nil || len(posts) < 1 {
			posts, _ = SelectMultipleContents(limit, 0, Post_date, DESC, "posts", dbInfo)
		}

		blogPosts.BlogPosts = append(blogPosts.BlogPosts, posts...)
	}

	tmpl := template.Must(template.New("base").
		Funcs(template.FuncMap{"formatDate": formatDate, "fieldExists": fieldExists}).
		ParseFiles("tmpl/blogpost.gohtml", "tmpl/blogpage.gohtml", "tmpl/base.gohtml"))
	return tmpl, blogPosts
}

func queryToUint32(queryParam string, min uint32) uint32 {
	value, err := strconv.Atoi(queryParam)
	if err != nil || value < int(min) {
		return min
	} else {
		return uint32(value)
	}
}

func fieldExists(name string, obj interface{}) bool {
	s := reflect.TypeOf(obj)
	_, b := s.FieldByName(name)
	return b
}
