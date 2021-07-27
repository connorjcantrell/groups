package models

type Section struct {
	ID     int    `db:"id"`
	Title  string `db:"title"`
	Number int    `db:"number"`
}

type SectionStore interface {
	GetSection(id int) (Section, error)
	GetSectionsByChapter(id int) ([]Section, error)
	CreateSection(s *Section) (Section, error)
	UpdateSection(s *Section) (Section, error)
	DeleteSection(id int) error
}
