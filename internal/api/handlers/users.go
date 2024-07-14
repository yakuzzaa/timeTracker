package handlers

import (
	"net/http"
	"strings"
	"timeTracker/internal/api/serializers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Create a user
// @Description Create a new user based on the data provided in the request body
// @Accept json
// @Produce json
// @Tags users
// @Param input body serializers.CreateUserRequest true "Request"
// @Success 201 {object} serializers.CreateUserResponse
// @Failure 400 {object} serializers.ErrorResponse
// @Failure 500 {object} serializers.ErrorResponse
// @Router /users [post]
func (h *Handler) createUser(c *gin.Context) {
	var req serializers.CreateUserRequest

	h.logger.Debug("Entered createUser handler")

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Debug("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, serializers.ErrorResponse{
			Message: "Failed to bind JSON",
			Error:   err.Error(),
		})
		return
	}

	h.logger.Debug("Request JSON bound successfully", "request", req)

	userId, err := h.services.User.Create(req.Passport)
	if err != nil {
		h.logger.Error("Failed to create user", "error", err)
		c.JSON(http.StatusInternalServerError, serializers.ErrorResponse{
			Message: "Failed to create user",
			Error:   err.Error(),
		})
		return
	}

	h.logger.Info("User created successfully", "userId", userId)

	c.JSON(http.StatusCreated, serializers.CreateUserResponse{Id: userId})
}

// @Summary Retrieve a list of users
// @Description Returns a list of users with optional filtering and pagination capabilities
// @Accept json
// @Produce json
// @Tags users
// @Param id query string false "User ID"
// @Param passportSeries query string false "Passport series"
// @Param passportNumber query string false "Passport number"
// @Param name query string false "Name"
// @Param surname query string false "Surname"
// @Param patronymic query string false "Patronymic"
// @Param address query string false "Address"
// @Param page query int false "Page number"
// @Param pageSize query int false "Number of records per page"
// @Success 200 {object} serializers.GetUsersResponse
// @Failure 400 {object} serializers.ErrorResponse
// @Failure 500 {object} serializers.ErrorResponse
// @Router /users/info [get]
func (h *Handler) info(c *gin.Context) {
	h.logger.Debug("Entered info handler")

	var req serializers.GetUsersRequest
	idStr := c.Query("id")

	if idStr != "" {
		id, err := uuid.Parse(idStr)
		if err != nil {
			h.logger.Debug("Invalid UUID", "error", err)
			c.JSON(http.StatusBadRequest, serializers.ErrorResponse{
				Message: "Invalid UUID",
				Error:   err.Error(),
			})
			return
		}
		req.Id = &id
	}

	query := c.Request.URL.Query()
	query.Del("id")
	c.Request.URL.RawQuery = query.Encode()

	if err := c.ShouldBindQuery(&req); err != nil {
		h.logger.Debug("Failed to bind query parameters", "error", err)
		c.JSON(http.StatusBadRequest, serializers.ErrorResponse{
			Message: "Failed to bind query parameters",
			Error:   err.Error(),
		})
		return
	}

	h.logger.Debug("Query parameters bound successfully", "request", req)

	usersResponse, err := h.services.User.Get(req)
	if err != nil {
		h.logger.Error("Failed to get users", "error", err)
		c.JSON(http.StatusInternalServerError, serializers.ErrorResponse{
			Message: "Failed to get users",
			Error:   err.Error(),
		})
		return
	}

	h.logger.Info("Users retrieved successfully", "response", usersResponse)

	c.JSON(http.StatusOK, usersResponse)
}

// @Summary Update user data
// @Description Updates information of the user with the specified identifier
// @Accept json
// @Produce json
// @Tags users
// @Param id path string true "User ID"
// @Param input body serializers.UpdateUserRequest true "Request"
// @Success 200 {object} serializers.UpdateUserResponse
// @Failure 400 {object} serializers.ErrorResponse
// @Failure 404 {object} serializers.ErrorResponse
// @Failure 500 {object} serializers.ErrorResponse
// @Router /users/{id} [put]
func (h *Handler) updateUser(c *gin.Context) {
	h.logger.Debug("Entered updateUser handler")

	userId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		h.logger.Debug("Invalid UUID format", "error", err)
		c.JSON(http.StatusBadRequest, serializers.ErrorResponse{
			Message: "Invalid UUID",
			Error:   err.Error(),
		})
		return
	}

	var req serializers.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Debug("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.logger.Debug("Request JSON bound successfully", "request", req)

	if err := h.services.User.Update(userId, req); err != nil {
		if strings.Contains(err.Error(), "user with ID "+userId.String()+" not found") {
			c.JSON(http.StatusNotFound, serializers.ErrorResponse{
				Message: "User not found",
				Error:   err.Error(),
			})
			return
		}
		h.logger.Error("Failed to update user", "error", err)
		c.JSON(http.StatusInternalServerError, serializers.ErrorResponse{
			Message: "Failed to update user",
			Error:   err.Error(),
		})
		return
	}

	h.logger.Info("User updated successfully", "userId", userId)

	c.JSON(http.StatusOK, serializers.UpdateUserResponse{Status: "updated"})

}

// @Summary Delete user
// @Description Deletes the user with the specified identifier
// @Accept json
// @Produce json
// @Tags users
// @Param id path string true "User ID"
// @Success 200 {object} serializers.DeleteUserResponse
// @Failure 400 {object} serializers.ErrorResponse
// @Failure 404 {object} serializers.ErrorResponse
// @Failure 500 {object} serializers.ErrorResponse
// @Router /users/{id} [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	h.logger.Debug("Entered deleteUser handler")

	userId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		h.logger.Debug("Invalid UUID format", "error", err)
		c.JSON(http.StatusBadRequest, serializers.ErrorResponse{
			Message: "Invalid UUID",
			Error:   err.Error(),
		})
		return
	}

	h.logger.Debug("Parsed userId successfully", "userId", userId)

	if err := h.services.User.Delete(c, userId); err != nil {
		h.logger.Error("Failed to delete user", "error", err)
		if strings.Contains(err.Error(), "user with ID "+userId.String()+" not found") {
			c.JSON(http.StatusNotFound, serializers.ErrorResponse{
				Message: "User not found",
				Error:   err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, serializers.ErrorResponse{
			Message: "Failed to delete user",
			Error:   err.Error(),
		})
		return
	}

	h.logger.Info("User deleted successfully", "userId", userId)

	c.JSON(http.StatusOK, serializers.DeleteUserResponse{Status: "deleted"})
}
