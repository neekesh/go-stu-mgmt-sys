package api

import "github.com/gin-gonic/gin"

func GetStudent(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "get all teh student",
	})
}

func CreateStudent(ctx *gin.Context) {
	ctx.JSON(201, gin.H{
		"msg": "create new student",
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
