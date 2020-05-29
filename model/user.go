package model

type User struct {
	// gorm.Model
	ID       uint `gorm:"primary_key"`
	Name     string
	LastName string
}
