package main

import (
	"SkillsRock/internal/config"
	"SkillsRock/internal/handler"
	"SkillsRock/internal/repos"
	"SkillsRock/internal/server"
	"SkillsRock/internal/service"
	"context"
	"log/slog"
	"os"
)

// @title SkillsRock
// @version 1.0.0
// @description TODO-list

// @host localhost:8080
// @BasePath /list
func main() {
	ctx := context.Background()

	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	log := slog.New(slog.NewTextHandler(os.Stdout, opts))

	cfg := config.InitConfig()

	log.Info("Successfully initialized config")

	db, err := repos.NewPostgresDB(ctx, repos.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		DBName:   cfg.DB.DBName,
		SSLMode:  cfg.DB.SSLMode,
	})
	if err != nil {
		log.Error("Failed to connect to database", err.Error())
		os.Exit(1)
	}

	log.Info("Successfully connected to database")

	rep := repos.NewRepos(db)
	services := service.NewService(rep)
	handlers := handler.NewHandler(services)
	fib := handlers.InitRoutes()

	log.Info("Successfully initialized handlers")
	log.Info("Starting server")

	err = server.Start(fib, cfg.Server.Port)
	if err != nil {
		log.Error("Server is stopped", err.Error())
		os.Exit(1)
	}

	if err := db.Close(ctx); err != nil {
		log.Error("Failed to close database connection", err.Error())
	}
}
