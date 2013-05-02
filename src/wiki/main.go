package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "wiki/delivery"
)

func main() {
  r := mux.NewRouter()
  delivery.Handle(&r)
  http.Handle("/", r)
  http.ListenAndServe(":8080", nil)
}
