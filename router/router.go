package router

import (
	"Manifolk/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(todoHandler handler.TodoHandler) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")

	api.GET("/get-all-todos", todoHandler.GetAll)
	api.POST("/create-todo", todoHandler.Create)
	api.DELETE("/delete-todo", todoHandler.Delete)
	api.PUT("/update-todo-status", todoHandler.UpdateStatus)
	api.PUT("/update-todo-name", todoHandler.UpdateName)

	return router
}
