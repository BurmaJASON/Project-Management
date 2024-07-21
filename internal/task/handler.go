package task

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

// CreateTask godoc
// @Summary Create a new task
// @Description Create a new task
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body Task true "Task"
// @Success 201 {object} Task
// @Failure 400 {string} string "Bad Request"
// @Router /tasks [post]
func (h *Handler) CreateTask(c echo.Context) error {
	var task Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := h.service.CreateTask(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, task)
}

// GetTask godoc
// @Summary Get a task by id
// @Description Get a task by id
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Success 200 {object} Task
// @Failure 404 {string} string "Not Found"
// @Router /tasks/{id} [get]
func (h *Handler) GetTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	task, err := h.service.GetTask(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, task)
}

// UpdateTask godoc
// @Summary Update a task by id
// @Description Update a task by id
// @Tags tasks
// @Produce json
// @Param id path int true "Task ID"
// @Param task body Task true "Task"
// @Success 200 {object} Task
// @Failure 400 {string} string "Bad Request"
// @Router /tasks/{id} [put]
func (h *Handler) UpdateTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var task Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	task.ID = id
	if err := h.service.UpdateTask(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, task)
}

// DeleteTask godoc
// @Summary Delete a task
// @Description Delete a task
// @Tags tasks
// @Param id path int true "Task ID"
// @Success 204 {string} string "No Content"
// @Failure 500 {string} string "Internal Server Error"
// @Router /tasks/{id} [delete]
func (h *Handler) DeleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.service.DeleteTask(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
