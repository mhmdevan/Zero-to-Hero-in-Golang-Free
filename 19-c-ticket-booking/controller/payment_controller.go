package controllers

import (
	"net/http"
	"os"
	"ticket-booking/database"
	"ticket-booking/models"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/charge"
)

func InitStripe() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY") // Set your Stripe Secret Key as an environment variable
}

func ProcessPayment(c *gin.Context) {
	var paymentRequest struct {
		UserID   uint   `json:"user_id" binding:"required"`
		EventID  uint   `json:"event_id" binding:"required"`
		Amount   int64  `json:"amount" binding:"required"`   // Amount in cents
		Currency string `json:"currency" binding:"required"` // e.g., "usd"
		Token    string `json:"token" binding:"required"`    // Stripe token
	}

	if err := c.ShouldBindJSON(&paymentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a charge using the Stripe API
	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(paymentRequest.Amount),
		Currency:    stripe.String(paymentRequest.Currency),
		Description: stripe.String("Ticket Booking Payment"),
		Source:      &stripe.SourceParams{Token: stripe.String(paymentRequest.Token)},
	}
	ch, err := charge.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Payment failed: " + err.Error()})
		return
	}

	// Save the payment and update booking status in the database
	booking := models.Booking{
		UserID:  paymentRequest.UserID,
		EventID: paymentRequest.EventID,
		Status:  "Paid",
	}
	if err := database.DB.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment successful", "charge_id": ch.ID})
}
