package teacup

import (
	"database/sql"
	"strconv"
)

func SelectContentByUid(uid string, table string, dbInfo string) (*PageContents, error) {
	queryRow := func(db sql.DB) (*sql.Row, error) {
		i_uid, err := strconv.Atoi(uid)
		if err != nil { return nil, err }
		return db.QueryRow(`
SELECT pa.uid, pa.summary, pa.title, pa.body, pa.post_date 
FROM ` + table + ` pa
WHERE uid = $1;`, i_uid), nil
	}

	return selectContent(queryRow, dbInfo)
}

func SelectContentByTitle(title string, table string, dbInfo string) (*PageContents, error) {
	queryRow := func(db sql.DB) (*sql.Row, error) {
		return db.QueryRow(`
SELECT pa.uid, pa.summary, pa.title, pa.body, pa.post_date
FROM ` + table + ` pa
WHERE title = $1;`, title), nil
	}

	return selectContent(queryRow, dbInfo)
}

func SelectNSortedByUid(n int, offset int, table string, dbInfo string) ([]*PageContents, error) {
	
	return nil, nil
}

func SelectMultipleSortedByDate(n int, offset int, table string, dbInfo string) ([]*PageContents, error) {
	return nil, nil
}

func selectContent(queryRow func(db sql.DB) (*sql.Row, error), dbInfo string) (*PageContents, error){
	db, _ := sql.Open("postgres", dbInfo)
	defer db.Close()

	var p PageContents

	row, err := queryRow(*db)
	if err != nil { return nil, err }

	err = row.Scan(&p.Uid, &p.Summary, &p.Title, &p.Body, &p.PostDate)
	if err != nil { return nil, err }

	return &p, nil
}

func selectContents(query func(db sql.DB, p PageContents) (*sql.Rows, error), dbInfo string) (*PageContents, error){
	db, _ := sql.Open("postgres", dbInfo)
	defer db.Close()

	var p PageContents

	rows, err := query(*db, p)
	if err != nil { return nil, err }

	//TODO("Fill in rest of function")
	//err = rows.Scan(&p.Uid, &p.Summary, &p.Title, &p.Body, &p.PostDate)
	//if err != nil { return nil, err }

	return nil, nil
}

func (t *teacup) CreateTable(name string, uniqueTitles bool) {
	db, _ := sql.Open("postgres", t.DbInfo)
	defer db.Close()

	unique := ""
	if uniqueTitles {
		unique = "UNIQUE"
	}

	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS ` + name + ` (
		uid serial PRIMARY KEY,
		optional_summary varchar(192),
		title varchar(128) ` + unique + ` NOT NULL,
		body text NOT NULL,
		post_date date DEFAULT CURRENT_DATE NOT NULL
	) WITH (OIDS=FALSE)`)

	if err != nil { t.Log.Fatal(err, "\nTeacup could not connect to postgresql database.") }

	_, err = db.Exec(`
-- Summary function for body.
CREATE OR REPLACE FUNCTION summary(rec `+ name +`)
  RETURNS varchar(192)
LANGUAGE SQL
AS
$$
SELECT
       CASE WHEN $1.optional_summary IS NULL
                 THEN
           CASE WHEN length($1.body) > 192
                     THEN
               $1.body::varchar(191) || '…'
                ELSE
               $1.body::varchar(192)
               END
            ELSE
           $1.optional_summary
           END
$$;`)

	if err != nil { t.Log.Fatal(err) }

	t.tables[name] = uniqueTitles
	//TODO("check for duplicates")
}