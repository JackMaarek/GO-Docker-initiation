package main

import (
	"docker_app_go/goApp/db"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Server started on: http://localhost:8000")
	http.HandleFunc("/", db.Index)
	http.HandleFunc("/show", db.Show)
	http.HandleFunc("/new",db. New)
	http.HandleFunc("/edit", db.Edit)
	http.HandleFunc("/insert", db.Insert)
	http.HandleFunc("/update", db.Update)
	http.HandleFunc("/delete", db.Delete)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("Failed to listen on port 8000")
	}
}
