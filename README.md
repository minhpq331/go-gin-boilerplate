# Go Gin Boilerplate

A modern Go application boilerplate using Gin framework with multiple entrypoints and environment configuration.

## Features

- Multiple entrypoints (API server and worker)
- Environment configuration using godotenv
- Clean project structure
- Latest Go version (1.22)
- Gin web framework
- Health check endpoint
- Dependency injection using Google Wire

## Project Structure

```
.
├── cmd/
│   ├── api/        # API server entrypoint
│   └── worker/     # Worker entrypoint
├── internal/
│   ├── app/
│   │   ├── api/
│   │   │   ├── controller/
│   │   │   ├── router/
│   │   │   └── middleware/
│   │   │   └── binding/
│   │   └── worker/
│   │       └── handler/
│   ├── config/     # Configuration package
│   ├── database/   # Database connections
│   ├── eventbus/   # Event bus
│   ├── logger/     # Logging
│   ├── repository/ # Data access
│   ├── service/    # Business logic
│   └── model/      # Model
├── .env            # Environment variables
├── go.mod          # Go module file
└── README.md       # This file
```

## Getting Started

1. Clone the repository
2. Copy `.env.example` to `.env` and adjust the values
3. Install necessary tools:
   ```bash
   go install github.com/githubnemo/CompileDaemon
   go install github.com/google/wire/cmd/wire
   ```

   Edit `.zshrc` or `.bashrc` to add go bin path to PATH:
   ```bash
   export PATH="$PATH:$(go env GOPATH)/bin"
   ```
4. Install dependencies:
   ```bash
   go mod tidy
   ```
5. Run the API server with hot reloading:
   ```bash
   ./dev.sh api
   ```
6. Run the worker with hot reloading:
   ```bash
   ./dev.sh worker
   ```

## API Endpoints

- `GET /health` - Health check endpoint with uptime and timestamp
- `POST /api/v1/notes` - Create a new note
- `GET /api/v1/notes/:id` - Get a note by ID
- `GET /api/v1/notes` - Get all notes
- `PUT /api/v1/notes/:id` - Update a note by ID
- `DELETE /api/v1/notes/:id` - Delete a note by ID

## Environment Variables

See `.env.example` file for available configuration options.

## Dependency Injection

This project uses Google Wire for dependency injection. The wire configuration is in `cmd/api/wire.go` and `cmd/worker/wire.go`. To regenerate the wire_gen.go file after making changes:

```bash
wire ./cmd/api/
wire ./cmd/worker/
```

## Features

- [x] Multiple entrypoints (API server and worker)
- [x] API server with Gin web framework
- [x] Worker with Asynq
- [x] Environment configuration using godotenv
- [x] Dependency injection using Google Wire
- [x] Event-driven processing with EventBus
- [x] Logger
- [x] Error handling
- [x] Validation
- [x] Hot reloading
- [x] Redis
- [x] MongoDB
- [x] Docker

## License

MIT
