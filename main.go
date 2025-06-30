package main

import (
	"github.com/gin-gonic/gin"
	"github.com/su-de-sh/gin-api/database"
	"github.com/su-de-sh/gin-api/routes"
)

func main() {
     database.Connect()
    r := gin.Default()
    routes.PostRoutes(r)
    routes.CatchAllRoutes(r)
    r.Run() // Use PORT from .env or default to 3000
}
