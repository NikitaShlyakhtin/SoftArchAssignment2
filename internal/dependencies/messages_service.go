package dependencies

import "Assignment2/internal/pkg/models"

// IMessageService defines the methods for working with messages
type IMessageService interface {
	CreateMessage(message string) (*models.Message, error)
	GetMessageCount() int
	GetAllMessages() []*models.Message
}
