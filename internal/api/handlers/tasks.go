package handlers

import (
	"net/http"
	"timeTracker/internal/api/serializers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Создание задачи
// @Description Создает новую задачу для пользователя и начинает отсчет
// @Accept json
// @Produce json
// @Param user_id path string true "id пользователя"
// @Success 201 {object} serializers.CreateTaskResponse
// @Router /tasks/start_timing/{user_id} [post]
func (h *Handler) startTiming(c *gin.Context) {
	userId, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID format"})
		return
	}

	taskId, err := h.services.Task.Create(userId)

	c.JSON(http.StatusCreated, serializers.CreateTaskResponse{
		Id:     taskId,
		Status: "Отсчет начат",
	})
}

// @Summary Обновление задачи
// @Description Обновляет задачу (заканчивает отсчет)
// @Accept json
// @Produce json
// @Param user_id path string true "id пользователя"
// @Param input body serializers.UpdateTaskRequest true "Запрос"
// @Success 200 {object} serializers.UpdateTaskResponse
// @Router /tasks/end_timing/{user_id} [put]
func (h *Handler) endTiming(c *gin.Context) {
	var req serializers.UpdateTaskRequest
	userId, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID format"})
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Task.Update(userId, req.Id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, serializers.UpdateTaskResponse{Status: "Отсчет закончен"})
}

// @Summary Получение задач
// @Description Получить все задачи пользователся, отсортированные по убыванию затраченного времени
// @Accept json
// @Produce json
// @Param user_id path string true "id пользователя"
// @Success 200 {object} serializers.GetTaskResponse
// @Router /tasks/{user_id} [get]
func (h *Handler) getTasks(c *gin.Context) {
	userId, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID format"})
	}

	taskResponse, err := h.services.Task.Get(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, taskResponse)
}
