package users

import "wiph-api/store/postgres"

// UserOps ...
type UserOps struct {
	Pg *postgres.Conn
}

// GetUsers returns all user data
func (u *UserOps) GetUsers() ([]*User, error) {
	return nil, nil
}
