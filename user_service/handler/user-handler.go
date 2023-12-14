package handler

import (
	"carpooling-platform/user_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Implement APIs for user registration, profile update, account deletion, and past trip retrieval

func RegisterUser(c *gin.Context) {
	var user user_service.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert user into the database (replace with your database query)
	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func UpdateUser(c *gin.Context) {
	userID := c.Param("id")

	// Update user in the database (replace with your database query)
	result := db.Model(&user_service.User{}).Where("id = ?", userID).Updates(&user_service.User{
		FirstName: c.PostForm("FirstName"),
		LastName:  c.PostForm("LastName"),
		Mobile:    c.PostForm("Mobile"),
		Email:     c.PostForm("Email"),
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User profile updated successfully"})
}

func DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	// Delete user from the database (replace with your database query)
	result := db.Where("id = ?", userID).Delete(&user_service.User{})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user account"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User account deleted successfully"})
}

func GetUserTrips(c *gin.Context) {
	userID := c.Param("id")

	// Retrieve past trips from the database (replace with your database query)
	// ...

	// Example response
	trips := []string{"Trip1", "Trip2", "Trip3"}
	c.JSON(http.StatusOK, gin.H{"trips": trips})
}
