package postgres

import (
	"github.com/connorjcantrell/groups"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SectionStore struct {
	*sqlx.DB
}

func (*SectionStore) Section(id uuid.UUID) (groups.Section, error) {
	return groups.Section{}, nil
}

func (*SectionStore) SectionsByBook(id uuid.UUID) ([]groups.Section, error) {
	var ss []groups.Section
	return ss, nil
}

func (*SectionStore) CreateSection(s *groups.Section) (groups.Section, error) {
	return groups.Section{}, nil
}

func (*SectionStore) UpdateSection(s *groups.Section) (groups.Section, error) {
	return groups.Section{}, nil
}

func (*SectionStore) DeleteSection(id uuid.UUID) error {
	return nil
}
