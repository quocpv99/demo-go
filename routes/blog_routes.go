package routes

import (
	"go_api/config"
	"go_api/controllers"
	"go_api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupBlogRoutes() *gin.Engine {
	config.ConnectDB()

	r := gin.Default()

	blogService := services.NewBlogService(config.DB)
	blogController := controllers.NewBlogController(blogService)

	r.LoadHTMLGlob("views/*")
	r.Static("/public", "./public")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	v1 := r.Group("/api/v1")
	{
		v1.POST("/blogs", blogController.CreateBlog)
		v1.GET("/blogs", blogController.GetBlogs)
	}
	return r
}
