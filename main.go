package main

import (
	"github.com/mortdeus/go-wiki/entity"
	"github.com/mortdeus/go-wiki/repository/file"
	"github.com/mortdeus/go-wiki/usecase/wikipage"
	"html/template"
	"net/http"
)

const lenPath = len("/view/")

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	p := new(file.Page)
	if err := p.LoadPage(title); err != nil {
		panic(err)
	}

	renderTemplate(w, "view", &p.Page)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	p := new(file.Page)
	if err := p.LoadPage(title); err != nil {
		panic(err)
	}
	renderTemplate(w, "edit", &p.Page)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	body := r.FormValue("body")

	err := wikipage.ExecuteSave(new(file.Page), title, []byte(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *entity.Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.ListenAndServe(":8080", nil)
}

func init() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

}
