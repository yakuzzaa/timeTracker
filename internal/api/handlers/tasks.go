package handlers

import (
	"net/http"
	"strings"

	"github.com/yakuzzaa/timeTracker/internal/api/serializers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Create task
// @Description Creates a new task for a user and starts the timer
// @Accept json
// @Produce json
// @Tags tasks
// @Param user_id path string true "User ID"
// @Success 201 {object} serializers.CreateTaskResponse
// @Failure 400 {object} serializers.ErrorResponse
// @Failure 500 {object} serializers.ErrorResponse
// @Router /tasks/start_timing/{user_id} [post]
func (h *Handler) startTiming(c *gin.Context) {
	h.logger.Debug("Entered startTiming handler")

	userId, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		h.logger.Debug("Invalid UUID format", "error", err)
		c.JSON(http.StatusBadRequest, serializers.ErrorResponse{
			Message: "Invalid UUID format",
			Error:   err.Error(),
		})
		return
	}

	h.logger.Debug("Parsed userId successfully", "userId", userId)

	taskId, err := h.services.Task.Create(userId)
	if err != nil {
		h.logger.Error("Failed to create task", "error", err)
		c.JSON(http.StatusInternalServerError, serializers.ErrorResponse{
			Message: "Failed to create task",
			Error:   err.Error(),
		})
		return
	}

	h.logger.Info("Task created successfully", "taskId", taskId)

	c.JSON(http.StatusCreated, serializers.CreateTaskResponse{
		Id:     taskId,
		Status: "Start timing",
	})
}

// @Summary Update task
// @Description Updates a task (ends the timer)
// @Accept json
// @Produce json
// @Tags tasks
// @Param user_id path string true "User ID"
// @Param input body serializers.UpdateTaskRequest true "Request"
// @Success 200 {object} serializers.UpdateTaskResponse
// @Failure 400 {object} serializers.ErrorResponse
// @Failure 404 {object} serializers.ErrorResponse
// @Failure 500 {object} serializers.ErrorResponse
// @Router /tasks/end_timing/{user_id} [put]
func (h *Handler) endTiming(c *gin.Context) {
	h.logger.Debug("Entered endTiming handler")

	var req serializers.UpdateTaskRequest

	userId, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		h.logger.Debug("Invalid UUID format for user_id", "error", err)
		c.JSON(http.StatusBadRequest, serializers.ErrorResponse{
			Message: "Invalid UUID format for user_id",
			Error:   err.Error(),
		})
		return
	}

	h.logger.Debug("Parsed userId successfully", "userId", userId)

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Debug("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, serializers.ErrorResponse{
			Message: "Failed to bind JSON",
			Error:   err.Error(),
		})
		return
	}

	h.logger.Debug("Request JSON bound successfully", "request", req)

	if err := h.services.Task.Update(userId, req.Id); err != nil {
		h.logger.Error("Failed to update task", "error", err)
		if strings.Contains(err.Error(), "task with ID "+req.Id.String()+" not found for user with ID "+userId.String()+"") {
			c.JSON(http.StatusNotFound, serializers.ErrorResponse{
				Message: "Task or User not found",
				Error:   err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, serializers.ErrorResponse{
			Message: "Failed to update task",
			Error:   err.Error(),
		})
		return
	}

	h.logger.Info("Task updated successfully", "taskId", req.Id)

	c.JSON(http.StatusOK, serializers.UpdateTaskResponse{Status: "End timing"})
}

// @Summary Get tasks
// @Description Retrieves all tasks of a user, sorted by descending total time spent
// @Accept json
// @Produce json
// @Tags tasks
// @Param user_id path string true "User ID"
// @Success 200 {object} serializers.GetTaskResponse
// @Failure 400 {object} serializers.ErrorResponse
// @Failure 500 {object} serializers.ErrorResponse
// @Router /tasks/{user_id} [get]
func (h *Handler) getTasks(c *gin.Context) {
	h.logger.Debug("Entered getTasks handler")

	userId, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		h.logger.Debug("Invalid UUID format for user_id", "error", err)
		c.JSON(http.StatusBadRequest, serializers.ErrorResponse{
			Message: "Invalid UUID format for user_id",
			Error:   err.Error(),
		})
		return
	}

	h.logger.Debug("Parsed userId successfully", "userId", userId)

	taskResponse, err := h.services.Task.Get(userId)
	if err != nil {
		h.logger.Error("Failed to get tasks", "error", err)
		c.JSON(http.StatusInternalServerError, serializers.ErrorResponse{
			Message: "Failed to get tasks",
			Error:   err.Error(),
		})
		return
	}

	h.logger.Info("Tasks retrieved successfully", "userId", userId)

	c.JSON(http.StatusOK, taskResponse)
}
