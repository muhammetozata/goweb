package models

type User struct {
	ID       int
	Username string
	Name     string
	Email    string
	Phone    string
	Website  string
	Todos    []Todo
}
