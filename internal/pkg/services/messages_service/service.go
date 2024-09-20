package messages_service

import (
	"Assignment2/internal/dependencies"
	"Assignment2/internal/pkg/models"
	"errors"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"sync"
	"time"
)

// MessageService implementation of IMessageService interface
type MessageService struct {
	Logger   *zap.SugaredLogger
	messages []*models.Message
	mu       sync.RWMutex
}

var _ dependencies.IMessageService = (*MessageService)(nil)

// NewMessageService creates a new instance of MessageService
func NewMessageService(l *zap.Logger) (dependencies.IMessageService, error) {
	if l == nil {
		return nil, errors.New("logger must be provided")
	}

	return &MessageService{
		Logger:   l.Sugar(),
		messages: make([]*models.Message, 0),
	}, nil
}

// CreateMessage adds a new message to the service
func (m *MessageService) CreateMessage(content string) (*models.Message, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	msg := models.NewMessage(uuid.Must(uuid.NewV7()), content, time.Now())

	m.messages = append(m.messages, msg)

	return msg, nil
}

// GetMessageCount returns the current message count
func (m *MessageService) GetMessageCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.messages)
}

// GetAllMessages returns all messages
func (m *MessageService) GetAllMessages() []*models.Message {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.messages
}
