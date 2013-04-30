package view

import "wiki/entity"

type Page struct {
  Content *entity.Page
}

func (page Page) Title() string {
  return page.Content.Title
}

func (page Page) Body() string {
  return string(page.Content.Body)
}
