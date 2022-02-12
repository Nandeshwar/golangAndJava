package dto

import (
	"time"
)

type Event struct {
	Id          int       `json:"id"`
	Day         int       `json:"day" validate:"required,gte=1,lte=31"`
	Month       int       `json:"month"`
	Year        int       `json:"year"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
