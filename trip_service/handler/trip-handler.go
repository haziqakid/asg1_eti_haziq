package handler

import (
	"carpooling-platform/trip_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Implement APIs for trip publishing, enrollment, and past trip retrieval

func PublishTrip(c *gin.Context) {
	var trip trip_service.Trip
	if err := c.ShouldBindJSON(&trip); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert trip into the database
	result := db.Create(&trip)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish trip"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Trip published successfully"})
}

func EnrollInTrip(c *gin.Context) {
	var enrollment trip_service.TripEnrollment
	if err := c.ShouldBindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert enrollment into the database
	result := db.Create(&enrollment)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enroll in trip"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Enrolled in trip successfully"})
}

func GetPastTrips(c *gin.Context) {
	// Retrieve past trips from the database (replace with your database query)
	// ...

	// Example response
	trips := []string{"PastTrip1", "PastTrip2", "PastTrip3"}
	c.JSON(http.StatusOK, gin.H{"pastTrips": trips})
}
