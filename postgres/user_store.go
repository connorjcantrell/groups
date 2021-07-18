package postgres

import (
	"fmt"

	"github.com/connorjcantrell/groups"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserStore struct {
	*sqlx.DB
}

func (s *UserStore) User(id uuid.UUID) (groups.User, error) {
	var u groups.User
	if err := s.Get(&u, `SELECT * FROM users WHERE id = $1`, id); err != nil {
		return groups.User{}, fmt.Errorf("error getting user: %w", err)
	}
	return u, nil
}

func (s *UserStore) UserByEmail(email string) (groups.User, error) {
	var u groups.User
	if err := s.Get(&u, `SELECT * FROM users WHERE username = $1`, email); err != nil {
		return groups.User{}, fmt.Errorf("error getting user: %w", err)
	}
	return u, nil
}

func (s *UserStore) Users() ([]groups.User, error) {
	var uu []groups.User
	if err := s.Select(&uu, `SELECT * FROM users`); err != nil {
		return []groups.User{}, fmt.Errorf("error getting users: %w", err)
	}
	return uu, nil
}

func (s *UserStore) CreateUser(u *groups.User) (groups.User, error) {
	if err := s.Get(u, `INSERT INTO users VALUES ($1, $2, $3) RETURNING *`,
		u.ID,
		u.Email,
		u.Password); err != nil {
		return groups.User{}, fmt.Errorf("error creating user: %w", err)
	}
	// TODO: Return User
	return groups.User{}, nil
}

func (s *UserStore) UpdateUser(u *groups.User) error {
	if err := s.Get(u, `UPDATE users SET username = $1, password = $2 WHERE id = $3 RETURNING *`,
		u.Email,
		u.Password,
		u.ID); err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}

func (s *UserStore) DeleteUser(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM users WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}
