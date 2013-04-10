package wikipage

import (
  "wiki/entity"
)

type Saver interface {
  SavePage(p *entity.Page) error
}

type Loader interface {
  LoadPage() error
}

