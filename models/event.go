package models

import "time"

type Event struct {
	ID          int           `db:"id"`
	GroupID     int           `db:"group_id"`
	BookID      int           `db:"book_id"`
	ChapterID   int           `db:"chapter_id"`
	VideoLink   string        `db:"video_link"`
	StartTime   time.Time     `db:"time"`
	Duration    time.Duration `db:"duration"`
	Description string        `db:"description"`
	CreatedAt   time.Time     `db:"created_at"`
}

type EventStore interface {
	GetEvent(id int) (Event, error)
	GetEventsByUser() ([]Event, error)
	GetEventsByGroup(id int) ([]Event, error)
	CreateEvent(e *Event) (Event, error)
	UpdateEvent(e *Event) (Event, error)
	DeleteEvent(id int) error
}
