package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/su-de-sh/gin-api/database"
	"github.com/su-de-sh/gin-api/models"
)

func CreatePost(c *gin.Context) {
    var post models.Post
    if err := c.ShouldBindJSON(&post); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    database.DB.Create(&post)
    c.JSON(200, post)
}

func GetPosts(c *gin.Context) {
    var posts []models.Post
    database.DB.Find(&posts)
    c.JSON(200, posts)
}

func GetPost(c *gin.Context) {
    id := c.Param("id")
    var post models.Post
    if err := database.DB.First(&post, id).Error; err != nil {
        c.JSON(404, gin.H{"error": "Post not found"})
        return
    }
    c.JSON(200, post)
}

func UpdatePost(c *gin.Context) {
    id := c.Param("id")
    var post models.Post
    if err := database.DB.First(&post, id).Error; err != nil {
        c.JSON(404, gin.H{"error": "Post not found"})
        return
    }

    var input models.Post
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    database.DB.Model(&post).Updates(input)
    c.JSON(200, post)
}

func DeletePost(c *gin.Context) {
    id := c.Param("id")
    database.DB.Delete(&models.Post{}, id)
    c.JSON(204, gin.H{})
}