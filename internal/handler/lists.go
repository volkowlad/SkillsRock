package handler

import (
	sr "SkillsRock"
	"github.com/gofiber/fiber/v3"
	"log/slog"
	"strconv"
)

// @Summary Create task
// @Tags TODO-list
// @Description create task
// @ID create-task
// @Accept  json
// @Produce  json
// @Param input body sr.List true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /list/tasks [post]
func (h *Handler) createTask(c fiber.Ctx) error {
	var input sr.List

	if err := c.Bind().Body(&input); err != nil {
		slog.Error("Failed to parse body", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(respError{
			Massage: "Failed to parse body",
		})
	}

	id, err := h.services.CreateTask(c.Context(), input)
	if err != nil {
		slog.Error("Failed to create list", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(respError{
			Massage: "Failed to create list",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": id,
	})
}

type getList struct {
	Data []sr.List `json:"data"`
}

// @Summary Get List
// @Tags TODO-list
// @Description get list
// @ID get-list
// @Accept  json
// @Produce  json
// @Success 200 {integer} getList
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /list/tasks [get]
func (h *Handler) getList(c fiber.Ctx) error {
	tasks, err := h.services.GetList(c.Context())
	if err != nil {
		slog.Error("Failed to get tasks", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(respError{
			Massage: "Failed to get tasks",
		})
	}

	return c.Status(fiber.StatusOK).JSON(getList{
		Data: tasks,
	})
}

// @Summary Update task
// @Tags TODO-list
// @Description update task
// @ID update-task
// @Accept  json
// @Produce  json
// @Success 200 {integer} sr.List
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /list/tasks/:id [put]
func (h *Handler) updateTask(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		slog.Error("Failed to parse id", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(respError{
			Massage: "Failed to parse params",
		})
	}

	var input sr.List
	if err := c.Bind().Body(&input); err != nil {
		slog.Error("Failed to parse body", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(respError{
			Massage: "Failed to parse body",
		})
	}

	err = h.services.UpdateTask(c.Context(), input, id)
	if err != nil {
		slog.Error("Failed to create list", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(respError{
			Massage: "Failed to update list",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": id,
	})
}

// @Summary Delete task
// @Tags TODO-list
// @Description delete task
// @ID delete-task
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /list/tasks/:id [delete]
func (h *Handler) deleteTask(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		slog.Error("Failed to parse id", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(respError{
			Massage: "Failed to parse params",
		})
	}

	err = h.services.DeleteTask(c.Context(), id)
	if err != nil {
		slog.Error("Failed to delete task", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(respError{
			Massage: "Failed to delete task",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}
