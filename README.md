# Backend

A Go backend service built with the **Fiber framework** and **PostgreSQL**, implementing a user management API with **date of birth (DOB)** and **dynamic age calculation**.

---

## Features

- RESTful API using GoFiber
- PostgreSQL database with SQLC for type-safe queries
- Dynamic age calculation using Go's `time` package
- Input validation using `go-playground/validator`
- Structured logging using Uber Zap
- Clean layered architecture (handler → service → database)

---

## Prerequisites

- Go 1.21 or higher
- PostgreSQL
- Git

---

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/Sudhanva05/Backend.git
cd Backend
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Set up environment variables

Create a `.env` file in the root directory:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=backend_db
```

### 4. Database setup

Create the database:

```sql
CREATE DATABASE backend_db;
```

Run SQLC code generation:

```bash
sqlc generate
```

---

## Running the Server

Start the development server:

```bash
go run cmd/server/main.go
```

The server runs at: `http://127.0.0.1:3000`

---

## API Endpoints

### Create User

**POST** `/users`

```json
{
  "name": "Alice",
  "dob": "1990-05-10"
}
```

### Get All Users

**GET** `/users`

### Get User by ID

**GET** `/users/:id`

Each response includes a dynamically calculated age.

---

## Project Structure

```
Backend/
├── cmd/
│   └── server/
│       └── main.go
├── db/
│   ├── migrations/
│   └── sqlc/
├── internal/
│   ├── handler/
│   ├── service/
│   ├── routes/
│   ├── models/
│   └── logger/
├── go.mod
├── go.sum
└── README.md
```

---

## Development

### Build

```bash
go build -o bin/server cmd/server/main.go
```

### Run Tests

```bash
go test ./...
```

---

## License

Add your license information here.
