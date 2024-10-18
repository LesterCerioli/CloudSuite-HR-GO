# Cloudsuite-HR-GO

This API was developed in Go using the Fiber framework to manage time records. It allows the creation, viewing, and filtering of time records by date, year, month, and day.

## Technologies Used

- **Golang**: Main programming language.
- **Fiber**: Fast and minimalist web framework for Go.
- **GORM**: ORM (Object Relational Mapping) for database manipulation.
- **PostgreSQL**: Database used.
- **Docker**: Containerization for development and production environments.
- **UUID**: Generation of unique identifiers.

## Features

The API allows:

- Create a new time record.
- Get all time records.
- Get time records filtered by date, year, month, or day.

## Docker 

- **Terminal Path**: CloudSuite-HR-GO\cloudsuite-hr-api
- **Command to start docker**: docker-compose up -d

<img src="./img/docker_comando_init.png" alt="Command to start docker">

- **Docker Images**

<img src="./img/docker_img.png" alt="API and Database images in docker">

- **Docker Container**

<img src="./img/docker_container.png" alt="Docker Container">

## Migration

- **Terminal Path**: CloudSuite-HR-GO\cloudsuite-hr-api\src
- **Command to create the times table**: migrate -path ./migrations -database "postgres://postgres:lucastechnology@localhost:5432/cloudsuite_hr_db?sslmode=disable" up

<img src="./img/migration_comand_init.png" alt="Command to create the times table">

- **Table View in DBeaver**

<img src="./img/migration_table_times.png" alt="Times table view">

## Routes

	Post("/times") Create a new time record.
	Get("/times") Get all time records.
	Get("/times/date/:date") Get time records filtered by date.
	Get("/times/year/:year") Get time records filtered by year.
	Get("/times/month/:month") Get time records filtered by month.
	Get("/times/day/:day") Get time records filtered by day.

### Create

- **Route**: `POST /times`
- **Description**: Creates a new time record.
- **Body Example**:
  ```json
  {
    "date": "2024-10-18",
    "entry_time": "2024-10-18T08:00:00Z",
    "lunch_entry_time": "2024-10-18T12:00:00Z",
    "lunch_exit_time": "2024-10-18T13:00:00Z",
    "exit_time": "2024-10-18T17:00:00Z"
  }