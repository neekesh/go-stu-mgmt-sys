package routes

import (
	"learn-go/api/controllers"
	"learn-go/infrastructure"
)

type StudentRoutes struct {
	studentController controllers.StudentControllers
	router            infrastructure.Router
}

func NewStudentRoutes(studentController controllers.StudentControllers, router infrastructure.Router) StudentRoutes {
	return StudentRoutes{
		studentController: studentController,
		router:            router,
	}
}

func (c StudentRoutes) Setup() {
	student := c.router.Gin.Group("student").Use()
	{
		student.GET("/", c.studentController.GetAllStudent)
		student.POST("/create", c.studentController.PostStudent)
		student.PUT("/update/:id", c.studentController.UpdateStudent)
		student.DELETE("/delete/:id", c.studentController.DeleteStudent)
	}
}
