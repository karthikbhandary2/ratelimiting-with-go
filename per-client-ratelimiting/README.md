# Per-Client Rate Limiting

Custom rate limiting implementation that applies limits per client IP address.

## Configuration
- **Rate limit:** 2 requests per second, burst of 4
- **Client cleanup:** Every minute
- **Client timeout:** 3 minutes of inactivity
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
for i in {1..6}; do curl http://localhost:8081/ping; done
```

## Features
- Individual limits per client IP
- Automatic cleanup of inactive clients
- Thread-safe with mutex locks
- Memory-efficient client tracking

## Implementation Details
- Uses `golang.org/x/time/rate` for token bucket algorithm
- Maintains map of client limiters keyed by IP
- Background goroutine cleans up stale clients
- Extracts IP from `r.RemoteAddr`

## Dependencies
- `golang.org/x/time v0.14.0`
