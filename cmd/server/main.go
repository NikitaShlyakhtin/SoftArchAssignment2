package main

import (
	"Assignment2/internal/app"
	"Assignment2/internal/pkg/services/client_service"
	"Assignment2/internal/pkg/services/messages_service"
	"Assignment2/internal/server"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(getFxOptions()).Run()
}

func getFxOptions() fx.Option {
	return fx.Options(
		fx.WithLogger(getEventLogger),
		fx.Provide(
			zap.NewDevelopment,
			messages_service.NewMessageService,
			client_service.NewClientService,
			app.NewApplication,
			server.NewServer,
		),
		fx.Invoke(startServer),
	)
}
