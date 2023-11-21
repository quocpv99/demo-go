package services

import (
	"go_api/models"

	"github.com/jinzhu/gorm"
)

type BlogService struct {
	DB *gorm.DB
}

func NewBlogService(db *gorm.DB) *BlogService {
	return &BlogService{DB: db}
}

func (s *BlogService) GetBlogs() ([]models.Blog, error) {
	var blogs []models.Blog
	if err := s.DB.Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

func (s *BlogService) CreateBlog(blog *models.Blog) {
	s.DB.Create(blog)
}
