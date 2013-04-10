package file

import (
  "io/ioutil"
  "wiki/entity"
)

type WikiPage struct {
}

func (repo WikiPage) LoadPage(title string) (*entity.Page, error) {
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename)

  if err != nil {
    return nil, err
  }

  return &entity.Page{Title: title, Body: body}, nil
}

func (repo WikiPage) SavePage(p *entity.Page) error {
  filename := p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
}