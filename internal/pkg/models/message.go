package models

import (
	"github.com/google/uuid"
	"time"
)

// Message represents a chat message
type Message struct {
	id uuid.UUID

	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// NewMessage creates a new Message instance
func NewMessage(id uuid.UUID, content string, createdAt time.Time) *Message {
	return &Message{
		id:        id,
		Content:   content,
		CreatedAt: createdAt,
	}
}
