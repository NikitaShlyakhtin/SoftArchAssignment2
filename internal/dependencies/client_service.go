package dependencies

import (
	"Assignment2/internal/pkg/models"
	"github.com/gorilla/websocket"
)

// IClientService defines the methods for managing WebSocket clients
type IClientService interface {
	AddClient(conn *websocket.Conn)
	RemoveClient(conn *websocket.Conn)
	BroadcastMessage(message *models.Message)
	GetClients() map[*websocket.Conn]struct{}
}
