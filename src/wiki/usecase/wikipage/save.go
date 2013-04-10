package wikipage

import (
  "wiki/repository"
  "wiki/entity"
)

func ExecuteSave(s repository.Saver, title string, body []byte) error {
  p := &entity.Page{Title: title, Body: body}
  return s.SavePage(p)
}
