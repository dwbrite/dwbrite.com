/*  A personal web server intended for hosting dwbrite.com
 *  Copyright (C) 2018 Devin Brite
 *  Licensing information can be found in `README.md` and `LICENSE.md`
 */

package main

import (
	. "./teacup"
	"database/sql"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
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

type staticPage struct {
	Title string
	StaticContent template.HTML
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
			Cert: "certs/dwbrite.com.cert",
			Key:  "certs/dwbrite.com.key",
		},
		[]string{"html", "css", "scss", "map", "js", "png", "jpg", "gif", "webm", "ico", "md", "mp3", "mp4", "ttf", "woff", "woff2", "eot"},//*regexp.MustCompile("^/.*\\.(html|css|scss|map|js|png|jpg|gif|webm|ico|md|mp3|mp4|ttf|woff|woff2|eot)$"),
		[]string{"/certs", "/examples", "/tmpl"},//*regexp.MustCompile("^/(certs|examples|tmpl)/?.*$"),

		//*regexp.MustCompile("^/.*\\.(html|css|scss|map|js|png|jpg|gif|webm|ico|md|mp3|mp4|ttf|woff|woff2|eot)$")

		log.New(file, "", log.LstdFlags|log.Lshortfile),
	)

	t.CreateTable("posts", false)
	t.CreateTable("projects", true)

	// error template
	errTmpl := template.Must(template.New("base").
		Funcs(template.FuncMap{"fieldExists": fieldExists}).
		ParseFiles("tmpl/errors.gohtml", "tmpl/base.gohtml"))
	t.SetErrorTemplate(errTmpl)



	t.HandleFunc("/blog", blogQuery)
	t.HandleFunc("/blog/page", blogQuery)
	t.HandleFunc("/blog/post", blogQuery)

	t.HandleFunc("/portfolio", projectQuery)
	t.HandleFunc("/", home)
	t.Mux.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./resources"))))

	t.StartServer()
}

func home(_ http.Request, _ string) (*TemplateContent, error) {
	pageTmpl := template.Must(template.New("base").
		Funcs(template.FuncMap{"fieldExists": fieldExists}).
		ParseFiles("tmpl/static.gohtml", "tmpl/base.gohtml"))

	content, err := ioutil.ReadFile("home.html")
	if err != nil {
		return nil, nil
	}

	home := staticPage {
		"Home Page",
		template.HTML(string(content)),
	}

	return &TemplateContent{ pageTmpl, home }, nil
}

func projectQuery(_ http.Request, _ string) (*TemplateContent, error) {
	return nil, nil
}

func formatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func page(request http.Request, dbInfo string) (*TemplateContent, error) {
	//var limit uint32 = 3

	/*blog := blogPage{
		[]*PageContents{},
		"Blog",
		nil, nil, nil, nil, nil,
	}*/

	return nil, nil
}

func blogQuery(request http.Request, dbInfo string) (*TemplateContent, error) {
	// If a specific post is requested, return that post.
	//   If no such post exists, return (nil, nil) | (404)
	// Otherwise
	//   If a page is requested, return that page.
	//   If the page doesn't exist, just return the first page.
	var limit uint32 = 3

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
		blog.LastPage = &onum {calcNumBlogPages(dbInfo, limit) }

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

	return &TemplateContent{tmpl, blog}, nil
}

func strToOnum(str string, min uint32) *onum {
	value, err := strconv.Atoi(str)
	if err != nil || value < int(min) {
		return nil
	} else {
		return &onum { uint32(value) }
	}
}

func fieldExists(name string, obj interface{}) bool {
	s := reflect.TypeOf(obj)
	_, b := s.FieldByName(name)
	return b
}

func calcNumBlogPages(dbInfo string, limit uint32) uint32 {
	db, _ := sql.Open("postgres", dbInfo)
	defer db.Close()

	var count int
	rows := db.QueryRow("SELECT COUNT(*) as count FROM posts")
	rows.Scan(&count)

	ct := uint32(count)

	pages := ct/limit
	if ct%limit != 0 {
		pages++
	}
	return pages
}