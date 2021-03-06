package delivery

import (
  "fmt"
  "net/http"
  "wiki/entity"
  "wiki/repository"
  "wiki/usecase"
  "wiki/view"
  "github.com/hoisie/mustache"
  "github.com/gorilla/mux"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
  usecase := usecase.NewLoadWikipage(repository.FileWikiPageRepository{})

  vars := mux.Vars(r)
  title := vars["title"]
  p, err := usecase.Execute(title)

  if err != nil {
    fmt.Println(err)
    http.Redirect(w, r, "/edit/"+title, http.StatusFound)
    return
  }

  renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
  usecase := usecase.NewLoadWikipage(repository.FileWikiPageRepository{})

  vars := mux.Vars(r)
  title := vars["title"]
  p, err := usecase.Execute(title)

  if err != nil {
    p = &entity.Page{Title: title}
  }

  renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  title := vars["title"]
  body := r.FormValue("body")
  usecase := usecase.NewSaveWikipage(repository.FileWikiPageRepository{})

  err := usecase.Execute(title, []byte(body))
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *entity.Page) {
  view := view.Page{Content: p}
  output := mustache.RenderFile("templates/"+tmpl+".mustache", view)
  fmt.Fprintf(w, output)
}
