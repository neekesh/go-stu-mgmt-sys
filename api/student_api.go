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

func GetAllStudent(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "get all teh student",
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
	ctx.JSON(200, gin.H{
		"msg": "Update student data",
	})
}
