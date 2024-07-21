package project

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

// CreateProject godoc
// @Summary Create a new project
// @Description Create a new project
// @Tags projects
// @Accept  json
// @Produce  json
// @Param project body Project true "Project"
// @Success 201 {object} Project
// @Failure 400 {string} string "Bad Request"
// @Router /projects [post]
func (h *Handler) CreateProject(c echo.Context) error {
	var project Project
	if err := c.Bind(&project); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := h.service.CreateProject(&project); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, project)
}

// GetProject godoc
// @Summary Get a project by id
// @Description Get a project by id
// @Tags projects
// @Accept  json
// @Produce  json
// @Param id path int true "Project ID"
// @Success 200 {object} Project
// @Failure 404 {string} string "Not Found"
// @Router /projects/{id} [get]
func (h *Handler) GetProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	project, err := h.service.GetProject(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, project)
}

// UpdateProject godoc
// @Summary Update a project by id
// @Description Update a project by id
// @Tags projects
// @Produce json
// @Param id path int true "Project ID"
// @Param project body Project true "Project"
// @Success 200 {object} Project
// @Failure 400 {string} string "Bad Request"
// @Router /projects/{id} [put]
func (h *Handler) UpdateProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var project Project
	if err := c.Bind(&project); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	project.ID = id
	if err := h.service.UpdateProject(&project); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, project)
}

// DeleteProject godoc
// @Summary Delete a project
// @Description Delete a project
// @Tags projects
// @Param id path int true "Project ID"
// @Success 204 {string} string "No Content"
// @Failure 500 {string} string "Internal Server Error"
// @Router /projects/{id} [delete]
func (h *Handler) DeleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.service.DeleteProject(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
