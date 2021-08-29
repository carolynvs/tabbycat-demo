package main

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	_ "embed"
)

//go:embed templates
var templates embed.FS

func main() {
	http.HandleFunc("/", List)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func List(w http.ResponseWriter, r *http.Request) {
	err := renderList(w)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
}

func renderList(w io.Writer) error {
	tmpl, err := template.ParseFS(templates, "templates/*")
	if err != nil {
		return fmt.Errorf("cannot load templates: %s", err.Error())
	}

	data := struct {
		Favorites []Favorite
	}{
		Favorites: []Favorite{
			{URL: "http://tabbycats.club/cat/jvzpqs"},
			{URL: "https://tabbycats.club/cat/a637t5"},
			{URL: "https://tabbycats.club/cat/w4266a"},
			{URL: "https://tabbycats.club/cat/01ujg2"},
		},
	}
	return tmpl.ExecuteTemplate(w, "list.html", data)
}

type Favorite struct {
	Name string
	URL  string
}
