package app

import (
	"Assignment2/internal/dependencies"
)

// Application holds the application state and dependencies
type Application struct {
	MessageService dependencies.IMessageService
	ClientService  dependencies.IClientService
}

// NewApplication initializes a new Application instance
func NewApplication(
	messageService dependencies.IMessageService,
	clientService dependencies.IClientService,
) *Application {
	return &Application{
		MessageService: messageService,
		ClientService:  clientService,
	}
}
