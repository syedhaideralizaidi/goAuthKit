# AuthKit

AuthKit is a Go-based application providing user authentication features, including email verification, password reset, and JWT-based authentication. It features robust logging for authenticated API requests and ensures each request to authenticated routes is properly authenticated.

## Features

- User Registration with Email Verification
- Password Reset with Email Token
- JWT Authentication
- Logging for Authenticated API Requests
- Docker-based PostgreSQL setup
- Migration management with `migrate` tool
- SQL code generation with `sqlc`

## Prerequisites

Before running the project, ensure you have the following installed:

- Go (1.18+)
- Docker
- Docker Compose (optional, for more complex setups)
- `migrate` CLI tool
- `sqlc` CLI tool

## Installation

1. **Clone the Repository**

   ```sh
   git clone https://github.com/yourusername/authkit.git
   cd authkit

2. **Set Up Environment Variables**
   ```sh
   POSTGRES_USER=your_postgres_user
   POSTGRES_PASSWORD=your_postgres_password
   POSTGRES_CONTAINER_NAME=authkit_postgres
   DATABASE_NAME=authkit_db
   DATABASE_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(DATABASE_NAME)?sslmode=disable

## Configuration

1. **Start PostgreSQL Container**
   ```sh
   make postgres_up
2. **Create the Database**
   ```sh
   make createdb
3. **Apply Migrations**
   ```sh
   make migrate_up
4. **Run the Application**
   ```sh
   make run

**Contributing**
Feel free to fork the repository and submit pull requests. For bug reports or feature requests, please open an issue on GitHub.
