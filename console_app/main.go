package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "http://localhost:"

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
}

type CarOwner struct {
	DriverLicense  string `json:"driver_license"`
	CarPlateNumber string `json:"car_plate_number"`
}

type Trip struct {
	PickupLocation string `json:"pickup_location"`
	// Additional trip details...
}

type Enrollment struct {
	TripID int `json:"trip_id"`
	// Additional enrollment details...
}

func main() {
	for {
		printMenu()

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			registerUser()
		case 2:
			publishTrip()
		case 3:
			enrollInTrip()
		case 4:
			getPastTrips()
		case 5:
			fmt.Println("Exiting console application.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func printMenu() {
	fmt.Println("\nCarpooling Platform Console Application")
	fmt.Println("1. Register User")
	fmt.Println("2. Publish Trip")
	fmt.Println("3. Enroll in Trip")
	fmt.Println("4. Get Past Trips")
	fmt.Println("5. Exit")
	fmt.Print("Enter your choice: ")
}

func registerUser() {
	var user User

	fmt.Print("Enter First Name: ")
	fmt.Scanln(&user.FirstName)

	fmt.Print("Enter Last Name: ")
	fmt.Scanln(&user.LastName)

	fmt.Print("Enter Mobile Number: ")
	fmt.Scanln(&user.Mobile)

	fmt.Print("Enter Email Address: ")
	fmt.Scanln(&user.Email)

	userJSON, _ := json.Marshal(user)
	makeRequest("POST", "/users/register", userJSON)
}

func publishTrip() {
	var trip Trip

	// Get car owner details
	var carOwner CarOwner
	fmt.Print("Are you a car owner? (yes/no): ")
	fmt.Scanln(&carOwner)

	if carOwner.DriverLicense != "" && carOwner.CarPlateNumber != "" {
		fmt.Print("Enter Driver's License Number: ")
		fmt.Scanln(&carOwner.DriverLicense)

		fmt.Print("Enter Car Plate Number: ")
		fmt.Scanln(&carOwner.CarPlateNumber)
	}

	// Get trip details
	fmt.Print("Enter Pickup Location: ")
	fmt.Scanln(&trip.PickupLocation)

	// Additional trip details...

	tripJSON, _ := json.Marshal(trip)
	makeRequest("POST", "/trips/publish", tripJSON)
}

func enrollInTrip() {
	var enrollment Enrollment

	fmt.Print("Enter Trip ID to Enroll: ")
	fmt.Scanln(&enrollment.TripID)

	// Additional enrollment details...

	enrollmentJSON, _ := json.Marshal(enrollment)
	makeRequest("POST", "/trips/enroll", enrollmentJSON)
}

func getPastTrips() {
	var userID int
	fmt.Print("Enter User ID: ")
	fmt.Scanln(&userID)

	makeRequest("GET", fmt.Sprintf("/users/%d/trips", userID), nil)
}

func makeRequest(method, endpoint string, payload []byte) {
	url := baseURL + "8081" + endpoint
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("HTTP Status Code:", resp.StatusCode)
	// Read and print the response body if needed
	// ...

	fmt.Println("Request processed successfully.")
}
