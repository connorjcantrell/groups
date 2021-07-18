package postgres

import (
	"github.com/connorjcantrell/groups"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type GroupStore struct {
	*sqlx.DB
}

func (s *GroupStore) Group(id uuid.UUID) (groups.Group, error) {
	return groups.Group{}, nil
}

func (s *GroupStore) Groups() ([]groups.Group, error) {
	var gg []groups.Group
	return gg, nil
}

func (s *GroupStore) CreateGroup(g *groups.Group) (groups.Group, error) {
	return groups.Group{}, nil
}

func (s *GroupStore) UpdateGroup(g *groups.Group) (groups.Group, error) {
	return groups.Group{}, nil
}

func (s *GroupStore) DeleteGroup(id uuid.UUID) error {
	return nil
}
