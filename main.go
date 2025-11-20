package main

import (
	"log"

	"github.com/Cheemx/go-hrms/internal/config"
	"github.com/Cheemx/go-hrms/routes"
	"github.com/gin-gonic/gin"
)

const port = "8080"

func main() {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)

	cfg := config.Load()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Human Resource Management System API is working!!!"})
	})

	routes.StudentRoutes(r, cfg)
	routes.AttendanceRoutes(r, cfg)

	log.Printf("Serving Human Resource Management System API on Port: %s\n", port)
	log.Fatal(r.Run(":" + port))
}
