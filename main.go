package main

import (
	"go-curd/controllers"
	"go-curd/initializers"

	"github.com/gin-gonic/gin"
)
func init(){
	initializers.LoadEnvVariables()
	initializers.Connection()
}

func main() {
	r := gin.Default()
	//create post
	r.POST("/", controllers.CreatePost)

	//get posts
	r.GET("/", controllers.GetPosts)

	r.GET("/:id", controllers.GetPost)

	r.PUT("/:id", controllers.UpdatePost)

	//delete post
	r.DELETE("/:id", controllers.DeletePost)

	r.Run() 
}