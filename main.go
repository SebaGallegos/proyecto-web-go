package main

import (
	"net/http"
	"proyecto-web-go/html"
)

func index(w http.ResponseWriter, r *http.Request) {
	p := html.IndexParams{Title: "Inicio"}
	html.Index(w, p)
}

func pagina1(w http.ResponseWriter, r *http.Request) {
	p := html.IndexParams{Title: "Página 1"}
	html.Pagina1(w, p)
}

func pagina2(w http.ResponseWriter, r *http.Request) {
	p := html.IndexParams{Title: "Página 2"}
	html.Pagina2(w, p)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/1", pagina1)
	http.HandleFunc("/2", pagina2)
	http.ListenAndServe(":8080", nil)
}
