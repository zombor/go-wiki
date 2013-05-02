package repository

import (
  "wiki/entity"
  "io/ioutil"
)

type WikipageRepository interface {
  SavePage(p *entity.Page) error
  LoadPage(title string) (*entity.Page, error)
}

type FileWikiPageRepository struct {
}

func (repo FileWikiPageRepository) LoadPage(title string) (*entity.Page, error) {
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename)

  if err != nil {
    return nil, err
  }

  return &entity.Page{Title: title, Body: body}, nil
}

func (repo FileWikiPageRepository) SavePage(p *entity.Page) error {
  filename := p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
}
