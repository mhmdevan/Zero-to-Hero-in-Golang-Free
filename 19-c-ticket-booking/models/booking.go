package models

import "github.com/jinzhu/gorm"

type Booking struct {
	gorm.Model
	UserID  uint   `json:"user_id" binding:"required"`
	EventID uint   `json:"event_id" binding:"required"`
	Status  string `json:"status"` // e.g., "Booked" or "Cancelled"
}
