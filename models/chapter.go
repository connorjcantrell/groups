package models

type Chapter struct {
	ID       int       `db:"id"`
	BookID   int       `db:"book_id"`
	Title    string    `db:"title"`
	Number   int       `db:"number"`
	Sections []Section `db:"sections"`
}

type ChapterStore interface {
	GetChapter(id int) (Chapter, error)
	GetChaptersByBook(id int) ([]Chapter, error)
	CreateChapter(c *Chapter) (Chapter, error)
	UpdateChapter(c *Chapter) (Chapter, error)
	DeleteChapter(id int) error
}
