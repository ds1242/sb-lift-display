# Snowbird Lift Status
Building a small Snowbird ski lift status sign using a combination between an
ESP32 and a small HTTP server hosted on Digital Ocean. 

## About/Description
This goal of this project is to periodically query Snowbird's lift status API 
and store it in memory cache.  It removes a lot of the extra data they provide
in the API.  This is then hosted with an endpoint that my ESP32 can query
and display the current lift status.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Tests](#tests)
- [Future Enhancements](#future-enhancements)

## Installation
```bash
go mod download
```

## Usage
Set up environment variables:
```bash
PORT=8080
API_KEY=your-secret-key
```

Run the server:
```bash
go run .
```

### Endpoints
- `GET /api/lifts` - Fetch cached lift status (requires `X-API-Key` header)
- `GET /api/test-all-open` - Fetch lift status from test file (requires `X-API-Key` header)

## Tests
Run all tests:
```bash
go test -v
```

Current test coverage:
- **json_test.go**: HTTP response helpers (`RespondWithJSON`, `RespondWithError`)
  - Verifies correct status codes, Content-Type headers, and JSON response bodies
  - Table-driven tests for multiple error code scenarios

- **cache_test.go**: Cache logic (`GetLifts()`)
  - Cache hit scenario with fresh data

- **main_test.go**: Handler functions for `/api/lifts`
  - Authorization checks (valid and invalid API keys)
  - Response formatting and error handling

## Future Enhancements
- Cache miss handling with mocked API calls
- Integration tests with real API responses
- `/api/test-all-open` endpoint tests
