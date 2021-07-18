package groups

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

type Group struct {
	ID          uuid.UUID `db:"id"`
	OrganizerID uuid.UUID `db:"organizer_id"`
	Name        string    `db:"name"`
	// TODO: How do I represent a Group's relationship to books in a database?
	// I suspect an aggregate table will be needed ("GroupBook" table)
	// However, I would like to search for better alternatives
	Books     []Book    `db:"books"`
	Members   []User    `db:"members"`
	CreatedAt time.Time `db:"created_at"`
}

type Event struct {
	ID          uuid.UUID     `db:"id"`
	GroupID     uuid.UUID     `db:"group_id"`
	BookID      uuid.UUID     `db:"book_id"`
	ChapterID   uuid.UUID     `db:"chapter_id"`
	StartTime   time.Time     `db:"time"`
	Duration    time.Duration `db:"duration"`
	Description string        `db:"description"`
	CreatedAt   time.Time     `db:"created_at"`
}

type Book struct {
	ID       uuid.UUID `db:"id"`
	Title    string    `db:"title"`
	Author   string    `db:"author"`
	Category string    `db:"category"`
}

type Chapter struct {
	ID       uuid.UUID `db:"id"`
	BookID   uuid.UUID `db:"book_id"`
	Title    string    `db:"title"`
	Number   int       `db:"number"`
	Sections []Section `db:"sections"`
}

type Section struct {
	ID     uuid.UUID `db:"id"`
	Title  string    `db:"title"`
	Number int       `db:"number"`
}

type Store interface {
	UserStore
	GroupStore
	EventStore
	BookStore
	ChapterStore
	SectionStore
}

type UserStore interface {
	User(id uuid.UUID) (User, error)
	UserByEmail(username string) (User, error)
	CreateUser(u *User) (User, error)
	UpdateUser(u *User) (User, error)
	DeleteUser(id uuid.UUID) error
}

type GroupStore interface {
	Group(id uuid.UUID) (Group, error)
	Groups() ([]Group, error)
	CreateGroup(g *Group) (Group, error)
	UpdateGroup(g *Group) (Group, error)
	DeleteGroup(id uuid.UUID) error
}

type EventStore interface {
	Event(id uuid.UUID) (Event, error)
	EventsByUser() ([]Event, error)
	EventsByGroup(id uuid.UUID) ([]Event, error)
	CreateEvent(e *Event) (Event, error)
	UpdateEvent(e *Event) (Event, error)
	DeleteEvent(id uuid.UUID) error
}

type BookStore interface {
	Book(id uuid.UUID) (Book, error)
	Books() ([]Book, error)
	BooksByCategory() ([]Book, error)
	CreateBook(c *Book) (Book, error)
	UpdateBook(c *Book) (Book, error)
	DeleteBook(id uuid.UUID) error
}

type ChapterStore interface {
	Chapter(id uuid.UUID) (Chapter, error)
	ChaptersByBook(id uuid.UUID) ([]Chapter, error)
	CreateChapter(c *Chapter) (Chapter, error)
	UpdateChapter(c *Chapter) (Chapter, error)
	DeleteChapter(id uuid.UUID) error
}

type SectionStore interface {
	Section(id uuid.UUID) (Section, error)
	SectionsByChapter(id uuid.UUID) ([]Section, error)
	SectionsByBook(id uuid.UUID) ([]Section, error)
	CreateSection(s *Section) (Section, error)
	UpdateSection(s *Section) (Section, error)
	DeleteSection(id uuid.UUID) error
}
