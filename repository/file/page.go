package file

import (
	"github.com/mortdeus/go-wiki/entity"
)

type Page struct {
	entity.Page
}

func (p *Page) LoadPage(title string) error {
	return (&p.Page).LoadPage(title)

}

func (p *Page) SavePage() error {
	return (&p.Page).SavePage()
}
