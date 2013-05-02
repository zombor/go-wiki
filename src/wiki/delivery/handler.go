package delivery

import (
  "github.com/gorilla/mux"
)

func Handle(r **mux.Router) {
  r.HandleFunc("/view/{title}", viewHandler)
  r.HandleFunc("/edit/{title}", editHandler)
  r.HandleFunc("/save/{title}", saveHandler)
}

