package routes

import (
	"github.com/Cheemx/go-hrms/controllers"
	"github.com/Cheemx/go-hrms/internal/config"
	"github.com/gin-gonic/gin"
)

func AttendanceRoutes(router *gin.Engine, cfg *config.APIConfig) {
	router.POST("/api/attendance/mark", controllers.MarkAttendance(cfg))
	router.GET("/api/attendance/:student_id", controllers.GetAttendance(cfg))
}
