package main

import (
	"github.com/alexandersantosdev/go-orm-api/controllers"
	"github.com/alexandersantosdev/go-orm-api/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}
func main() {

	r := gin.Default()

	r.POST("/api/posts", controllers.PostCreate)
	r.GET("/api/posts", controllers.GetPosts)
	r.GET("/api/posts/:id", controllers.GetPostById)
	r.PUT("/api/posts/:id", controllers.PostUpdate)
	r.DELETE("/api/posts/:id", controllers.PostDelete)

	r.Run()
}
