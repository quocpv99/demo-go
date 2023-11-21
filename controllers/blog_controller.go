package controllers

import (
	"fmt"
	"go_api/models"
	"go_api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	BlogService *services.BlogService
}

func NewBlogController(blogService *services.BlogService) *BlogController {
	return &BlogController{BlogService: blogService}
}

func (c *BlogController) GetBlogs(ctx *gin.Context) {
	todos, err := c.BlogService.GetBlogs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, todos)
}

func (c *BlogController) CreateBlog(ctx *gin.Context) {
	var blog models.Blog
	if err := ctx.ShouldBindJSON(&blog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Received JSON: %+v\n", ctx)
	c.BlogService.CreateBlog(&blog)
	ctx.JSON(http.StatusCreated, blog)
}
