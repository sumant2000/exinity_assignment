
Payment Gateway Service

Overview

The Payment Gateway Service is a scalable and extensible service that processes deposit and withdrawal transactions through multiple payment gateways. The architecture is designed to accommodate future gateway integrations, support different protocols (like HTTP and ISO8583), and handle various data formats (JSON, XML, etc.).

Architecture

	•	Gateways: Each payment gateway is implemented as a separate struct that adheres to the Gateway interface. This ensures that new gateways can be easily integrated by implementing the Deposit and Withdraw methods.
	•	Gateway A: Uses HTTP and JSON.
	•	Gateway B: Uses HTTP and SOAP/XML.
	•	ISO8583 Gateway: (Example for future extensibility) Uses TCP for ISO8583 message format.
	•	Service Layer: The service layer (TransactionService) abstracts the gateway logic and provides a unified API for processing deposits and withdrawals. The service delegates requests to the appropriate gateway based on user input.
	•	Resilience: The system is designed with resilience features such as:
	•	Circuit Breakers: Prevents overloading gateways with requests when they are failing.
	•	Retries: Attempts to retry failed transactions before returning an error.
	•	Timeouts: Ensures that requests to gateways don’t hang indefinitely.
	•	API:
	•	/deposit: Handles deposit transactions through different gateways.
	•	/withdraw: Handles withdrawal transactions through different gateways.
	•	/health: Health check endpoint for monitoring the service status.

Key Components

	•	Gateways: Each gateway (A, B, etc.) implements the Gateway interface. Future gateways (e.g., ISO8583 over TCP) can be added easily.
	•	TransactionService: Encapsulates the business logic for processing transactions and interacts with the gateways.
	•	Handlers: Exposes HTTP endpoints for deposits, withdrawals, and health checks.
	•	Resilience: Includes circuit breakers, retries, and timeouts to ensure system reliability.

Build, Run, and Test Instructions

Prerequisites

	•	Go: Version 1.23.2 or later
	•	Docker (Optional): If you wish to use Docker for running the service in a container.
	•	Postman or curl: For testing the API endpoints.

1. Build the Service

First, ensure you are in the project directory and initialize Go modules:

go mod tidy

To build the Go application:

go build

2. Run the Service

Run the application locally on localhost:8080:

go run main.go

You should see output like:

[GIN-debug] Listening and serving HTTP on :8080

3. API Endpoints

You can test the following endpoints using Postman, curl, or other HTTP clients.

3.1 Deposit Funds

	•	URL: http://localhost:8080/deposit
	•	Method: POST
	•	Body (JSON):

{
"amount": 100.0,
"accountID": "123456",
"gateway": "A"
}


	•	Response (JSON):

{
"transactionID": "transaction_id_a"
}



3.2 Withdraw Funds

	•	URL: http://localhost:8080/withdraw
	•	Method: POST
	•	Body (JSON):

{
"amount": 50.0,
"accountID": "123456",
"gateway": "B"
}


	•	Response (JSON):

{
"transactionID": "withdrawal_id_b"
}



3.3 Health Check

	•	URL: http://localhost:8080/health
	•	Method: GET
	•	Response (JSON):

{
"status": "OK"
}



4. Running Tests

To run unit and integration tests, use the following command:

go test ./...

This will run all tests in the services, handlers, and other packages. You should see output indicating the number of tests that passed or failed.

5. Docker (Optional)

If you want to run the service inside a Docker container, you can use the provided Dockerfile.

Build the Docker Image

docker build -t payment-gateway-service .

Run the Docker Container

docker run -p 8080:8080 payment-gateway-service

The service will now be accessible at http://localhost:8080.
