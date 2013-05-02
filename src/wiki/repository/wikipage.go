package repository

import (
  "wiki/entity"
)

type WikipageRepository interface {
  SavePage(p *entity.Page) error
  LoadPage(title string) (*entity.Page, error)
}
