package view

import (
	"html/template"
	"net/http"
	"regexp"
)

type Page struct {
	Title string
}

var templates = template.Must(template.ParseFiles("static/index.html"))
var valid_path = regexp.MustCompile("send")

func Render(w http.ResponseWriter, tmpl string, page Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
