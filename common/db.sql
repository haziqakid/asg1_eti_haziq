-- common/setup.sql

-- Create the users table
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    mobile VARCHAR(20) NOT NULL,
    email VARCHAR(255) NOT NULL
);

-- Create the car_owners table
CREATE TABLE IF NOT EXISTS car_owners (
    user_id INT PRIMARY KEY,
    driver_license VARCHAR(20) NOT NULL,
    car_plate_number VARCHAR(20) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Create the trips table
CREATE TABLE IF NOT EXISTS trips (
    id INT AUTO_INCREMENT PRIMARY KEY,
    car_owner_id INT NOT NULL,
    pickup_location VARCHAR(255) NOT NULL,
    alt_pickup_locations VARCHAR(255),
    start_time DATETIME NOT NULL,
    destination VARCHAR(255) NOT NULL,
    available_seats INT NOT NULL,
    FOREIGN KEY (car_owner_id) REFERENCES car_owners(user_id)
);

-- Create the trip_enrollments table
CREATE TABLE IF NOT EXISTS trip_enrollments (
    trip_id INT,
    user_id INT,
    PRIMARY KEY (trip_id, user_id),
    FOREIGN KEY (trip_id) REFERENCES trips(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
