// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"regexp"
	"database/sql"
        _ "github.com/lib/pq"

)
const (
	port = 5432
)


func connDB() (*sql.DB, error) {
	// connection string
	host := os.Getenv("DBHOST")
	password :=  os.Getenv("DBPSW")
	dbname := os.Getenv("POSTGRES_DB")
	user := os.Getenv("POSTGRES_USER")
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println("psqlconn: %s", psqlconn)

	// open database
	db, err := sql.Open("postgres", psqlconn)	
	
	return db, err
}

func readDB(key string) ([]byte, error) {
	db, err := connDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	
	rows, err := db.Query("SELECT name FROM mypages WHERE key = $1", key)
	if err != nil {
		return nil, err
	}

	nxt := rows.Next()
	if !nxt {
		return nil, nil
	}
	var name string
	err = rows.Scan(&name)
	return []byte(name), err
}
func writeDB(key string, name string) error {
	db, err := connDB()
	if err != nil {
		return err
	}
	defer db.Close()

	deleteDynStmt := `delete from mypages where key = $1`
	_, err = db.Exec(deleteDynStmt, key)
	if err != nil {
		return err
	}

	insertDynStmt := `insert into mypages (key,name) values($1, $2)`
	_, err = db.Exec(insertDynStmt, key, name)
	return err
}

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	if strings.HasPrefix(p.Title,"db") {
		return writeDB(p.Title, string(p.Body))
	}
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	var body []byte
	var err error 
	if strings.HasPrefix(title,"db") {
		body, err = readDB(title)
	} else {
		filename := title + ".txt"
		body, err = os.ReadFile(filename)
	}
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("url path %s\n", r.URL.Path)
		m := validPath.FindStringSubmatch(r.URL.Path)
		fmt.Printf("match found %s\n", m)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
