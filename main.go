package main

import (
	"api-go/components"
	"api-go/db"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	db := db.ConnDb()
	defer db.Close()
	http.HandleFunc("/", components.Index)
	http.ListenAndServe(":8000", nil)
}
