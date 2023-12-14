# asg1_eti_haziq

# Carpooling Platform

Welcome to the Carpooling Platform project! This platform allows users to participate in a carpooling service, offering features for user registration, trip publishing, enrollment, and past trip retrieval.

## Project Structure

The project is organized into microservices for user and trip management, a common module for database interaction, a console application for interaction, and an optional web frontend.

carpooling-platform/
|--main.go
|-- user_service/
| |-- main.go
| |-- user.go
| |-- handler/
| |-- user_handler.go
|-- trip_service/
| |-- main.go
| |-- trip.go
| |-- handler/
| |-- trip_handler.go
|-- common/
| |-- database.go
  |-- setup.sql
|-- console_app/
| |-- main.go
|-- web_frontend/ (optional, only if implementing bonus)
| |-- main.go
| |-- static/
| |-- index.html
|-- README.md
|-- docker-compose.yml


## Getting Started

Follow these steps to set up the Carpooling Platform environment and run the project:

### Prerequisites

- [Go](https://golang.org/) installed on your machine.
- [Docker](https://www.docker.com/) installed for running the MySQL database.
- [Git](https://git-scm.com/) for cloning the repository.

### Setting Up the Environment


cd carpooling-platform
Database Setup
Open a terminal and run the following command to start the MySQL server Docker container:
go run main.go

Running Microservices
Start the User Service:

cd user_service
go run main.go
Start the Trip Service:

cd trip_service
go run main.go
Running Console Application
Start the Console Application:

cd console_app
go run main.go
Follow the prompts in the console to interact with the Carpooling Platform.

Running Web Frontend (Optional)
Start the Web Frontend:

cd web_frontend
go run main.go
Access the web frontend at http://localhost:8080 in your browser.