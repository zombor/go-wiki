package wikipage

import (
	"github.com/mortdeus/go-wiki/entity"
	"github.com/mortdeus/go-wiki/repository"
	"github.com/mortdeus/go-wiki/repository/file"
	"reflect"
)

func ExecuteSave(s repository.Saver, title string, body []byte) error {
	switch p := s.(type) {
	case *file.Page:
		p.Page.Title = title
		p.Page.Body = body
		ExecuteSave(&p.Page, "", nil)
	case *entity.Page:
		return p.SavePage()
	default:
		pv := reflect.ValueOf(p)
		pv.FieldByName("Title").SetString(title)
		pv.FieldByName("Body").SetBytes(body)
		return p.SavePage()
	}
	return nil
}
