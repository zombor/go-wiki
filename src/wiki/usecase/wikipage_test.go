package usecase

import (
  "wiki/entity"
  "wiki/repository"
  "github.com/orfjackal/gospec/src/gospec"
  . "github.com/orfjackal/gospec/src/gospec"
)

func WikipageSpec(c gospec.Context) {
  repository := repository.MemoryWikiPageRepository{}
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
