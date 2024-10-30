package models

type Todo struct {
	ID          int
	Title       string
	Description string
}

type User struct {
	Name     string
	Email    string
	Password string
}
