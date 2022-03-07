package controllers

import (
	"log"
	"net/http"
	"embed"
	"html/template"
	"io"
)

var data interface{}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

// PostRenderer renders data into HTML
type PostRenderer struct {
	templ    *template.Template
}

type WordList struct {
	Entries string
	Words []string
	Vocabs []Vocab
}

type Vocab struct {
	Word, Translation, Example string
}

var store WordList

// NewPostRenderer creates a new PostRenderer
func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ: templ}, nil 
}

func (r *PostRenderer) RenderIndex(w io.Writer, data interface{}) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", data)
}

func (r *PostRenderer) RenderBlog(w io.Writer, data interface{}) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", data)
}

func Home(w http.ResponseWriter, r *http.Request){
	store = WordList{Entries: "Hallo", Words: []string{"Hallo", "Selina"}, Vocabs: []Vocab{{"Hallo", "Hola", "Hola a todos"}, {"Tsch[ss", "Ciao", "Ciao a todos"},},}
	postRenderer, err := NewPostRenderer()
	if err != nil {
		log.Fatal("NewPostRenderer execution failed")
	}
	err = postRenderer.RenderIndex(w, store)
	if err != nil {
		log.Fatal("RenderIndex execution failed")
	}
}

func Blog(w http.ResponseWriter, r *http.Request){
	store = WordList{Entries: "Hallo", Words: []string{"Hallo", "Selina"}, Vocabs: []Vocab{{"Hallo", "Hola", "Hola a todos"}, {"Tsch[ss", "Ciao", "Ciao a todos"},},}
	postRenderer, err := NewPostRenderer()
	if err != nil {
		log.Fatal("NewPostRenderer execution failed")
	}
	err = postRenderer.RenderBlog(w, store)
	if err != nil {
		log.Fatal("RenderIndex execution failed")
	}
}