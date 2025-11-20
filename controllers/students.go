package controllers

import (
	"database/sql"
	"errors"
	"net/mail"
	"strings"
	"time"

	"github.com/Cheemx/go-hrms/internal/config"
	"github.com/Cheemx/go-hrms/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateStudent(cfg *config.APIConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request handling
		req := struct {
			Email      string `json:"email"`
			Name       string `json:"name"`
			Department string `json:"department"`
		}{}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			respondWithError(ctx, 500, "error unmarshalling request", err)
			return
		}

		// email validation
		if _, err := mail.ParseAddress(req.Email); err != nil {
			respondWithError(ctx, 400, "invalid email address", err)
			return
		}

		// Creating Student in Database
		err := cfg.DB.CreateStudent(ctx, database.CreateStudentParams{
			Email:      req.Email,
			Name:       req.Name,
			Department: req.Department,
		})
		if err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				respondWithError(ctx, 400, "User already exists", err)
				return
			}
			respondWithError(ctx, 500, "error creating user", err)
			return
		}

		// Retrieving the User for Response
		student, err := cfg.DB.GetStudentByEmail(ctx, req.Email)
		if err != nil {
			respondWithError(ctx, 500, "error retrieving the user just created", err)
			return
		}

		// Creating Response
		res := struct {
			ID         string    `json:"id"`
			Name       string    `json:"name"`
			Department string    `json:"department"`
			Email      string    `json:"email"`
			CreatedAt  time.Time `json:"created_at"`
		}{
			ID:         student.ID,
			Name:       student.Name,
			Department: student.Department,
			Email:      student.Email,
			CreatedAt:  student.CreatedAt,
		}

		ctx.JSON(201, res)
	}
}

func GetStudents(cfg *config.APIConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get all students from table
		students, err := cfg.DB.GetAllStudents(ctx)
		if err != nil {
			respondWithError(ctx, 500, "error retrieving students from DB", err)
		}

		// return array of students
		ctx.JSON(200, students)
	}
}

func GetStudent(cfg *config.APIConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get id parameter from request
		id := ctx.Param("id")

		// parsing UUID
		if _, err := uuid.Parse(id); err != nil {
			respondWithError(ctx, 400, "path param id is not valid uuid", err)
			return
		}

		// Get student from Database
		stud, err := cfg.DB.GetStudentByID(ctx, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				respondWithError(ctx, 400, "No Such ID associated to student", err)
				return
			}
			respondWithError(ctx, 500, "error retrieving student with ID", err)
			return
		}

		// responsing
		ctx.JSON(200, stud)
	}
}

func UpdateStudent(cfg *config.APIConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get id from request parameters
		id := ctx.Param("id")

		// parsing UUID
		if _, err := uuid.Parse(id); err != nil {
			respondWithError(ctx, 400, "path param id is not valid uuid", err)
			return
		}

		// req body to unmarshal to
		req := struct {
			Name       string `json:"name"`
			Department string `json:"department"`
		}{}

		// unmarshall the request accordingly
		if err := ctx.ShouldBindJSON(&req); err != nil {
			respondWithError(ctx, 400, "bad request", err)
			return
		}

		// updating the student in DB
		err := cfg.DB.UpdateStudent(ctx, database.UpdateStudentParams{
			Name:       req.Name,
			Department: req.Department,
			ID:         id,
		})
		if err != nil {
			respondWithError(ctx, 500, "failed to update the student", err)
			return
		}

		// return the response
		stud, err := cfg.DB.GetStudentByID(ctx, id)
		if err != nil {
			respondWithError(ctx, 500, "failed to retrieve the updated student", err)
			return
		}

		// respond with student
		ctx.JSON(200, stud)
	}
}

func DeleteStudent(cfg *config.APIConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get id from request parameters
		id := ctx.Param("id")

		// parsing UUID
		if _, err := uuid.Parse(id); err != nil {
			respondWithError(ctx, 400, "path param id string is not valid uuid", err)
			return
		}

		// Delete the user with id from students
		err := cfg.DB.DeleteStudent(ctx, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				respondWithError(ctx, 401, "no such ID associated with any student", err)
				return
			}
			respondWithError(ctx, 500, "error deleting the student", err)
			return
		}

		ctx.JSON(200, gin.H{"message": "Deletion Successful"})
	}
}
