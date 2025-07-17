# PDF Generator Service

A simple Go microservice that generates a PDF report for a given student ID.

### Author
- Moses Idowu

## Overview

This service exposes a single HTTP endpoint to generate a PDF document containing a student's details. It works by:

1.  Receiving a request with a student ID.
2.  Calling an external student information service to fetch the student's data.
3.  Using the fetched data to generate a PDF file on-the-fly.
4.  Returning the generated PDF as an attachment in the HTTP response.

The project follows a clean architecture pattern, separating concerns into distinct layers: handlers (API), services (business logic), and clients (external communication).

## Project Structure

The project is organized into several packages, each with a specific responsibility:

```
.
├── cmd/
│   └── main.go              # Main application entry point. Initializes and starts the server.
├── client/
│   └── student_client.go    # Client for fetching student data from an external service.
├── config/
│   └── config.go            # Configuration management (e.g., API keys, URLs).
├── handler/
│   ├── student_handler.go      # HTTP handlers for the API endpoints.
│   └── student_handler_test.go # Tests for the handlers.
├── model/
│   └── student.go           # Defines the Student data model.
├── service/
│   ├── pdf_generator_service.go # Logic for creating the PDF document.
│   ├── student_service.go      # Business logic to orchestrate report generation.
│   └── student_service_test.go # Tests for the service layer.
└── go.mod                   # Go module definitions.
```

You can use a tool like `curl` to test the endpoint:

```sh
curl -o report.pdf http://localhost:4500/api/v1/students/12345/report
```

This will download the generated PDF report for the student with ID `123` and save it as `report.pdf`.

## Getting Started

### Running the Service

1.  Clone the repository and navigate into the project directory.
2.  Install dependencies:
    ```sh
    go mod tidy
    ```
3.  Run the application:
    ```sh
    go run ./cmd/main.go
    ```
    The service will start and listen on `http://localhost:4500`.

## Running Tests

To run the unit tests for the project, execute the following command from the root directory:

```sh
go test -v ./...
```