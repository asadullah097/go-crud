package controllers

import (
	"go-curd/initializers"
	"go-curd/models"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
     //get request body
	  var body struct {
		Title string
		Body  string
	  }
	  c.Bind(body)
	// //  //create post
	 post :=models.Post{Title:body.Title,Body:body.Body}

	 initializers.DB.Create(&post)
	c.JSON(200, gin.H{
		"message": "Post created successfully!",
		"data": post,
	})
}

func GetPosts(c *gin.Context) {
	posts := []models.Post{}
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"message": "All posts retrieved successfully!",
		"data": posts,
	})
}

func GetPost(c *gin.Context) {
	post:=models.Post{}
	initializers.DB.First(&post, c.Param("id"))
	c.JSON(200, gin.H{
		"message": "Post retrieved successfully!",
		"data": post,
	})
}

func UpdatePost(c *gin.Context) {
	var id = c.Param("id")

	//get request body
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)
	post:=models.Post{}
	//find post
	initializers.DB.First(&post, id)
	post.Title = body.Title
	post.Body = body.Body
	//Update
	initializers.DB.Save(&post)
	c.JSON(200, gin.H{
		"message": "POST updated successfully!",
		"data": post,
	})
}
func DeletePost(c *gin.Context) {
	post:=models.Post{}
    result:=initializers.DB.First(&post, c.Param("id"))
	if result != nil {
		c.JSON(404, gin.H{
			"message": "post not found",
		})
		
	}
	initializers.DB.Delete(&post,c.Param("id"))

	c.JSON(200, gin.H{
		"message": "POST deleted successfully!",
		"data": post,
	})
}
