package app

import (
	"Assignment2/internal/dependencies"
	"go.uber.org/zap"
)

// Application holds the application state and dependencies
type Application struct {
	Logger         *zap.SugaredLogger
	MessageService dependencies.IMessageService
	ClientService  dependencies.IClientService
}

// NewApplication initializes a new Application instance
func NewApplication(
	logger *zap.Logger,
	messageService dependencies.IMessageService,
	clientService dependencies.IClientService,
) *Application {
	return &Application{
		Logger:         logger.Sugar(),
		MessageService: messageService,
		ClientService:  clientService,
	}
}
