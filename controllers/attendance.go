package controllers

import (
	"fmt"
	"strings"

	"github.com/Cheemx/go-hrms/internal/config"
	"github.com/Cheemx/go-hrms/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func MarkAttendance(cfg *config.APIConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Parse the request first in a format
		req := struct {
			StudentID string `json:"student_id"`
			Status    string `json:"status"`
		}{}

		// Unmarshall the request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			respondWithError(ctx, 500, "error unmarshalling request", err)
			return
		}

		// parsing UUID
		if _, err := uuid.Parse(req.StudentID); err != nil {
			respondWithError(ctx, 400, "entered student ID is not UUID", err)
			return
		}

		// mark attendance in attendance table
		err := cfg.DB.MarkAttendance(ctx, database.MarkAttendanceParams{
			StudentID: req.StudentID,
			Status:    database.AttendanceStatus(req.Status),
		})
		if err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				respondWithError(ctx, 400, "Attendance already marked", err)
				return
			}
			respondWithError(ctx, 500, "failed to mark attendance", err)
			return
		}

		// respond with success
		ctx.JSON(200, gin.H{"message": fmt.Sprintf("attendance marked successfully for id: %s\n", req.StudentID)})
	}
}

func GetAttendance(cfg *config.APIConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get student id from path parameter
		studentID := ctx.Param("student_id")

		// parsing UUID
		if _, err := uuid.Parse(studentID); err != nil {
			respondWithError(ctx, 400, "path parameter student ID is not valid UUID", err)
			return
		}

		// get attendance report for the student
		attendance, err := cfg.DB.GetAttendanceByStudentID(ctx, studentID)
		if err != nil {
			respondWithError(ctx, 500, "error retrieving attendance report", err)
			return
		}

		// respond with success
		ctx.JSON(200, attendance)
	}
}
