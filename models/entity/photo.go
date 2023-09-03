package entity

import "time"

type Photo struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"type:VARCHAR(255)" json:"title"`
	Caption   string    `gorm:"type:TEXT" json:"caption"`
	PhotoUrl  string    `gorm:"type:VARCHAR(255)" json:"photoUrl"`
	UserID    int64     `json:"userId"`            // Ini akan menjadi foreign key
	User      User      `gorm:"foreignKey:UserID"` // Hubungan ke tabel "user"
	CreatedAt time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:DATETIME;default:NULL" json:"updatedAt"`
}

type PhotoResponse struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photoUrl"`
	UserID    int64     `json:"userId"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
