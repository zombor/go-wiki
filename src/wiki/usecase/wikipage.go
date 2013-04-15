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

type LoadWikipage struct {
  PageRepository wikipage.Loader
}

func (u *LoadWikipage) Execute(title string) (*entity.Page, error) {
  page, err := u.PageRepository.LoadPage(title)

  return page, err
}
