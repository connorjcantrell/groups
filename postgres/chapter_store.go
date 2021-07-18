package postgres

import (
	"github.com/connorjcantrell/groups"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ChapterStore struct {
	*sqlx.DB
}

func (s *ChapterStore) Chapter(id uuid.UUID) (groups.Chapter, error) {
	return groups.Chapter{}, nil
}

func (s *ChapterStore) ChaptersByBook(id uuid.UUID) ([]groups.Chapter, error) {
	var gg []groups.Chapter
	return gg, nil
}

func (s *ChapterStore) CreateChapter(c *groups.Chapter) (groups.Chapter, error) {
	return groups.Chapter{}, nil
}
func (s *ChapterStore) UpdateChapter(c *groups.Chapter) (groups.Chapter, error) {
	return groups.Chapter{}, nil
}

func (s *ChapterStore) DeleteChapter(id uuid.UUID) error {
	return nil
}
