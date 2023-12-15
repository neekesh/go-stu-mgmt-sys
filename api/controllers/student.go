package controllers

import (
	"learn-go/api/repository"
	"learn-go/models"
	"learn-go/types"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentControllers struct {
	student repository.StudentRepository
}

func NewStudentControllers(
	student repository.StudentRepository,
) StudentControllers {
	return StudentControllers{
		student: student,
	}
}

func (cc StudentControllers) GetAllStudent(ctx *gin.Context) {
	var students []models.Student
	db := ctx.MustGet("db").(*gorm.DB)

	db.Find(&students)

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "get all teh student",
		"data": students,
	})
}

func (cc StudentControllers) PostStudent(ctx *gin.Context) {
	var newStudent types.StudentCreate
	// db := ctx.MustGet("db").(*gorm.DB)

	if err := ctx.BindJSON(&newStudent); err != nil {
		log.Print("Error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student := models.Student{
		FullName: newStudent.FullName,
		Age:      newStudent.Age,
		Address:  newStudent.Address,
		Major:    newStudent.Major,
		Grade:    newStudent.Grade,
	}
	cc.student.Create(student)
	ctx.JSON(http.StatusCreated, gin.H{
		"msg":  "new student craete",
		"data": student,
	})
}

func (cc StudentControllers) DeleteStudent(ctx *gin.Context) {
	var students models.Student
	// db := ctx.MustGet("db").(*gorm.DB)

	if err := cc.student.CheckStudent(ctx.Param("id"), &students); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Student Doesnot exists",
		})
	}

	cc.student.Delete(students)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "student deleted",
	})
}

func (cc StudentControllers) UpdateStudent(ctx *gin.Context) {
	// var student models.Student
	// db := ctx.MustGet("db").(*gorm.DB)
	var input types.UpdateStudentInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := strconv.Atoi(ctx.Param("id"))
	student := models.Student{
		Id:       uint64(id),
		FullName: input.FullName,
		Age:      input.Age,
		Address:  input.Address,
		Major:    input.Major,
		Grade:    input.Grade,
	}

	if err := cc.student.Update(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Update student data",
		"data": student,
	})
}
