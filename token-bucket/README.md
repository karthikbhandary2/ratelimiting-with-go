# Token Bucket Rate Limiting

Custom token bucket algorithm implementation using Go's rate limiter.

## Configuration
- **Rate limit:** 2 requests per second, burst of 4
- **Scope:** Global (all clients share the same bucket)
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
- Global rate limiting across all clients
- Token bucket algorithm for burst handling
- Clean separation of rate limiting logic in `limit.go`
- Simple middleware pattern

## Implementation Details
- Uses `rate.NewLimiter(2, 4)` for 2 tokens/second with burst of 4
- Rate limiting logic separated into `rateLimiter` function
- Returns HTTP 429 status when rate limited

## Files
- `main.go` - HTTP server and endpoint handler
- `limit.go` - Rate limiting middleware implementation

## Dependencies
- `golang.org/x/time v0.14.0`
