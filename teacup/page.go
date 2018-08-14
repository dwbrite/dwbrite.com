package teacup

import (
	"html/template"
		)

//var pageTemplate = template.Must(template.ParseFiles("tmpl/page.gohtml", "tmpl/base.gohtml"))

type page struct {
	Title string
	Body template.HTML
}

/*
func servePage(writer http.ResponseWriter, title string) {
	p, err := selectPage(title)
	if checkError(writer, err) { return }

	err = pageTemplate.ExecuteTemplate(writer, "base", p)
	if checkError(writer, err) { return }
}

func selectPage(title string) (*page, error){
	db, _ := sql.Open("postgres", DB_INFO)
	defer db.Close()

	var body template.HTML
	err := db.QueryRow("SELECT * FROM pages WHERE title = $1;", title).Scan(&title, &body)
	if err != nil { return nil, err }
	return &page{Title: title, Body: body}, nil
}*/