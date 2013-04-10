package usecase

import (
  "wiki/repository"
  "wiki/entity"
)

type SaveWikipage struct {
  PageRepository wikipage.Saver
}

func (u *SaveWikipage) Execute(title string, body []byte) error {
  p := &entity.Page{Title: title, Body: body}
  return u.PageRepository.SavePage(p)
}
