
```markdown
# HTTP Authentication Server

This project implements a simple HTTP server with various features including basic authentication, rate limiting, and handling different HTTP methods and status codes.

## Features

- **Home Page**: Displays a welcome message.
- **GET and POST Handlers**: Handles query parameters and POST data.
- **Rate Limiting**: Limits the number of requests to prevent abuse.
- **Basic Authentication**: Protects certain routes with basic authentication.
- **Status Handlers**: Returns different HTTP status codes.

## Endpoints

- `/`: Home page
- `/200`: Returns a "200 OK" message
- `/404`: Returns a "404 Not Found" message
- `/500`: Returns a "500 Internal Server Error" message
- `/authenticated`: Protected route requiring basic authentication
- `/limited`: Rate-limited endpoint

## Prerequisites

- Go 1.16 or higher
- `golang.org/x/time/rate` package for rate limiting

## Setup

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/http-auth.git
   cd http-auth
   ```

2. Set up environment variables:
   Create a `.env` file in the project root and add:
   ```
   AUTH_USERNAME=your_username
   AUTH_PASSWORD=your_password
   ```

3. Install dependencies:
   ```
   go mod tidy
   ```

## Running the Server

Run the following command in the project root:

```
go run .
```

The server will start on `http://localhost:8080`.

## Testing

### Manual Testing with cURL

1. Test the home page:
   ```
   curl -i 'http://localhost:8080/'
   ```

2. Test status codes:
   ```
   curl -i 'http://localhost:8080/200'
   curl -i 'http://localhost:8080/404'
   curl -i 'http://localhost:8080/500'
   ```

3. Test authentication:
   ```
   curl -i 'http://localhost:8080/authenticated'
   curl -i 'http://localhost:8080/authenticated' -H 'Authorization: Basic YWRtaW46c2VjcmV0'
   ```
   Note: Replace 'YWRtaW46c2VjcmV0' with the Base64 encoding of your username:password.

4. Test rate limiting:
   ```
   ab -n 1000 -c 100 'http://localhost:8080/limited'
   ```

### Automated Testing

Run test with

```
go test
```

## Load Testing

Install ApacheBench (ab) and run:

```
ab -n 10000 -c 100 'http://localhost:8080/'
```


