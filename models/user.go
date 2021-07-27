package models

type User struct {
	ID       int    `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

type UserStore interface {
	GetUser(id int) (User, error)
	GetUserByEmail(username string) (User, error)
	CreateUser(u *User) (User, error)
	UpdateUser(u *User) (User, error)
	DeleteUser(id int) error
}
