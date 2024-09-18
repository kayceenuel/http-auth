
```markdown
# HTTP Authentication Server

This project implements a simple HTTP server with various features including basic authentication, rate limiting, and handling different HTTP methods and status codes.

## Features

- Responds to GET and POST requests
- Handles different HTTP status codes (200, 404, 500)
- Implements basic HTTP authentication
- Includes rate limiting
- Escapes HTML in user inputs to prevent XSS attacks

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

Run the test suite with:

```
go test
```

## Load Testing

Install ApacheBench (ab) and run:

```
ab -n 10000 -c 100 'http://localhost:8080/'
```

Analyze the results to understand the server's performance under load.



