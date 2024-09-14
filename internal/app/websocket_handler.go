package app

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Upgrade the connection to a WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections
	},
}

// HandleWebSocket manages WebSocket connections
func (app *Application) HandleWebSocket() echo.HandlerFunc {
	return func(c echo.Context) error {
		conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			c.Logger().Errorf("HandleWebSocket: error upgrading connection to websocket: %v", err)
			return err
		}
		defer conn.Close()
		defer app.ClientService.RemoveClient(conn)

		app.ClientService.AddClient(conn)

		err = app.writeMessageHistoryToNewClient(conn)
		if err != nil {
			c.Logger().Errorf("HandleWebSocket: error writing message history: %v", err)
			return err
		}

		for {
			newMessageContent, err := app.readNewMessage(conn)
			if err != nil {
				c.Logger().Errorf("HandleWebSocket: error reading new message: %v", err)
				return err
			}

			message, err := app.MessageService.CreateMessage(newMessageContent)
			if err != nil {
				c.Logger().Errorf("HandleWebSocket: error creating message: %v", err)
				return err
			}

			app.ClientService.BroadcastMessage(message)
		}
	}
}

func (app *Application) writeMessageHistoryToNewClient(conn *websocket.Conn) error {
	previousMessages := app.MessageService.GetAllMessages()
	for _, msg := range previousMessages {
		err := conn.WriteJSON(msg)
		if err != nil {
			return err
		}
	}

	return nil
}

func (app *Application) readNewMessage(conn *websocket.Conn) (string, error) {
	_, rawMessage, err := conn.ReadMessage()
	if err != nil {
		return "", err
	}

	var parsedContent struct {
		Content string `json:"content"`
	}

	err = json.Unmarshal(rawMessage, &parsedContent)
	if err != nil {
		return "", err
	}

	return parsedContent.Content, nil
}
