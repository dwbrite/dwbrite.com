package teacup

import (
	"database/sql"
	"strconv"
)

func SelectContentByUid(uid string, table string, dbInfo string) (*PageContents, error) {
	queryRow := func(db sql.DB, p PageContents) (*sql.Row, error) {
		i_uid, err := strconv.Atoi(uid)
		if err != nil { return nil, err }
		return db.QueryRow(`
SELECT pa.uid, pa.summary, pa.title, pa.body, pa.post_date 
FROM ` + table + ` pa
WHERE uid = $1;`, i_uid), nil
	}

	return selectContents(queryRow, dbInfo)
}

func SelectContentByTitle(title string, table string, dbInfo string) (*PageContents, error) {
	queryRow := func(db sql.DB, p PageContents) (*sql.Row, error) {
		return db.QueryRow(`
SELECT pa.uid, pa.summary, pa.title, pa.body, pa.post_date
FROM ` + table + ` pa
WHERE title = $1;`, title), nil
	}

	return selectContents(queryRow, dbInfo)
}

func selectContents(queryRow func(db sql.DB, p PageContents) (*sql.Row, error), dbInfo string) (*PageContents, error){
	db, _ := sql.Open("postgres", dbInfo)
	defer db.Close()

	var p PageContents

	row, err := queryRow(*db, p)
	if err != nil { return nil, err }

	err = row.Scan(&p.Uid, &p.Summary, &p.Title, &p.Body, &p.PostDate)
	if err != nil { return nil, err }

	return &p, nil
}

func (t *teacup) CreateTable(name string, uniqueTitles bool) sql.Result {
	db, _ := sql.Open("postgres", t.DbInfo)
	defer db.Close()

	unique := ""
	if uniqueTitles {
		unique = "UNIQUE"
	}

	result, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS ` + name + ` (
		uid serial PRIMARY KEY,
		summary varchar(192),
		title varchar(128) ` + unique + ` NOT NULL,
		body text NOT NULL,
		post_date date DEFAULT CURRENT_DATE NOT NULL
	) WITH (OIDS=FALSE)`)

	if err != nil {
		t.Log.Fatal(err)
	}

	t.tables[name] = uniqueTitles
	//TODO("check for duplicates")

	return result
}