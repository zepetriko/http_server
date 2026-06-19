# Chirpy

Chirpy is a Twitter-inspired REST API built in Go. It supports user authentication, chirp creation and management, refresh tokens, webhooks, and PostgreSQL persistence.

The project was built to explore backend development concepts including authentication, authorization, database design, migrations, and API architecture.

## Features

### Authentication & Users

* User registration
* User login
* JWT access tokens
* Refresh tokens
* Token revocation
* Password hashing with bcrypt
* User profile updates
* Chirpy Red membership upgrades via webhooks

### Chirps

* Create chirps
* Retrieve all chirps
* Retrieve a single chirp by ID
* Delete chirps
* Filter chirps by author
* Sort chirps by creation date (ascending or descending)
* Chirp validation and profanity replacement

### Administration

* Health check endpoint
* Request metrics tracking
* Metrics reset endpoint (development/admin)

### Database

* PostgreSQL persistence
* SQLC-generated queries
* Goose database migrations

## Tech Stack

* Go
* PostgreSQL
* SQLC
* Goose
* JWT
* bcrypt
* net/http

## API Overview

### Health Check

| Method | Endpoint       |
| ------ | -------------- |
| GET    | `/api/healthz` |

Checks server health.

---

### Users

| Method | Endpoint     |
| ------ | ------------ |
| POST   | `/api/users` |
| PUT    | `/api/users` |

Create and update users.

---

### Authentication

| Method | Endpoint       |
| ------ | -------------- |
| POST   | `/api/login`   |
| POST   | `/api/refresh` |
| POST   | `/api/revoke`  |

Handles login, token refresh, and token revocation.

---

### Chirps

| Method | Endpoint                |
| ------ | ----------------------- |
| POST   | `/api/chirps`           |
| GET    | `/api/chirps`           |
| GET    | `/api/chirps/{chirpID}` |
| DELETE | `/api/chirps/{chirpID}` |

#### Query Parameters

`GET /api/chirps`

| Parameter | Values    |
| --------- | --------- |
| author_id | UUID      |
| sort      | asc, desc |

Examples:

```http
GET /api/chirps
```

```http
GET /api/chirps?sort=desc
```

```http
GET /api/chirps?author_id=<user-id>
```

```http
GET /api/chirps?author_id=<user-id>&sort=desc
```

---

### Webhooks

| Method | Endpoint              |
| ------ | --------------------- |
| POST   | `/api/polka/webhooks` |

Processes Chirpy Red upgrades from Polka.

---

### Metrics

| Method | Endpoint         |
| ------ | ---------------- |
| GET    | `/admin/metrics` |
| POST   | `/admin/reset`   |

Used for tracking and resetting API metrics.

## Getting Started

### Prerequisites

* Go 1.24+
* PostgreSQL
* SQLC
* Goose

### Installation

Clone the repository:

```bash
git clone https://github.com/<your-username>/chirpy.git
cd chirpy
```

Create a `.env` file:

```env
DB_URL=postgres://postgres:postgres@localhost:5432/chirpy?sslmode=disable
JWT_SECRET=your-secret-key
POLKA_API_KEY=your-polka-api-key
```

Run migrations:

```bash
goose postgres "$DB_URL" up
```

Generate SQLC code:

```bash
sqlc generate
```

Start the server:

```bash
go run .
```

The API will be available at:

```text
http://localhost:8080
```

## Project Structure

```text
.
├── main.go
├── internal/
│   ├── auth/
│   ├── database/
│   └── ...
├── sql/
│   ├── schema/
│   └── queries/
├── assets/
├── .env
└── README.md
```

## What I Learned

* Designing RESTful APIs
* Authentication and authorization with JWTs
* Refresh token workflows
* Secure password storage
* PostgreSQL schema design
* Database migrations
* SQLC query generation
* Webhook processing
* Go HTTP servers and middleware
* Structuring production-style Go projects

## Future Improvements

* Follow system
* Likes and reactions
* Chirp editing
* Media uploads
* Rate limiting
* API documentation with OpenAPI/Swagger
* Containerization with Docker

## License

MIT
