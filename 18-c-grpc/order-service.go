package main

import (
	"context"
	"fmt"
	"myproject/userpb"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Order struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Item   string `json:"item"`
}

var orders = []Order{
	{ID: 1, UserID: 1, Item: "Laptop"},
	{ID: 2, UserID: 2, Item: "Smartphone"},
}

func main() {
	// Connect to UserService gRPC
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic("Failed to connect to UserService: " + err.Error())
	}
	defer conn.Close()
	userClient := userpb.NewUserServiceClient(conn)

	router := gin.Default()

	router.GET("/order", func(c *gin.Context) {
		userID := c.Query("user_id")
		if userID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}

		var userOrders []Order
		for _, order := range orders {
			if fmt.Sprintf("%d", order.UserID) == userID {
				userOrders = append(userOrders, order)
			}
		}

		if len(userOrders) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No orders found for this user"})
			return
		}

		// Fetch user info from UserService
		userIDInt, _ := strconv.Atoi(userID)
		userResp, err := userClient.GetUser(context.Background(), &userpb.UserRequest{Id: int32(userIDInt)})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user info"})
			return
		}

		response := gin.H{
			"user": gin.H{
				"id":   userResp.Id,
				"name": userResp.Name,
			},
			"orders": userOrders,
		}

		c.JSON(http.StatusOK, response)
	})

	fmt.Println("OrderService is running on http://localhost:8080")
	router.Run(":8080")
}
