package users

import (
	"context"
	"wiph-api/store/postgres"
)

// UserOps ...
type UserOps struct {
	Pg *postgres.Conn
}

// GetUsers returns all user data
func (u *UserOps) GetUsers() ([]*User, error) {
	rows, err := u.Pg.Queries.QueryContext(context.Background(), u.Pg, "get-users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]*User, 0)
	for rows.Next() {
		user := &User{}
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, rows.Err()
}
