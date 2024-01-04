package models

type User struct {
	ID        string
	Email     string
	Username  string
	Password  string
	PhotoURL  string
	FirstName string
	LastName  string
	IsOnline  bool
	IsBanned  bool
}
