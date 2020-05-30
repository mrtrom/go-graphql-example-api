package model

import "time"

type Chat struct {
	ID        uint `gorm:"primary_key"`
	From      string
	Content   string
	CreatedAt time.Time `gorm:"column:created_at"`
}
