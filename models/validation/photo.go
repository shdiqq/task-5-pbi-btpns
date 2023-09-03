package validation

import "time"

type PhotoRequestPost struct {
	ID        int64
	Title     string `json:"title" validate:"required"`
	Caption   string `json:"caption" validate:"required"`
	PhotoUrl  string `json:"photoUrl" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PhotoRequestPut struct {
	ID        int64
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoUrl  string `json:"photoUrl"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
