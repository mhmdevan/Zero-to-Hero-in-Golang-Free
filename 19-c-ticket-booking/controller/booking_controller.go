package controllers

import (
	"net/http"
	"ticket-booking/database"
	"ticket-booking/models"

	"github.com/gin-gonic/gin"
)

func BookTicket(c *gin.Context) {
	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the event exists and has available tickets
	var event models.Event
	if err := database.DB.First(&event, booking.EventID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if event.Tickets <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No tickets available"})
		return
	}

	// Reduce the number of available tickets and create a booking
	event.Tickets--
	database.DB.Save(&event)

	booking.Status = "Booked"
	if err := database.DB.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to book ticket"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket booked successfully"})
}
