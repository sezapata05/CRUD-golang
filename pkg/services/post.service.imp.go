package services

import (
	"gin_gonic_api/pkg/models"

	"gorm.io/gorm"
)

type PostServiceImpl struct {
	DB *gorm.DB
}

func NewPostService(db *gorm.DB) PostService {
	return &PostServiceImpl{
		DB: db,
	}
}

// CreatePost implements PostService
func (p *PostServiceImpl) CreatePost(post *models.Post) (*models.Post, error) {
	result := p.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		return nil, result.Error
	}

	return post, nil
}

// DeletePost implements PostService
func (p *PostServiceImpl) DeletePost(id *string) error {
	// delete the post
	p.DB.Delete(&models.Post{}, id)

	return nil
}

// GetPost implements PostService
func (p *PostServiceImpl) GetPosts() ([]*models.Post, error) {
	var posts []*models.Post

	p.DB.Find(&posts)

	return posts, nil
}

// GetPosts implements PostService
func (p *PostServiceImpl) GetPost(id *string) (*models.Post, error) {
	var post models.Post

	p.DB.First(&post, id)

	return &post, nil
}

// UpdatePost implements PostService
func (p *PostServiceImpl) UpdatePost(body *models.Body, id *string) (*models.Post, error) {
	// Find the post were updating
	//get the posts
	var post models.Post
	// Get all records
	p.DB.First(&post, id)

	// update it
	p.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	return &post, nil
}
