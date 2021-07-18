package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewStore(dataSourceName string) (*Store, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return &Store{
		UserStore:    &UserStore{DB: db},
		GroupStore:   &GroupStore{DB: db},
		EventStore:   &EventStore{DB: db},
		BookStore:    &BookStore{DB: db},
		ChapterStore: &ChapterStore{DB: db},
		SectionStore: &SectionStore{DB: db},
	}, nil
}

type Store struct {
	*UserStore
	*GroupStore
	*EventStore
	*BookStore
	*ChapterStore
	*SectionStore
}
