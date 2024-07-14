package handlers

import (
	"net/http"
	"timeTracker/internal/api/serializers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Создание пользователя
// @Description Создает нового пользователя на основе переданных данных в теле запроса
// @Accept json
// @Produce json
// @Tags users
// @Param input body serializers.CreateUserRequest true "Запрос"
// @Success 201 {object} serializers.CreateUserResponse
// @Router /users [post]
func (h *Handler) createUser(c *gin.Context) {
	var req serializers.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := req.Passport.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userId, err := h.services.User.Create(req.Passport)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, serializers.CreateUserResponse{Id: userId})
}

// @Summary Получение списка пользователей
// @Description Возвращает список пользователей с возможностью фильтрации и пагинации
// @Accept json
// @Produce json
// @Tags users
// @Param id query string false "Id пользователя"
// @Param passportSeries query string false "Серия паспорта"
// @Param passportNumber query string false "Номер паспорта"
// @Param name query string false "Имя"
// @Param surname query string false "Фамилия"
// @Param patronymic query string false "Отчество"
// @Param address query string false "Адрес"
// @Param page query int false "Страница"
// @Param pageSize query int false "Кол-во записей на странице"
// @Success 200 {object} serializers.GetUsersResponse
// @Router /users/info [get]
func (h *Handler) info(c *gin.Context) {
	var req serializers.GetUsersRequest

	idStr := c.Query("id")
	if idStr != "" {
		id, err := uuid.Parse(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
			return
		}
		req.Id = &id
	}
	query := c.Request.URL.Query()
	query.Del("id")
	c.Request.URL.RawQuery = query.Encode()

	if err := c.ShouldBindQuery(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usersResponse, err := h.services.User.Get(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usersResponse)
}

// @Summary Обновление данных пользователя
// @Description Обновляет информацию о пользователе с указанным идентификатором
// @Accept json
// @Produce json
// @Tags users
// @Param id path string true "Id пользователя"
// @Param input body serializers.UpdateUserRequest true "Запрос"
// @Success 200 {object} serializers.UpdateUserResponse
// @Router /users/{id} [put]
func (h *Handler) updateUser(c *gin.Context) {
	userId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID format"})
		return
	}

	var req serializers.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.User.Update(userId, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, serializers.UpdateUserResponse{Status: "updated"})

}

// @Summary Удаление пользователя
// @Description Удаляет пользователя с указанным идентификатором
// @Accept json
// @Produce json
// @Tags users
// @Param id path string true "Id пользователя"
// @Success 200 {object} serializers.DeleteUserResponse
// @Router /users/{id} [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	userId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID format"})
		return
	}

	if err := h.services.User.Delete(c, userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, serializers.DeleteUserResponse{Status: "deleted"})

}
