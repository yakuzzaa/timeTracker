package handlers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
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
			tasks.GET("/:id", h.getTasks)
			tasks.POST("/start_timing", h.startTiming)
			tasks.PUT("/end_timing/:id", h.endTiming)
		}

	}
	return router
}
