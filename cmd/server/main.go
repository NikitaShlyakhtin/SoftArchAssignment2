package main

import (
	"Assignment2/internal/app"
	"Assignment2/internal/pkg/services/client_service"
	"Assignment2/internal/pkg/services/messages_service"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	messageService := messages_service.NewMessageService()
	clientService := client_service.NewClientService()

	application := app.NewApplication(messageService, clientService)

	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	e.GET("/messages/count", application.GetMessageCount())
	e.GET("/messages/ws", application.HandleWebSocket())

	port := 8080
	e.Logger.Infof("Starting server on port %d", port)
	if err := e.Start(fmt.Sprintf(":%d", port)); err != nil {
		return
	}
}
