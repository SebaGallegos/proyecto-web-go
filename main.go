package main

import (
	"net/http"
	"proyecto-web-go/html"
)

func index(w http.ResponseWriter, r *http.Request) {
	p := html.IndexParams{Title: "Inicio"}
	html.Index(w, p)
}
func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
