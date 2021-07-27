package models

import "time"

type Group struct {
	ID          int       `db:"id"`
	OrganizerID int       `db:"organizer_id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Books       []Book    `db:"books"` // Comes from group_books table.
	Members     []User    `db:"members"`
	CreatedAt   time.Time `db:"created_at"`
}

type GroupStore interface {
	GetGroup(id int) (Group, error)
	GetGroups() ([]Group, error)
	CreateGroup(g *Group) (Group, error)
	UpdateGroup(g *Group) (Group, error)
	DeleteGroup(id int) error
}
