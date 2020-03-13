package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"text/template"
)
/*Create mysql connection*/
func CreateCon() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "testgoapp"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
	panic(err.Error())
	} else {
		fmt.Printf("DB transaction done \n")
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("templates/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := CreateCon()
	selDB, err := db.Query("SELECT * FROM user ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	usr := User {}
	var res []User
	for selDB.Next() {
		var id int
		var name string
		err = selDB.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		usr.ID = id
		usr.Name = name
		res = append(res, usr)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := CreateCon()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM user WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	usr := User {}
	for selDB.Next() {
		var id int
		var name string
		err = selDB.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		usr.ID = id
		usr.Name = name
	}
	tmpl.ExecuteTemplate(w, "Show", usr)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := CreateCon()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM user WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	usr := User {}
	for selDB.Next() {
		var id int
		var name string
		err = selDB.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		usr.ID = id
		usr.Name = name
	}
	tmpl.ExecuteTemplate(w, "Edit", usr)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := CreateCon()
	if r.Method == "POST" {
		name := r.FormValue("name")
		insForm, err := db.Prepare("INSERT INTO user(name) VALUES(?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name)
		log.Println("INSERT: Name: " + name )
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := CreateCon()
	if r.Method == "POST" {
		name := r.FormValue("name")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE user SET name=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, id)
		log.Println("UPDATE: Name: " + name)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := CreateCon()
	usr := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM user WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(usr)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}