package usecase

import (
  "wiki/entity"
  "errors"
  "github.com/orfjackal/gospec/src/gospec"
  . "github.com/orfjackal/gospec/src/gospec"
)

type Repository struct {
  Pages map[string]*entity.Page
}

func (repo Repository) SavePage(p *entity.Page) error {
  repo.Pages[p.Title] = p

  if true == false {
    return errors.New("foobar")
  }

  return nil
}

func (repo Repository) LoadPage(title string) (*entity.Page, error) {
  return repo.Pages[title], nil
}

func WikipageSpec(c gospec.Context) {
  repository := Repository{}
  repository.Pages = make(map[string]*entity.Page)

  c.Specify("Saves a wiki page", func() {
    usecase := SaveWikipage{repository}
    usecase.Execute("foo", []byte("body"))

    page, err := repository.LoadPage("foo")

    c.Expect(err, Equals, nil)
    c.Expect(string(page.Body), Equals, "body")
  })

  c.Specify("Loads a wiki page", func() {
    repository.Pages["testing"] = &entity.Page{Title: "testing", Body: []byte("the body")}
    usecase := LoadWikipage{repository}
    page, err := usecase.Execute("testing")

    c.Expect(err, Equals, nil)
    c.Expect(string(page.Body), Equals, "the body")
  })
}
