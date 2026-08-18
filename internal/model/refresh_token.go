package model

import "time"

type RefreshToken struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Token     string    `gorm:"unique;not null" json:"-"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time
}
