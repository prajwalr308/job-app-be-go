package main

import (
	"job-app/main/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.RegisterJobRoutes(router)

	router.Run(":8080")
}
