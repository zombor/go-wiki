package main

import (
  "wiki/entity"
  "wiki/repository/file"
  "wiki/usecase/wikipage"
  "net/http"
  "html/template"
)

const lenPath = len("/view/")
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))


func viewHandler(w http.ResponseWriter, r *http.Request) {
  repository := file.Page{}

  title := r.URL.Path[lenPath:]
  p, err := repository.LoadPage(title)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
  repository := file.Page{}

  title := r.URL.Path[lenPath:]
  p, err := repository.LoadPage(title)

  if err != nil {
    p = &entity.Page{Title: title}
  }

  renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[lenPath:]
  body := r.FormValue("body")
  repository := file.Page{}

  err := wikipage.ExecuteSave(repository, title, []byte(body))
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
  http.HandleFunc("/view/", viewHandler)
  http.HandleFunc("/edit/", editHandler)
  http.HandleFunc("/save/", saveHandler)
  http.ListenAndServe(":8080", nil)
}
