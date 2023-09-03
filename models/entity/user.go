package entity

import "time"

type User struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"VARCHAR(255);not null" json:"username"`
	Email     string    `gorm:"VARCHAR(255);unique;not null" json:"email"`
	Password  string    `gorm:"VARCHAR(255);not null;check:CHAR_LENGTH(Password) >= 6" json:"password"`
	CreatedAt time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:DATETIME;default:NULL" json:"updatedAt"`
}

type UserResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
