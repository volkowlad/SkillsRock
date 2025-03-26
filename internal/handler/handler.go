package handler

import (
	"SkillsRock/internal/service"
	"github.com/Flussen/swagger-fiber-v3"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"

	_ "SkillsRock/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) InitRoutes() *fiber.App {
	r := fiber.New()

	r.Use(limiter.New(limiter.Config{
		Max: 1000,
	}))

	list := r.Group("/list")
	{
		list.Post("/tasks", h.createTask)
		list.Get("/tasks", h.getList)
		list.Put("/tasks/:id", h.updateTask)
		list.Delete("/tasks/:id", h.deleteTask)
	}

	r.Get("/swagger/*", swagger.HandlerDefault)

	return r
}
