package api

import (
	"learn-go/infrastructure"
	"learn-go/models"
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
)

// var DB = infrastructure.DB

type StudentCreate struct {
	FullName string `json:"full_name" binding:"required"`
	Age      uint8  `json:"age,string" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Major    string `json:"major" binding:"required"`
	Grade    string `json:"grade"  binding:"required"`
}

type UpdateStudentInput struct {
	FullName string `json:"full_name"`
	Age      uint8  `json:"age,string"`
	Address  string `json:"address"`
	Major    string `json:"major"`
	Grade    string `json:"grade"`
}

func GetAllStudent(ctx *gin.Context) {
	var students []models.Student

	infrastructure.DB.Find(&students)

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "get all teh student",
		"data": students,
	})
}

func PostStudent(ctx *gin.Context) {
	var newStudent StudentCreate

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

	infrastructure.DB.Create(&student)
	ctx.JSON(http.StatusCreated, gin.H{
		"msg":  "new student craete",
		"data": student,
	})
}

func DeleteStudent(ctx *gin.Context) {
	ctx.JSON(204, gin.H{
		"msg": "Delete student",
	})
}

func UpdateStudent(ctx *gin.Context) {
	var student models.Student
	log.Println(student)
	log.Println("Am i here", ctx.Param("id"))
	if err := infrastructure.DB.Where("id = ?", ctx.Param("id")).First(&student).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record NOt Founds"})
		return
	}

	var input UpdateStudentInput
	log.Println("checked here", ctx.Param("id"))

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("before update", ctx.Param("id"))
	infrastructure.DB.Model(student).Updates(input)
	log.Println("after uupdate", ctx.Param("id"))

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Update student data",
		"data": student,
	})
}
