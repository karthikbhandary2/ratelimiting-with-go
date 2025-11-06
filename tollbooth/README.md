# Tollbooth Rate Limiting

Rate limiting implementation using the `github.com/didip/tollbooth/v7` library.

## Configuration
- **Rate limit:** 1 request per second
- **Port:** 8081
- **Endpoint:** `/ping`

## Usage
```bash
go mod tidy
go run main.go
```

## Testing
```bash
# Test rate limiting
for i in {1..3}; do curl http://localhost:8081/ping; done
```

## Features
- Production-ready third-party library
- Built-in middleware functionality
- Custom JSON error messages
- Simple integration with existing handlers

## Dependencies
- `github.com/didip/tollbooth/v7 v7.0.2`
