package routes

import (
	"go_api/config"
	"go_api/controllers"
	"go_api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupTodoRoutes() *gin.Engine {
	config.ConnectDB()

	r := gin.Default()

	todoService := services.NewTodoService(config.DB)
	todoController := controllers.NewTodoController(todoService)

	r.LoadHTMLGlob("views/*")
	r.Static("/public", "./public")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	v1 := r.Group("/api/v1")
	{
		v1.POST("/todos", todoController.CreateTodo)
		v1.GET("/todos", todoController.GetTodos)
		v1.GET("/todos/:id", todoController.GetTodoByID)
		v1.PUT("/todos/:id", todoController.UpdateTodo)
		v1.DELETE("/todos/:id", todoController.DeleteTodo)
	}

	return r
}
