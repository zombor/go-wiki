package repository

import (
  "wiki/entity"
)

type MemoryWikiPageRepository struct {
  Pages map[string]*entity.Page
}

func (repo MemoryWikiPageRepository) LoadPage(title string) (*entity.Page, error) {
  return repo.Pages[title], nil
}

func (repo MemoryWikiPageRepository) SavePage(p *entity.Page) error {
  repo.Pages[p.Title] = p

  return nil
}

