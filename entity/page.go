package entity

import "io/ioutil"

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) LoadPage(title string) error {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}
	p.Title = title
	p.Body = body
	return nil
}

func (p *Page) SavePage() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}
