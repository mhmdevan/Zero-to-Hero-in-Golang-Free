package controllers

import (
	"net/http"
	"ticket-booking/database"
	"ticket-booking/models"

	"github.com/gin-gonic/gin"
)

func ListEvents(c *gin.Context) {
	var events []models.Event
	if err := database.DB.Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch events"})
		return
	}
	c.JSON(http.StatusOK, events)
}
