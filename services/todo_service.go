package services

import (
	"go_api/models"

	"github.com/jinzhu/gorm"
)

type TodoService struct {
	DB *gorm.DB
}

func NewTodoService(db *gorm.DB) *TodoService {
	return &TodoService{DB: db}
}

func (s *TodoService) CreateTodo(todo *models.Todo) {
	s.DB.Create(todo)
}

func (s *TodoService) GetTodos() ([]models.Todo, error) {
	var todos []models.Todo
	if err := s.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *TodoService) GetTodoByID(id uint) (*models.Todo, error) {
	var todo models.Todo
	if err := s.DB.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (s *TodoService) UpdateTodo(todo *models.Todo) {
	s.DB.Save(todo)
}

func (s *TodoService) DeleteTodo(id uint) {
	var todo models.Todo
	s.DB.Delete(&todo, id)
}
