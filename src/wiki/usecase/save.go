package usecase

import (
  "wiki/repository"
  "wiki/entity"
)

func Execute(s repository.Saver, title string, body []byte) error {
  p := &entity.Page{Title: title, Body: body}
  err := s.SavePage(p)
}
