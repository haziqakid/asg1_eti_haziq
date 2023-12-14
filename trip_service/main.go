// trip_service/main.go

package main

import (
	"carpooling-platform/common"
	"carpooling-platform/trip_service/handler"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDatabase()

	r := gin.Default()

	// Trip publishing API
	r.POST("/trips/publish", handler.PublishTrip)

	// Trip enrollment API
	r.POST("/trips/enroll", handler.EnrollInTrip)

	// Past trip retrieval API
	r.GET("/trips/:id", handler.GetPastTrips)

	port := 8082
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Trip Microservice is running on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
