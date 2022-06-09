package types

import (
	"time"
)

// User is a representation of a User in the database
type User struct {
	Id        string
	FirstName string
	LastName  string
	Nickname  string
	Password  string
	Email     string
	Country   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
