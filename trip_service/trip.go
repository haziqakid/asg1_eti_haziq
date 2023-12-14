// trip_service/trip.go

package trip_service

import "time"

// Define the Trip and TripEnrollment structs

type Trip struct {
	ID                 int       `json:"id" gorm:"primaryKey"`
	CarOwnerID         int       `json:"car_owner_id" binding:"required"`
	PickupLocation     string    `json:"pickup_location" binding:"required"`
	AltPickupLocations string    `json:"alt_pickup_locations"`
	StartTime          time.Time `json:"start_time" binding:"required"`
	Destination        string    `json:"destination" binding:"required"`
	AvailableSeats     int       `json:"available_seats" binding:"required"`
}

type TripEnrollment struct {
	TripID int `json:"trip_id" binding:"required"`
	UserID int `json:"user_id" binding:"required"`
}
