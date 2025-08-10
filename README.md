
# Outbox Pattern Example in Go

![Outbox pattern diagram](diagram.png)

This repository is a complete demonstration of the Outbox Pattern implemented in Go. It uses Gin for the HTTP API, GORM for ORM, PostgreSQL for persistent storage, and Redis for real-time event streaming. The project is organized into several clear components, each with a specific responsibility:

- **Client Service** (`client/`): This is a Go script that automatically generates and sends HTTP requests to the Producer API to create new pizza orders. It is useful for simulating real-world order traffic and testing the system end-to-end.
- **Producer API Service** (`producer-api/`): This is a RESTful API built with Gin. It exposes endpoints for creating pizza orders. When a new order is received, it is saved to the main business table and an event is also written to the outbox table in the same database transaction, ensuring consistency.
- **Job Service** (`job/`): This background service continuously polls the outbox table for new events (pizza orders that need to be processed). For each event, it simulates payment and delivery, then publishes the event payload to a Redis topic and marks the event as completed in the database.
- **Consumer API Service** (`consumer-api/`): This service subscribes to the Redis topic and listens for new pizza order events in real time. When a new event is received, it processes or logs the event as needed. This simulates a downstream microservice or event consumer.

## Key Features

- **Reliable Pizza Ordering API**: Accepts pizza orders via HTTP POST and ensures they are safely stored.
- **Transactional Outbox Table**: Guarantees that order events are only published if the order is successfully saved, preventing data loss or duplication.
- **Background Job for Event Processing**: Efficiently polls and processes outbox events, simulates business logic, and publishes to Redis.
- **Real-Time Event Streaming with Redis**: Uses Redis Pub/Sub to deliver events instantly to any number of consumers.
- **PostgreSQL Database**: Central storage for both business data (orders) and outbox events.

## Project Structure

Each directory contains a focused part of the system:

```
producer-api/   # REST API for pizza orders (order creation and outbox writing)
   main.go      # Entry point for the API server
   service/     # Business logic for handling pizza orders and outbox
   db/          # Database connection logic
   types.go     # Data models
   docker-compose.yml # Service-specific compose file (optional)
job/            # Background worker for processing outbox events and publishing to Redis
   main.go      # Entry point for the job service
   service/     # Outbox polling and event publishing logic
   db/          # Database connection logic
   redis/       # Redis connection logic
consumer-api/   # Service that subscribes to Redis and processes events
   main.go      # Entry point for the consumer
   service/     # Event subscription and handling logic
   redis/       # Redis connection logic
client/         # Script to generate and send pizza orders
   main.go      # Entry point for the client script
docker-compose.yml # Top-level compose file for all services
```

## How the System Works

1. **Order Creation (Producer API)**: A client (or user) sends a POST request to the `/pizza` endpoint. The Producer API saves the order to the `pizza_orders` table and, in the same transaction, writes an event to the `pizza_order_outbox` table. This ensures that no order event is lost or duplicated.
2. **Outbox Event Processing (Job Service)**: The Job service runs in the background, polling the outbox table for new events with status `pending`. For each event, it simulates payment and delivery, then publishes the event payload to the `pizza-orders` topic in Redis. After successful publishing, it marks the event as `completed` in the database.
3. **Event Consumption (Consumer API)**: The Consumer API subscribes to the `pizza-orders` topic in Redis. Whenever a new event is published, it receives the event in real time and processes or logs it. This could be extended to trigger further business logic or notifications.

## Step-by-Step: Running the Project

Follow these instructions to run the entire system locally:

1. **Start the Database and Redis**

   Open a terminal in the project root and run:
   ```sh
   docker compose up -d
   ```
   This command will start both PostgreSQL and Redis containers in the background. PostgreSQL will be available on port 5432 and Redis on port 6379.

2. **Start the Producer API Service**

   Open a new terminal, navigate to the `producer-api/` directory, and run:
   ```sh
   go run .
   ```
   This will start the API server on `http://localhost:8080`.

3. **Create a Pizza Order (using curl or the client script)**

   You can manually create an order with:
   ```sh
   curl -X POST http://localhost:8080/pizza \
     -H 'Content-Type: application/json' \
     -d '{"flavor":"Pepperoni","size":"Large","quantity":2,"price":19.99,"address":"123 Main St","user_name":"Alice"}'
   ```
   Or, to generate many orders automatically, run the client script from the `client/` directory:
   ```sh
   go run .
   ```

4. **Start the Job Service**

   In another terminal, go to the `job/` directory and run:
   ```sh
   go run .
   ```
   This service will continuously process new outbox events and publish them to Redis.

5. **Start the Consumer API Service**

   In a final terminal, navigate to the `consumer-api/` directory and run:
   ```sh
   go run .
   ```
   The consumer will print every pizza order event it receives from Redis in real time.

## Technologies Used

- **Go**: Main programming language for all services.
- **Gin**: Web framework for building the REST API.
- **GORM**: ORM for database operations.
- **PostgreSQL**: Relational database for persistent storage.
- **Redis**: In-memory data store for real-time event streaming.
- **Docker Compose**: For orchestrating multi-container environments.

## About the Outbox Pattern

The Outbox Pattern is a proven approach for reliably publishing events or messages as part of a database transaction. By writing events to an outbox table in the same transaction as the business data, you ensure that no events are lost or duplicated, even if a service crashes. A background job then reads from the outbox and publishes to a message broker (here, Redis), decoupling event production from event consumption.

---

Feel free to explore, modify, and extend this project for your own event-driven or microservices architectures. Contributions and suggestions are welcome!
