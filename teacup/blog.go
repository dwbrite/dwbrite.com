package teacup

import (
				"time"
	)

type post struct {
	Uid  uint32
	PostDate time.Time
	Content page
}

type postList struct {
	BlogPosts []post
	Title string //Page Title
}

func formatDate(t time.Time) string {
	{ return t.Format("2006-01-02") }
}

/*
var postTemplate = template.Must(
	template.New("").Funcs(template.FuncMap{ "formatDate": formatDate }).
		ParseFiles("tmpl/blogpost.gohtml", "tmpl/base.gohtml"))
var postsTemplate = template.Must(
	template.New("").Funcs(template.FuncMap{ "formatDate": formatDate }).
		ParseFiles("tmpl/blogpost.gohtml", "tmpl/blogpage.gohtml", "tmpl/base.gohtml"))
*/

/*
func servePost(writer http.ResponseWriter, uid string) {
	p, err := selectPost(uid)
	if checkError(writer, err) { return }

	err = postTemplate.ExecuteTemplate(writer, "base", p)
	if checkError(writer, err) { return }
}

func selectPost(uid_s string) (*post, error){
	db, _ := sql.Open("postgres", DB_INFO)
	defer db.Close()

	var uid uint32
	var title string
	var body template.HTML
	var date time.Time

	err := db.QueryRow("SELECT * FROM posts WHERE uid = $1;", uid_s).Scan(&uid, &date, &title, &body)
	if err != nil { return nil, err }
	return &post{Uid: uid, PostDate: date, Content: page{Title: title, Body: body}}, nil
}

func selectPosts(query string) ([]post, error) {
	db, _ := sql.Open("postgres", DB_INFO)
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil { return nil, err }

	var posts []post
	for rows.Next() {
		var uid uint32
		var title string
		var body template.HTML
		var date time.Time
		err = rows.Scan(&uid, &title, &body, &date)
		if err != nil { return nil, err }
		posts = append(posts, post{Uid: uid, PostDate: date, Content: page{Title: title, Body: body}})
	}

	return posts, nil
}

func serveBlog(writer http.ResponseWriter, request *http.Request) {
	blogPosts := postList{ nil, "Blog" }

	if pageNum == -1 {
		posts, err := selectPosts("SELECT * FROM " +
			"(SELECT * FROM posts ORDER BY post_date DESC " +
			"LIMIT " + strconv.Itoa(POSTS_PER_PAGE) + ") AS derivedTable " +
			"ORDER BY post_date DESC;")
		if checkError(writer, err) {
			serveError(writer, http.StatusInternalServerError)
			return
		}

		blogPosts.BlogPosts = posts
		err = postsTemplate.ExecuteTemplate(writer, "base", blogPosts)
		if checkError(writer, err) {
			serveError(writer, http.StatusInternalServerError)
			return
		}
	} else {
		posts, err := selectPosts("SELECT * FROM posts ORDER BY post_date ASC " +
			"LIMIT " + strconv.Itoa(POSTS_PER_PAGE) +
			"OFFSET " + strconv.Itoa(POSTS_PER_PAGE * (pageNum)) + ";")
		if checkError(writer, err) {
			serveError(writer, http.StatusInternalServerError)
			return
		}

		blogPosts.BlogPosts = posts
		err = postsTemplate.ExecuteTemplate(writer, "base", blogPosts)
		if checkError(writer, err) {
			serveError(writer, http.StatusInternalServerError)
			return
		}
	}
}*/