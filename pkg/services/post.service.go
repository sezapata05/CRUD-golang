package services

import "gin_gonic_api/pkg/models"

type PostService interface {
	CreatePost(*models.Post) (*models.Post, error)
	GetPost(*string) (*models.Post, error)
	GetPosts() ([]*models.Post, error)
	UpdatePost(*models.Body, *string) (*models.Post, error)
	DeletePost(*string) error
}
