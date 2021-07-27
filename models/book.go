package models

type Book struct {
	ID       int    `db:"id"`
	Title    string `db:"title"`
	Author   string `db:"author"`
	Category string `db:"category"`
}

type BookStore interface {
	GetBook(id int) (Book, error)
	GetBooks() ([]Book, error)
	GetBooksByCategory() ([]Book, error)
	CreateBook(c *Book) (Book, error)
	UpdateBook(c *Book) (Book, error)
	DeleteBook(id int) error
}
