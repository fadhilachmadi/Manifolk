package router

import (
	"Manifolk/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(todoHandler handler.TodoHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/get-all-todos", todoHandler.GetAll)
	router.POST("/create-todo", todoHandler.Create)
	router.DELETE("/delete-todo", todoHandler.Delete)
	router.PUT("/update-todo-status", todoHandler.UpdateStatus)
	router.PUT("/update-todo-name", todoHandler.UpdateName)

	return router
}
