/*  A personal web server intended for hosting dwbrite.com
 *  Copyright (C) 2020 Devin Brite
 */

package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type BlogPost struct {
	Summary  string
	Title    string
	Body     template.HTML
	PostDate time.Time
	Link     string
}

type BlogPage struct {
	Posts     []*BlogPost
	Title     string
	CurrPage  *OptNum
	FirstPage *OptNum
	LastPage  *OptNum
	PrevPage  *OptNum
	NextPage  *OptNum
}

type OptNum struct {
	Value uint32
}

func (s *server) serveBlog(internalPath string, m *metadata) {
	pathDirectory, reqPath := getPaths(internalPath)

	data := m.Data.(map[string]interface{})
	title := data["title"].(string)

	b := &BlogRegister{
		s,
		[]*BlogPost{},
		title,
	}

	filepath.Walk(pathDirectory, b.walkBlog)

	// TODO: sort blog posts
	sort.Slice(b.posts, func(i, j int) bool {
		return b.posts[i].PostDate.Before(b.posts[j].PostDate)
	})

	page := BlogPage{
		Posts:     b.posts,
		Title:     title,
		CurrPage:  &OptNum{0},
		FirstPage: &OptNum{0},
		LastPage:  &OptNum{0},
		PrevPage:  nil,
		NextPage:  nil,
	}

	cache := cache{
		new(bytes.Buffer),
	}

	// TODO: execute the cache for each page

	_ = template.Must(template.New("base").
		Funcs(template.FuncMap{"formatDate": formatDate, "fieldExists": fieldExists}).
		ParseFiles("tmpl/blogpost.gohtml", "tmpl/blogpage.gohtml", "tmpl/base.gohtml")).
		Execute(cache, page)

	s.HandleFunc(reqPath, cache.serve)
}

type BlogRegister struct {
	s     *server
	posts []*BlogPost
	title string
}

func (b *BlogRegister) walkBlog(internalPath string, file os.FileInfo, err error) error {
	pathDirectory, reqPath := getPaths(internalPath)

	if "metadata.json" != file.Name() {
		return nil
	}

	metadata := b.s.readMetadata(internalPath)
	if metadata.Archetype != "blogpost" {
		return nil
	}

	data := metadata.Data.(map[string]interface{})

	date, _ := time.Parse("2006-01-02", data["date"].(string))

	postBody, _ := ioutil.ReadFile(filepath.Join(pathDirectory, "index.html"))

	post := &BlogPost{
		Title:    data["title"].(string),
		Body:     template.HTML(postBody),
		PostDate: date,
		Link:     reqPath,
	}

	b.posts = append(b.posts, post)

	cache := cache{
		new(bytes.Buffer),
	}

	// TODO: execute the cache for each page

	page := BlogPage{
		Posts: []*BlogPost{post},
		Title: b.title,
	}

	_ = template.Must(template.New("base").
		Funcs(template.FuncMap{"formatDate": formatDate, "fieldExists": fieldExists}).
		ParseFiles("tmpl/blogpost.gohtml", "tmpl/blogpage.gohtml", "tmpl/base.gohtml")).
		Execute(cache, page)

	b.s.HandleFunc(reqPath, cache.serve)

	return nil
}
