// user_service/user.go

package main

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
}

type Passenger struct {
	UserID int `json:"user_id"`
	// Additional fields for passenger profile
}

type CarOwner struct {
	UserID         int    `json:"user_id"`
	DriverLicense  string `json:"driver_license"`
	CarPlateNumber string `json:"car_plate_number"`
	// Additional fields for car owner profile
}
