# Chat Application Backend

A robust real-time chat application backend built with Go, featuring WebSocket communication, structured logging, and clean architecture.

## Features

- Real-time bidirectional communication using WebSockets
- Structured logging with different log levels
- Clean architecture with separation of concerns
- Configuration management
- Client connection pooling
- Message broadcasting
- User presence notifications

## Project Structure

```
backend/
├── config/
│   └── config.go         # Configuration management
├── logger/
│   └── logger.go         # Structured logging implementation
├── websocket/
│   ├── client.go         # WebSocket client handling
│   ├── message.go        # Message type definitions
│   └── pool.go           # Connection pool management
└── main.go               # Application entry point
```

## Prerequisites

- Go 1.21 or higher
- Git

## Setup and Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd chat-app/backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

The server will start on `localhost:8080` by default.

## Configuration

The application can be configured using environment variables:

- `PORT`: Server port (default: 8080)
- `LOG_LEVEL`: Logging level (debug, info, warn, error) (default: info)

## WebSocket API

### Connection

Connect to the WebSocket server:
```
ws://localhost:8080/ws
```

### Message Types

1. System Messages (Type: 1)
   - New user connection notification
   - User disconnection notification

2. Chat Messages (Type: 2)
   - Regular chat messages between users

### Message Format

```json
{
    "type": 1,
    "body": "message content"
}
```

## Architecture

### Components

1. **WebSocket Pool**
   - Manages client connections
   - Handles client registration/unregistration
   - Broadcasts messages to all connected clients

2. **Client Handler**
   - Manages individual WebSocket connections
   - Handles message reading and writing
   - Implements connection lifecycle

3. **Logger**
   - Provides structured logging
   - Supports multiple log levels
   - Thread-safe logging operations

4. **Configuration**
   - Centralized configuration management
   - Environment variable support
   - Default configuration values

## Error Handling

The application implements comprehensive error handling:
- Connection errors
- Message parsing errors
- Configuration errors
- Logging errors

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request
