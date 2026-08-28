package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"Manifolk/model"
	"Manifolk/repository"
)

type TodoHandler struct {
	repository repository.TodoReposPosgre
}

func NewTodoHandler(repository repository.TodoReposPosgre) TodoHandler {
	return TodoHandler{
		repository: repository,
	}
}

func (h TodoHandler) GetAll(c *gin.Context) {
	todos, err := h.repository.GetAllData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get all todo list",
		"todos":   todos,
	})
}

func (h TodoHandler) Create(c *gin.Context) {
	var todo model.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.repository.CreateData(todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully created todo",
		"todo":    todo,
	})
}

func (h TodoHandler) Delete(c *gin.Context) {
	var todo model.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.repository.DeletePostgre(todo.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted todo",
	})
}

func (h TodoHandler) UpdateStatus(c *gin.Context) {
	var todo model.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.repository.UpdateDataStatus(
		todo.ID,
		todo.IsDone,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully updated todo status",
	})
}

func (h TodoHandler) UpdateName(c *gin.Context) {
	var todo model.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.repository.UpdateDataName(
		todo.ID,
		todo.Name,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully updated todo name",
	})
}
