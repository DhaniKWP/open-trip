package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Phone     string    `gorm:"not null"`
	Password  string    `gorm:"not null"` // hashed password
	Role      string    `gorm:"not null"` // angler, eo, captain, admin
	Status    string    `gorm:"default:'pending'"` // pending, approved, rejected
	AvatarURL string
	CreatedAt time.Time
	UpdatedAt time.Time
}
