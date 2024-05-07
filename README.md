# Stock Management API

This is a Go application for managing stocks in a PostgreSQL database. It provides functionalities to Create, Read, Update, and Delete (CRUD) stock entries through a RESTful API.

## Features:

- Create new stock entries.
- Retrieve a stock by its ID.
- Fetch all stocks from the database.
- Update existing stock entries.
- Delete stock entries by ID.

## Technology Stack:

- Go Programming Language 
- PostgreSQL Database 
- Gorilla Mux 
- go-postgres  
- godotenv 

## Installation:

**Prerequisites:**

- Go installed on your system 
- PostgreSQL database server running

**Clone the repository:**

```bash
    https://github.com/Shreyank031/go-postgres.git
```

**Install dependencies:**

```bash
    cd go-postgres
    go mod download
```

**Configuration:**

- Create a .env file in the project root directory.

- Add the following environment variable to the .env file, replacing `<username>` with your postgres username, `<password>` with your postgres password of the mentioned user and `<db_name>` with name of database with your actual PostgreSQL connection string:

```bash

    POSTGRES_URL="postgres://<username>:<password>@localhost:5432/<db_name>?sslmode=disable"
```

## Running the API:

```bash
    go run main.go
```

This will start the API server on port 8080 by default

### API Endpoints:

- `/api/stock` = `GET` method, Fetches all stocks from the database.
- `/api/stock/{id}` = `GET` method, Retrieves a stock entry by its ID.
- `/api/stock{id}` = `PUT` method, Updates an existing stock entry based on the request body and ID.
- `/api/newstock{id}` = `POST` method, Creates a new stock entry based on the request body (JSON).
- `/api/deletestock{id}` = `DELETE` method, Deletes a stock entry by its ID.

**Request and Response Format:**

Requests should be sent in JSON format. Responses are also in JSON format

### Example Usage (using CURL):

**Create a new stock:**

```bash
    curl localhost:8080/api/newstock/1 --include --header "Content-Type: application/json" -d '{
  "name": "Apple Inc.",
  "price": 182,
  "company": "AAPL"
}' --request "POST"

```
**Response**

```bash
{
  "ID": 1,
  "Message": "Stock created successfully"
}
```

