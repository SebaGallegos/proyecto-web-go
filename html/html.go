package html

import (
	"embed"
	"io"
	"strings"
	"text/template"
)

//go:embed *
var files embed.FS

// TODO: parse page names
var (
	index   = parse("index.html")
	pagina1 = parse("pagina1.html")
	pagina2 = parse("pagina2.html")
)

type IndexParams struct {
	Title string
}

// Páginas

func Index(w io.Writer, p IndexParams) error {
	return index.Execute(w, p)
}

func Pagina1(w io.Writer, p IndexParams) error {
	return pagina1.Execute(w, p)
}

func Pagina2(w io.Writer, p IndexParams) error {
	return pagina2.Execute(w, p)
}

var funcs = template.FuncMap{
	"uppercase": func(v string) string {
		return strings.ToUpper(v)
	},
}

func parse(file string) *template.Template {
	return template.Must(
		template.New("layout.html").Funcs(funcs).ParseFS(files, "layout.html", file))
}
