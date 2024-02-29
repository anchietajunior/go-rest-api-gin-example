package main

import (
	routes "github.com/anchietajunior/go-rest-api/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	routes.AppRoutes(app)

	app.Run("localhost:8000")
}
