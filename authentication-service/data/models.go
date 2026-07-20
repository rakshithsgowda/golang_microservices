package data

import "time"

type User struct {
	ID        uint
	Email     string
	FirstName string
	LastName  string
	Password  string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Models struct {
	User User
}
