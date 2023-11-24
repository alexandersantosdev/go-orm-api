package controllers

import (
	"github.com/alexandersantosdev/go-orm-api/initializers"
	"github.com/alexandersantosdev/go-orm-api/models"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(201, gin.H{
		"error": "false",
		"post":  post,
	})
}

func GetPosts(c *gin.Context) {

	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"error": "false",
		"posts": posts,
	})
}

func GetPostById(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"error": "false",
		"post":  post,
	})
}

func PostUpdate(c *gin.Context) {

	id := c.Param("id")

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)
	post.Body = body.Body
	post.Title = body.Title
	initializers.DB.Save(&post)

	c.JSON(200, gin.H{
		"error": "false",
		"post":  post,
	})

}

func PostDelete(c *gin.Context) {

	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.JSON(204, gin.H{
		"error": "false",
	})
}
