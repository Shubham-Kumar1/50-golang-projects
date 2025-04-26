# Chat App Backend

A WebSocket-based chat application backend built with Go.

## Features

- Real-time WebSocket communication
- Configurable environment settings
- Structured logging
- Graceful shutdown
- Connection pooling

## Prerequisites

- Go 1.21 or higher
- Git

## Development Setup

1. Clone the repository:
```bash
git clone https://github.com/Shubham-Kumar1/chat-app.git
cd chat-app/backend
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file:
```bash
cp .env.example .env
```

4. Run the development server:
```bash
go run main.go
```

## Environment Variables

- `PORT`: Server port (default: 9090)
- `WS_PATH`: WebSocket endpoint path (default: /ws)
- `LOG_LEVEL`: Logging level (debug, info, warn, error)
- `MAX_CONNECTIONS`: Maximum number of WebSocket connections
- `PING_INTERVAL`: WebSocket ping interval in seconds
- `PONG_WAIT`: WebSocket pong wait time in seconds
- `WRITE_WAIT`: WebSocket write wait time in seconds
- `ALLOWED_ORIGINS`: Allowed CORS origins

## Project Structure

```
backend/
├── config/         # Configuration management
├── logger/         # Logging utilities
├── websocket/      # WebSocket implementation
├── main.go         # Application entry point
├── go.mod          # Go module file
└── README.md       # This file
```

## API Endpoints

### WebSocket Connection

```
ws://localhost:9090/ws
```

## Development

To run the server in development mode with hot reload:

```bash
go run main.go
```

## Logging

The application uses structured logging with different levels:
- DEBUG: Detailed debugging information
- INFO: General operational information
- WARN: Warning messages
- ERROR: Error messages

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request 