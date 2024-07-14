package handlers

import (
	"log/slog"
	"timeTracker/internal/api/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *services.Service
	logger   *slog.Logger
}

func NewHandler(services *services.Service, logger *slog.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	docs := router.Group("/docs")
	{
		docs.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("/info", h.info)
			users.POST("/", h.createUser)
			users.PUT("/:id", h.updateUser)
			users.DELETE("/:id", h.deleteUser)
		}
		tasks := api.Group("/tasks")
		{
			tasks.GET("/:user_id", h.getTasks)
			tasks.POST("/start_timing/:user_id", h.startTiming)
			tasks.PUT("/end_timing/:user_id", h.endTiming)
		}

	}
	return router
}
