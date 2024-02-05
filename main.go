package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kushalb-dev/go_crud/controllers"
	"github.com/kushalb-dev/go_crud/initializers"
	"github.com/kushalb-dev/go_crud/migrate"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	migrate.Migrate()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.GET("/posts", controllers.PostsRead)
	r.GET("/posts/:id", controllers.SinglePost)
	r.DELETE("/posts/:id", controllers.PostsDelete)
	r.Run()
}
