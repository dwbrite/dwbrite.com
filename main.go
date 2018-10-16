/*  A personal web server intended for hosting dwbrite.com
 *  Copyright (C) 2018 Devin Brite
 *  Licensing information can be found in `README.md` and `LICENSE.md`
 */

package main

import (
	. "./teacup"
	"database/sql"
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

type onum struct {
	Value uint32
}

type blogPage struct {
	Posts     []*PageContents
	Title     string
	CurrPage  *onum
	FirstPage *onum
	LastPage  *onum
	PrevPage  *onum
	NextPage  *onum
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

	t.CreateTable("Posts", false)
	t.CreateTable("projects", true)
	//t.CreateTable("pages", true)

	// error template
	errTmpl := template.Must(template.New("base").
		Funcs(template.FuncMap{"fieldExists": fieldExists}).
		ParseFiles("tmpl/page.gohtml", "tmpl/base.gohtml"))
	t.SetErrorTemplate(errTmpl)

	//"^...?.*?$"
	t.AddDynamicPage("^/blog/?.*?$", "Posts", blogQuery)

	t.StartServer()
}

func formatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func blogQuery(request http.Request, dbInfo string) (*template.Template, interface{}) {
	// If a specific post is requested, return that post.
	//   If no such post exists, return (nil, nil) | (404)
	// Otherwise
	//   If a page is requested, return that page.
	//   If the page doesn't exist, just return the first page.

	var limit uint32 = 5
	blog := blogPage{
		[]*PageContents{},
		"Blog",
		nil, nil, nil, nil, nil,
	}

	query := request.URL.Query()
	postNum := strToOnum(query.Get("post"), 1)

	blog.CurrPage = strToOnum(query.Get("page"), 0)

	if postNum != nil {
		postUid := strToOnum(query.Get("post"), 1)
		post, _ := SelectContentByUid(postUid.Value, "posts", dbInfo)
		blog.Posts = append(blog.Posts, post)

		if post == nil {
			return nil, nil
		}
	} else {
		var posts []*PageContents

		blog.FirstPage = &onum {1 }
		blog.LastPage = &onum {getNumPages(dbInfo, "posts", limit) }

		if blog.CurrPage == nil {
			blog.CurrPage = blog.LastPage
		} else if blog.CurrPage.Value == blog.LastPage.Value {
			return nil, nil
		}

		if blog.CurrPage.Value < blog.LastPage.Value {
			blog.NextPage = &onum { blog.CurrPage.Value +1 }
		}
		if blog.CurrPage.Value > blog.FirstPage.Value {
			blog.PrevPage = &onum { blog.CurrPage.Value -1 }
		}

		var tempPosts []*PageContents
		tempPosts, _ = SelectMultipleContents(limit, limit*(blog.CurrPage.Value-1), Post_date, ASC, "posts", dbInfo)

		if tempPosts == nil {
			return nil, nil
		} else {
			for _, p := range tempPosts {
				posts = append([]*PageContents{p}, posts...)
			}
		}

		blog.Posts = append(blog.Posts, posts...)
	}

	tmpl := template.Must(template.New("base").
		Funcs(template.FuncMap{"formatDate": formatDate, "fieldExists": fieldExists}).
		ParseFiles("tmpl/blogpost.gohtml", "tmpl/blogpage.gohtml", "tmpl/base.gohtml"))
	return tmpl, blog
}

func strToOnum(str string, min uint32) *onum {
	value, err := strconv.Atoi(str)
	if err != nil || value < int(min) {
		return nil
	} else {
		return &onum { uint32(value) };
	}
}

func fieldExists(name string, obj interface{}) bool {
	s := reflect.TypeOf(obj)
	_, b := s.FieldByName(name)
	return b
}

func getNumPages(dbInfo string, table string, limit uint32) uint32 {
	db, _ := sql.Open("postgres", dbInfo)
	defer db.Close()

	var count int
	rows := db.QueryRow("SELECT COUNT(*) as count FROM posts;")
	rows.Scan(&count)

	ct := uint32(count)

	pages := ct/limit
	if ct%limit != 0 {
		pages++
	}
	return pages
}
