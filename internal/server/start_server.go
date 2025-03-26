package server

import (
	"github.com/gofiber/fiber/v3"
	"log/slog"
	"os"
	"os/signal"
)

func Start(a *fiber.App, port string) error {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := a.Shutdown(); err != nil {
			slog.Error("Oops... Server is not shutting down! Reason: %v", err.Error())
		}

		close(idleConnsClosed)
	}()

	addr := ":" + port

	go func() {
		if err := a.Listen(addr); err != nil {
			slog.Error("Oops... Server is not running! Reason: %v", err.Error())
		}
	}()

	<-idleConnsClosed
	return nil
}
