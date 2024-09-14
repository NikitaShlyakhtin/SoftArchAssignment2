package app

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetMessageCount implements the IHandler interface
func (app *Application) GetMessageCount() echo.HandlerFunc {
	return func(c echo.Context) error {
		count := app.MessageService.GetMessageCount()
		return c.JSON(http.StatusOK, map[string]int{"count": count})
	}
}
