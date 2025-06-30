package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CatchAllRoutes(r *gin.Engine) {
    r.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
    })

	
	
}
