package repository

type Saver interface {
	SavePage() error
}

type Loader interface {
	LoadPage(string) error
}