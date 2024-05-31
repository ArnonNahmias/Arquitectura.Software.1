package dao

import (
	"time"
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Email        string `gorm:"unique"`
	PasswordHash string
	Type         string
	CreationDate time.Time
	LastUpdated  time.Time
}

type Course struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Description  string
	CreationDate time.Time
	LastUpdated  time.Time
}

type Subscription struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uint
	CourseID     uint
	CreationDate time.Time
	LastUpdated  time.Time
}
