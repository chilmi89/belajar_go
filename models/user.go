package models

import "time"

type User struct {
	ID        int       `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	IsActive  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "users"
}
