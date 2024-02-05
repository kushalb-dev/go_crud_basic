package controllers

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/kushalb-dev/go_crud/initializers"
	"github.com/kushalb-dev/go_crud/models"
)

func PostsCreate(c *gin.Context) {

	// Get Data Off Request Body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create a Post
	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{"post": post})
}

func PostsRead(c *gin.Context) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		log.Fatal("Couldn't retrieve post from DB.")
	}

	c.JSON(200, gin.H{"post": posts})
}

func SinglePost(c *gin.Context) {

	// Getting the id off the URL
	id := c.Param("id")

	// Storing the post
	var post models.Post
	initializers.DB.First(&post, id)

	// Sending back data
	c.JSON(200, gin.H{"post": post})
}

func PostsUpdate(c *gin.Context) {

	// finding id to update from URL
	id := c.Param("id")

	// getting the post
	var post models.Post
	initializers.DB.First(&post, id)

	// creating parameters to update post
	var body struct {
		Body  string
		Title string
	}

	// binding the to be updated paramters to post
	c.Bind(&body)

	// updating post in database
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(200, gin.H{"postUpdate": post})
}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	initializers.DB.Delete(&post, id)

	c.Status(200)
}
