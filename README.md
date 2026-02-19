# Go HTTP Server

Simple HTTP server written in Go with routing and logging middleware.

## Features

- Small, dependency-free HTTP server
- Routing (defined in `router.go`)
- Request logging middleware (see `middleware.go`)
- Health check and example endpoint

## Prerequisites

- Go 1.18+ installed (GOMOD enabled)
- Docker (optional, for container builds)

## Run locally

From the project root run:

```bash
go run .
```

The server listens on port 8080 by default.

## Build

To build a binary:

```bash
go build -o bin/server .
```

## Docker

Build the image:

```bash
docker build -t go-http-server .
```

Run the container:

```bash
docker run -p 8080:8080 go-http-server
```

## Endpoints

- GET /health — basic health check (returns 200 OK)
- GET /hello — example endpoint (returns greeting)

## Examples

Check health:

```bash
curl -i http://localhost:8080/health
```

Call hello:

```bash
curl http://localhost:8080/hello
```

## Logging & Middleware

Requests are logged by the middleware in `middleware.go`. The middleware prints method, path, status and duration to stdout.

## Development notes

- Router configuration lives in `router.go`.
- Main server setup is in `main.go`.
- Modify or extend middleware in `middleware.go`.

## License

MIT
