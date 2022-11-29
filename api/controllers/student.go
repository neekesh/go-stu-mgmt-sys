package controllers

import (
	"learn-go/api/repository"
	"learn-go/constants"
	"learn-go/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentControllers struct {
	student repository.StudentRepository
}

func NewStudentControllers(student repository.StudentRepository) StudentControllers {
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
	var newStudent constants.StudentCreate
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
	var student models.Student
	// db := ctx.MustGet("db").(*gorm.DB)

	if err := cc.student.CheckStudent(ctx.Param("id"), &student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record NOt Founds"})
		return
	}

	var input constants.UpdateStudentInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cc.student.Update(input, &student)

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Update student data",
		"data": student,
	})
}
