package postgres

import (
	"github.com/connorjcantrell/groups"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type EventStore struct {
	*sqlx.DB
}

func (s *EventStore) Event(id uuid.UUID) (groups.Event, error) {
	return groups.Event{}, nil
}

func (s *EventStore) EventsByUser() ([]groups.Event, error) {
	var ee []groups.Event
	return ee, nil
}

func (s *EventStore) EventsByGroup(id uuid.UUID) ([]groups.Event, error) {
	var ee []groups.Event
	return ee, nil
}

func (s *EventStore) CreateEvent(e *groups.Event) (groups.Event, error) {
	return groups.Event{}, nil
}

func (s *EventStore) UpdateEvent(e *groups.Event) (groups.Event, error) {
	return groups.Event{}, nil
}

func (s *EventStore) DeleteEvent(id uuid.UUID) error {
	return nil
}
