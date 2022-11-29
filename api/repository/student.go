package repository

import (
	"learn-go/constants"
	"learn-go/infrastructure"
	"learn-go/models"
)

type StudentRepository struct {
	database infrastructure.Database
}

func NewStudentRepository(db infrastructure.Database) StudentRepository {
	return StudentRepository{
		database: db,
	}
}

func (c StudentRepository) Create(Student models.Student) error {
	return c.database.DB.Create(&Student).Error
}

func (c StudentRepository) Delete(Student models.Student) error {
	return c.database.DB.Delete(&Student).Error
}

func (c StudentRepository) CheckStudent(id string, students *models.Student) error {
	return c.database.DB.Where("id=?", id).First(students).Error
}

func (c StudentRepository) Update(
	input constants.UpdateStudentInput,
	students *models.Student,
) error {
	return c.database.DB.Model(students).Updates(input).Error
}
