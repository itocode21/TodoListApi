package models

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type User struct {
	Name     string
	Email    string
	Password string
}
