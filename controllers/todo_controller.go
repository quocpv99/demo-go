package controllers

import (
	"go_api/models"
	"go_api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	TodoService *services.TodoService
}

func NewTodoController(todoService *services.TodoService) *TodoController {
	return &TodoController{TodoService: todoService}
}

func (c *TodoController) CreateTodo(ctx *gin.Context) {
	var todo models.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.TodoService.CreateTodo(&todo)
	ctx.JSON(http.StatusCreated, todo)
}

func (c *TodoController) GetTodos(ctx *gin.Context) {
	todos, err := c.TodoService.GetTodos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, todos)
}

func (c *TodoController) GetTodoByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}
	todo, err := c.TodoService.GetTodoByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	ctx.JSON(http.StatusOK, todo)
}

func (c *TodoController) UpdateTodo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}
	todo, err := c.TodoService.GetTodoByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.TodoService.UpdateTodo(todo)
	ctx.JSON(http.StatusOK, todo)
}

func (c *TodoController) DeleteTodo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}
	c.TodoService.DeleteTodo(uint(id))
	ctx.JSON(http.StatusNoContent, nil)
}
