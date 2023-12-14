// user_service/main.go

package main

import (
	"carpooling-platform/common"
	"carpooling-platform/user_service/handler"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDatabase()

	r := gin.Default()

	// User registration API
	r.POST("/users/register", handler.RegisterUser)

	// User profile update API
	r.PUT("/users/:id/update", handler.UpdateUser)

	// Account deletion API
	r.DELETE("/users/:id/delete", handler.DeleteUser)

	// Past trip retrieval API
	r.GET("/users/:id/trips", handler.GetUserTrips)

	port := 8081
	addr := fmt.Sprintf(":%d", port)
	log.Printf("User Microservice is running on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
