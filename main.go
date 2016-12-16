package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(writer http.ResponseWriter, request *http.Request) {
	page := "index.gohtml"
	zip := request.FormValue("zip")
	reps := zipLookup(zip)
	if zip != "" {
		page = "results.gohtml"
	}
	err := templates.ExecuteTemplate(writer, page, reps.Results)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		log.Fatalln(err)
	}
}
