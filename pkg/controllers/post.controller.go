package controllers

import (
	"gin_gonic_api/pkg/models"
	"gin_gonic_api/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	PostService services.PostService
}

func NewPost(postservice services.PostService) PostController {
	return PostController{
		PostService: postservice,
	}
}

func (pc *PostController) CreatePost(c *gin.Context) {
	var body models.Body

	c.Bind(&body)

	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	post_result, err := pc.PostService.CreatePost(&post)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post_result,
	})

}

func (pc *PostController) GetPosts(c *gin.Context) {

	posts, err := pc.PostService.GetPosts()

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func (pc *PostController) GetPost(c *gin.Context) {
	// get id from req
	id := c.Param("id")

	post, err := pc.PostService.GetPost(&id)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func (pc *PostController) UpdatePost(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	var body models.Body
	// Get data off req body
	c.Bind(&body)

	post_result, err := pc.PostService.UpdatePost(&body, &id)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// respond with it
	c.JSON(http.StatusOK, gin.H{
		"post": post_result,
	})
}

func (pc *PostController) DeletePost(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	err := pc.PostService.DeletePost(&id)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	//response
	c.Status(http.StatusOK)
}

func (pc *PostController) RegisterPostRoutes(rg *gin.RouterGroup) {
	postroute := rg.Group("/postservice")
	postroute.POST("/post", pc.CreatePost)
	postroute.GET("/posts", pc.GetPosts)
	postroute.GET("/post/:id", pc.GetPost)
	postroute.PUT("/post/:id", pc.UpdatePost)
	postroute.DELETE("/post/:id", pc.DeletePost)
}
