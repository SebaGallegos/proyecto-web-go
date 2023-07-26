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
	index = parse("index.html")
)

type IndexParams struct {
	Title string
}

func Index(w io.Writer, p IndexParams) error {
	return index.Execute(w, p)
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
