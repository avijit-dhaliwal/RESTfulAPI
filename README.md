# Go RESTful API Project

This project is a RESTful API built with Go, demonstrating proficiency in various aspects of backend development and Go programming.

## Project Overview

This API provides CRUD operations for managing items. It's built with a focus on clean architecture, scalability, and best practices in Go development.

## Features

- RESTful API endpoints for CRUD operations
- PostgreSQL database integration
- JWT-based authentication
- Structured logging
- Configuration management
- Error handling
- Middleware for authentication and logging

## Skills Demonstrated

1. **Go Programming**: Efficient use of Go's features and standard library.
2. **Web Development**: Building a RESTful API using the `net/http` package and `gorilla/mux` router.
3. **Database Integration**: Using PostgreSQL with Go's `database/sql` package.
4. **Authentication**: Implementing JWT-based authentication.
5. **Project Structure**: Organizing code into a clean, modular structure.
6. **Configuration Management**: Using Viper for flexible configuration.
7. **Error Handling**: Implementing proper error handling and returning appropriate HTTP status codes.
8. **Middleware**: Creating and using middleware for cross-cutting concerns.
9. **JSON Handling**: Working with JSON for request/response bodies.
10. **Dependency Management**: Using Go modules for dependency management.
11. **Logging**: Implementing structured logging for better observability.
12. **API Documentation**: (TODO) Using Swagger for API documentation.
13. **Testing**: (TODO) Writing unit and integration tests.

## Project Structure
.
├── cmd
│   └── server
│       └── main.go
├── internal
│   ├── api
│   │   ├── handlers.go
│   │   └── routes.go
│   ├── auth
│   │   └── auth.go
│   ├── config
│   │   └── config.go
│   ├── db
│   │   └── db.go
│   └── models
│       └── item.go
├── go.mod
├── go.sum
└── README.md

## Setup and Running

1. Clone the repository
2. Install dependencies: `go mod download`
3. Set up your PostgreSQL database
4. Create a `config.yaml` file in the root directory with your configuration
5. Run the server: `go run cmd/server/main.go`

## API Endpoints

- `GET /items`: Retrieve all items
- `POST /items`: Create a new item
- `GET /items/{id}`: Retrieve a specific item
- `PUT /items/{id}`: Update a specific item
- `DELETE /items/{id}`: Delete a specific item

## Future Improvements

- Implement comprehensive unit and integration tests
- Add Swagger documentation for the API
- Implement rate limiting
- Add caching layer
- Containerize the application using Docker

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.

## Contact

avijit.dhaliwal@gmail.com