package db

import (
	"github.com/connorjcantrell/groups/models"
)

// func (u *User) GetCreatedAt() sql.NullTime {
// 	return u.CreatedAt
// }

// If SQLC allows you to add a nested field, skip the 3rd step.
func GetGroupWithBooksAndMembers(id int32) models.Group {
	// 1st query: select group
	groupSQLC := GetGroupByID(id)
	// 2nd query: select all nested books and users
	books := GetGroupBooksByGroup(id)
	users := GetUsersByGroup(id)
	// 3rd step: convert `sqlc group` to your `group`
	group := convertSQLCGroup(groupSQLC)
	// last step: adding books and users to group
	// myGroup.Books = users -> complier will stop you because `Books` need `[]Book`, `users` is `[]User`
	group.Books = books
	group.Members = users
	return group
}

func convertSQLCGroup(g Group) models.Group {
	return models.Group{
		ID:          int(g.ID),
		Name:        g.Name,
		Description: g.Description.String,
		CreatedAt:   g.CreatedAt.Time,
	}

}
