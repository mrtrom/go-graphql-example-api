package model

type User struct {
	ID        string
	Email     string
	CreatedAt string `db:"created_at"`
}
