package groups

import (
	"time"
)

type User struct {
	ID        int       `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

type Group struct {
	ID          int       `db:"id"`
	OrganizerID int       `db:"organizer_id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Books       []Book    `db:"books"`
	Members     []User    `db:"members"`
	CreatedAt   time.Time `db:"created_at"`
}

type Event struct {
	ID          int           `db:"id"`
	GroupID     int           `db:"group_id"`
	BookID      int           `db:"book_id"`
	ChapterID   int           `db:"chapter_id"`
	StartTime   time.Time     `db:"time"`
	Duration    time.Duration `db:"duration"`
	Description string        `db:"description"`
	CreatedAt   time.Time     `db:"created_at"`
}

type Book struct {
	ID       int    `db:"id"`
	Title    string `db:"title"`
	Author   string `db:"author"`
	Category string `db:"category"`
}

type Chapter struct {
	ID       int       `db:"id"`
	BookID   int       `db:"book_id"`
	Title    string    `db:"title"`
	Number   int       `db:"number"`
	Sections []Section `db:"sections"`
}

type Section struct {
	ID     int    `db:"id"`
	Title  string `db:"title"`
	Number int    `db:"number"`
}

type UserStore interface {
	GetUser(id int) (User, error)
	GetUserByEmail(username string) (User, error)
	CreateUser(u *User) (User, error)
	UpdateUser(u *User) (User, error)
	DeleteUser(id int) error
}

type GroupStore interface {
	GetGroup(id int) (Group, error)
	GetGroups() ([]Group, error)
	CreateGroup(g *Group) (Group, error)
	UpdateGroup(g *Group) (Group, error)
	DeleteGroup(id int) error
}

type EventStore interface {
	GetEvent(id int) (Event, error)
	GetEventsByUser() ([]Event, error)
	GetEventsByGroup(id int) ([]Event, error)
	CreateEvent(e *Event) (Event, error)
	UpdateEvent(e *Event) (Event, error)
	DeleteEvent(id int) error
}

type BookStore interface {
	GetBook(id int) (Book, error)
	GetBooks() ([]Book, error)
	GetBooksByCategory() ([]Book, error)
	CreateBook(c *Book) (Book, error)
	UpdateBook(c *Book) (Book, error)
	DeleteBook(id int) error
}

type ChapterStore interface {
	GetChapter(id int) (Chapter, error)
	GetChaptersByBook(id int) ([]Chapter, error)
	CreateChapter(c *Chapter) (Chapter, error)
	UpdateChapter(c *Chapter) (Chapter, error)
	DeleteChapter(id int) error
}

type SectionStore interface {
	GetSection(id int) (Section, error)
	GetSectionsByChapter(id int) ([]Section, error)
	CreateSection(s *Section) (Section, error)
	UpdateSection(s *Section) (Section, error)
	DeleteSection(id int) error
}
