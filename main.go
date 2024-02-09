package main

import (
	"github.com/gin-gonic/gin"

	"api/src/routes"
)

func main() {
	app := gin.Default()
	routes.AppRoutes(app)

	app.Run("localhost:3001")
}
