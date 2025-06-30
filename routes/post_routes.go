package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/su-de-sh/gin-api/controllers"
)

func PostRoutes(r *gin.Engine) {
    r.POST("/posts", controllers.CreatePost)
    r.GET("/posts", controllers.GetPosts)
    r.GET("/posts/:id", controllers.GetPost)
    r.PUT("/posts/:id", controllers.UpdatePost)
    r.DELETE("/posts/:id", controllers.DeletePost)
}