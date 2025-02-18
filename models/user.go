package models

import "time"

type Role string

const (
	Admin Role = "admin"
	Staff Role = "staff"
)

type User struct {
	ID            int       `gorm:"primary_key;auto_increment" json:"id"`
	Name          string    `gorm:"size:255;not null" json:"name"`
	Email         string    `gorm:"size:100;not null;unique" json:"email"`
	Password_hash string    `gorm:"size:100;not null" json:"password_hash"`
	Role          Role      `gorm:"size:100;not null" json:"role"`
	Created_at    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at" `
	Updated_at    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
