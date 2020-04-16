package users

// Repo ...
type Repo interface {
	GetUsers() ([]*User, error)
}

// User ...
type User struct {
	ID   string
	Name string
}
