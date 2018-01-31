package main

import (
	"net/http"
	"html/template"
	"strings"
	"net/url"
)

const (
	resourcesDir = "./resources"
)

type extURL = url.URL

func (u *extURL) last() string {
	splitUrl := strings.Split(u.Path, "/")
	return splitUrl[len(splitUrl) -1]
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, _ := template.ParseFiles(resourcesDir + tmpl + ".html")
	t.Execute(w, data)
}

func renderErrorTemplate(w http.ResponseWriter) {
}
