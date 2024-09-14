# Server implementation for anonymous chat application

This is a backend for anonymous chat application. 
The backend is implemented in Go and provides REST API 
for retrieving total message count and real-time message 
creation/reading via WebSockets.

## Getting Started

Set up the server with docker-compose:

```bash
docker-compose up --build
```

Once the application is running, the base URL for the endpoints
will be `127.0.0.1:8080`.

## API Endpoints

### 1. Message Count
- **GET** `/messages/count` - Returns the total number of messages sent in the chat.

### 2. WebSocket Connection
- **GET** `/messages/ws` - Establishes a WebSocket connection for real-time messaging.

#### Message Format
The server accepts messages in the following format:

```json
{
  "content": "string"
}
```

Upon receiving a message, the server responds with:

```json
{
  "content": "string",
  "created_at": "timestamp"
}
```

When a client connects for the first time, all previously created messages will be sent as individual messages.