package routes

import (
	"job-app/main/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterJobRoutes(router *gin.Engine) {
	router.GET("/jobs", controllers.GetJobs)
	router.POST("/jobs", controllers.AddBook)
	router.PATCH("/jobs/:id", controllers.EditJob)
	router.DELETE("/jobs/:id", controllers.DeleteJob)

}
