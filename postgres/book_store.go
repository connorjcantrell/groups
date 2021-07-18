package postgres

import (
	"github.com/connorjcantrell/groups"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type BookStore struct {
	*sqlx.DB
}

func (s *BookStore) Book(id uuid.UUID) (groups.Book, error) {
	return groups.Book{}, nil
}

func (s *BookStore) Books() ([]groups.Book, error) {
	var bb []groups.Book
	return bb, nil
}

func (s *BookStore) BooksByCategory() ([]groups.Book, error) {
	var bb []groups.Book
	return bb, nil
}

func (s *BookStore) CreateBook(c *groups.Book) (groups.Book, error) {
	return groups.Book{}, nil
}

func (s *BookStore) UpdateBook(c *groups.Book) (groups.Book, error) {
	return groups.Book{}, nil
}

func (s *BookStore) DeleteToolEntry(id uuid.UUID) error {
	return nil
}
