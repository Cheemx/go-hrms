package routes

import (
	"github.com/Cheemx/go-hrms/controllers"
	"github.com/Cheemx/go-hrms/internal/config"
	"github.com/gin-gonic/gin"
)

func StudentRoutes(router *gin.Engine, cfg *config.APIConfig) {
	router.POST("/api/students", controllers.CreateStudent(cfg))
	router.GET("/api/students", controllers.GetStudents(cfg))
	router.GET("/api/students/:id", controllers.GetStudent(cfg))
	router.PUT("/api/students/:id", controllers.UpdateStudent(cfg))
	router.DELETE("/api/students/:id", controllers.DeleteStudent(cfg))
}
