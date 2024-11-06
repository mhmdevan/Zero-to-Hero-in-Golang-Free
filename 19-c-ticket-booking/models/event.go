package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Event struct {
	gorm.Model
	Name     string  `json:"name" binding:"required"`
	Location string  `json:"location" binding:"required"`
	Tickets  int     `json:"tickets" binding:"required,gte=1"` // Must be at least 1 ticket
	Price    float64 `json:"price" binding:"required,gte=0"`   // Price must be non-negative
}

// Validate checks if the event data is valid
func (e *Event) Validate() error {
	if e.Tickets <= 0 {
		return errors.New("number of tickets must be greater than zero")
	}
	if e.Price < 0 {
		return errors.New("price must be a non-negative value")
	}
	return nil
}
