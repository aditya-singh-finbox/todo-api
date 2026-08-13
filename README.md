# Todo API

A simple Go REST API for managing todos, built with Gin and GORM.

## Features

- CRUD endpoints for todos
- PostgreSQL database integration
- Environment-based configuration
- Docker Compose setup for local database

## Tech Stack

- Go
- Gin
- GORM
- PostgreSQL
- Docker Compose

## Project Structure

- `cmd/server` - application entry point
- `config` - environment and configuration loading
- `internal/handler` - HTTP handlers
- `internal/repository` - database access logic
- `internal/service` - business logic
- `internal/model` - data models
- `internal/routes` - route definitions
- `internal/database` - database connection setup

## Prerequisites

- Go 1.22+
- Docker and Docker Compose
- PostgreSQL client (optional)

## Setup

1. Copy the example environment file:

   ```bash
   cp .env.example .env
   ```

2. Update the values in `.env` if needed.

3. Start PostgreSQL with Docker Compose:

   ```bash
   docker compose up -d
   ```

4. Run the API:

   ```bash
   go run ./cmd/server
   ```



## API Endpoints

### Todos

- `GET /todos` - Get all todos
- `GET /todos/:id` - Get a todo by ID
- `POST /todos` - Create a todo
- `PUT /todos/:id` - Update a todo
- `DELETE /todos/:id` - Delete a todo

## Notes

- The server runs on port `8080` by default unless `PORT` is set.
- The PostgreSQL container is exposed on port `5433` to avoid conflicts with local PostgreSQL installations.

## License

This project is for learning and development purposes.
