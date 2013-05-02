package usecase

import (
  "wiki/repository"
  "wiki/entity"
)

type SaveWikipage struct {
  pageRepository repository.WikipageRepository
}

func NewSaveWikipage(repo repository.FileWikiPageRepository) (SaveWikipage) {
  return SaveWikipage{pageRepository: repo}
}

func (u *SaveWikipage) Execute(title string, body []byte) error {
  p := &entity.Page{Title: title, Body: body}
  return u.pageRepository.SavePage(p)
}

type LoadWikipage struct {
  pageRepository repository.WikipageRepository
}

func NewLoadWikipage(repo repository.FileWikiPageRepository) (LoadWikipage) {
  return LoadWikipage{pageRepository: repo}
}

func (u *LoadWikipage) Execute(title string) (*entity.Page, error) {
  page, err := u.pageRepository.LoadPage(title)

  return page, err
}
