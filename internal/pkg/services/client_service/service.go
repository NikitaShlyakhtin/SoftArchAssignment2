package client_service

import (
	"Assignment2/internal/dependencies"
	"Assignment2/internal/pkg/models"
	"errors"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"sync"
)

// ClientService is the implementation of IClientService
type ClientService struct {
	Logger  *zap.SugaredLogger
	clients map[*websocket.Conn]struct{}
	mu      sync.RWMutex
}

var _ dependencies.IClientService = (*ClientService)(nil)

// NewClientService creates a new instance of IClientService
func NewClientService(l *zap.Logger) (dependencies.IClientService, error) {
	if l == nil {
		return nil, errors.New("logger must be provided")
	}

	return &ClientService{
		Logger:  l.Sugar(),
		clients: make(map[*websocket.Conn]struct{}),
	}, nil
}

// AddClient adds a new WebSocket client
func (c *ClientService) AddClient(conn *websocket.Conn) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.clients[conn] = struct{}{}
}

// RemoveClient removes a WebSocket client
func (c *ClientService) RemoveClient(conn *websocket.Conn) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.clients, conn)
}

// BroadcastMessage sends a message to all connected clients
func (c *ClientService) BroadcastMessage(message *models.Message) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for client := range c.clients {
		err := client.WriteJSON(message)
		if err != nil {
			err = client.Close()
			if err != nil {
				return
			}

			delete(c.clients, client)
		}
	}
}

// GetClients returns the current clients
func (c *ClientService) GetClients() map[*websocket.Conn]struct{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.clients
}
