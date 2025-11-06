# Rate Limiting with Go

A comprehensive collection of rate limiting implementations in Go, demonstrating different approaches and strategies for controlling request rates in web applications.

## Overview

This repository contains three different rate limiting implementations:

1. **Tollbooth** - Third-party library implementation
2. **Per-Client Rate Limiting** - Custom implementation with client-specific limits
3. **Token Bucket** - Custom token bucket algorithm implementation

Each implementation serves different use cases and demonstrates various rate limiting strategies commonly used in production systems.

## Project Structure

```
ratelimiting-with-go/
├── tollbooth/              # Third-party library implementation
│   ├── main.go
│   ├── go.mod
│   └── go.sum
├── per-client-ratelimiting/ # Per-client custom implementation
│   ├── main.go
│   ├── go.mod
│   └── go.sum
├── token-bucket/            # Token bucket algorithm implementation
│   ├── main.go
│   ├── limit.go
│   ├── go.mod
│   └── go.sum
└── README.md
```

## Implementations

### 1. Tollbooth Implementation

**Location:** `./tollbooth/`

Uses the popular `github.com/didip/tollbooth/v7` library for rate limiting.

**Features:**
- Simple integration with existing HTTP handlers
- Configurable rate limits
- Custom error messages in JSON format
- Built-in middleware functionality

**Configuration:**
- Rate limit: 1 request per second
- Port: 8081
- Endpoint: `/ping`

**Usage:**
```bash
cd tollbooth
go run main.go
```

**Testing:**
```bash
curl http://localhost:8081/ping
```

### 2. Per-Client Rate Limiting

**Location:** `./per-client-ratelimiting/`

Custom implementation that applies rate limits per client IP address using Go's `golang.org/x/time/rate` package.

**Features:**
- Individual rate limits per client IP
- Automatic cleanup of inactive clients (3-minute timeout)
- Thread-safe implementation with mutex locks
- Memory-efficient client tracking

**Configuration:**
- Rate limit: 2 requests per second with burst of 4
- Client cleanup: Every minute
- Client timeout: 3 minutes of inactivity
- Port: 8081
- Endpoint: `/ping`

**Usage:**
```bash
cd per-client-ratelimiting
go run main.go
```

**Testing:**
```bash
# Test from different IPs or simulate multiple clients
curl http://localhost:8081/ping
```

### 3. Token Bucket Implementation

**Location:** `./token-bucket/`

Custom token bucket algorithm implementation using Go's rate limiter.

**Features:**
- Global rate limiting across all clients
- Token bucket algorithm for burst handling
- Clean separation of rate limiting logic
- Configurable rate and burst parameters

**Configuration:**
- Rate limit: 2 requests per second with burst of 4
- Port: 8081
- Endpoint: `/ping`

**Usage:**
```bash
cd token-bucket
go run main.go
```

**Testing:**
```bash
curl http://localhost:8081/ping
```

## API Response Format

All implementations return consistent JSON responses:

**Success Response:**
```json
{
  "status": "success",
  "body": "Request processed successfully"
}
```

**Rate Limited Response:**
```json
{
  "status": "Request Failed",
  "body": "Rate limit exceeded. Please try again later."
}
```

## Rate Limiting Strategies Comparison

| Implementation | Scope | Algorithm | Use Case |
|----------------|-------|-----------|----------|
| Tollbooth | Global | Token Bucket | Simple, production-ready solution |
| Per-Client | Per IP | Token Bucket | Multi-tenant applications |
| Token Bucket | Global | Token Bucket | Custom control and learning |

## Prerequisites

- Go 1.24.0 or higher
- Internet connection for downloading dependencies

## Installation & Setup

1. Clone or navigate to the project directory:
```bash
cd /home/ubuntu/ratelimiting-with-go
```

2. Choose an implementation and navigate to its directory:
```bash
cd tollbooth  # or per-client-ratelimiting or token-bucket
```

3. Install dependencies:
```bash
go mod tidy
```

4. Run the application:
```bash
go run main.go
```

## Testing Rate Limits

### Basic Testing
```bash
# Single request
curl http://localhost:8081/ping

# Multiple rapid requests to trigger rate limiting
for i in {1..6}; do curl http://localhost:8081/ping; done
```

### Advanced Testing with Apache Bench
```bash
# Test with concurrent requests
ab -n 100 -c 10 http://localhost:8081/ping
```

### Testing Per-Client Rate Limiting
For the per-client implementation, you can simulate different clients using proxy or different network interfaces, or test the cleanup mechanism by waiting for the 3-minute timeout.

## Key Concepts Demonstrated

### 1. Token Bucket Algorithm
- Tokens are added to a bucket at a fixed rate
- Each request consumes a token
- Burst capacity allows temporary spikes in traffic
- Requests are rejected when no tokens are available

### 2. Per-Client Rate Limiting
- Maintains separate rate limiters for each client IP
- Automatic cleanup prevents memory leaks
- Thread-safe operations using mutexes
- Suitable for multi-tenant applications

### 3. Middleware Pattern
- Rate limiting implemented as HTTP middleware
- Clean separation of concerns
- Reusable across different endpoints
- Easy integration with existing applications

## Production Considerations

### Security
- Consider using `X-Forwarded-For` header for client IP detection behind proxies
- Implement proper logging for rate limit violations
- Add monitoring and alerting for rate limit patterns

### Performance
- Monitor memory usage for per-client implementations
- Consider using Redis for distributed rate limiting
- Implement graceful degradation strategies

### Configuration
- Make rate limits configurable via environment variables
- Implement different limits for different endpoints
- Consider implementing rate limit headers (X-RateLimit-*)

## Dependencies

### Tollbooth Implementation
- `github.com/didip/tollbooth/v7 v7.0.2`
- `github.com/go-pkgz/expirable-cache/v3 v3.0.0` (indirect)

### Per-Client & Token Bucket Implementations
- `golang.org/x/time v0.14.0`
