# Outbox Pattern Example (Go)

This project demonstrates the Outbox Pattern using Go, Gin, GORM, and PostgreSQL. It is split into two main components:

- **API Service** (`api/`): Handles pizza order creation via a REST API and writes events to an outbox table.
- **Job Service** (`job/`): Periodically processes outbox events, simulating payment and delivery, and updates their status.

## Features

- **Pizza Ordering API**: Submit pizza orders via HTTP POST requests.
- **Outbox Table**: Ensures reliable event publishing by writing order events to a dedicated outbox table in the database.
- **Background Job**: Polls the outbox table for new events, processes them, and marks them as completed.
- **PostgreSQL Database**: Stores orders and outbox events.

## Project Structure

```
api/         # REST API for pizza orders
  main.go
  service.go
  db/
  service/
  types.go
  docker-compose.yml
job/         # Background job for processing outbox events
  main.go
  service/
  db/
```

## How It Works

1. **Order Creation**: The API receives a pizza order and stores it in the database. It also writes an event to the outbox table.
2. **Outbox Processing**: The job service polls the outbox table for pending events, processes each (e.g., simulates payment and delivery), and updates their status.

## Running the Project

1. **Start PostgreSQL**

   In the `api/` directory:
   ```sh
   docker compose up -d
   ```

2. **Run the API Service**

   In the `api/` directory:
   ```sh
   go run .
   ```

   The API will be available at `http://localhost:8080`.

3. **Create a Pizza Order**

   Example request:
   ```sh
   curl -X POST http://localhost:8080/pizza \
     -H 'Content-Type: application/json' \
     -d '{"flavor":"Pepperoni","size":"Large","quantity":2,"price":19.99,"address":"123 Main St","user_name":"Alice"}'
   ```

4. **Run the Job Service**

   In the `job/` directory:
   ```sh
   go run .
   ```

   The job will log processing of pizza orders from the outbox.

## Technologies Used
- Go
- Gin (API framework)
- GORM (ORM)
- PostgreSQL
- Docker Compose

## Outbox Pattern
This pattern is used to reliably publish events or messages as part of a database transaction, ensuring that no events are lost even if a service crashes after writing to the database.

---

Feel free to explore and extend this project for your own event-driven or microservices architectures!
