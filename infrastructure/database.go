package infrastructure

import (
	"fmt"
	"learn-go/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func NewConnectDB() Database {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DBHost"),
		os.Getenv("DBUsername"),
		os.Getenv("DBPassword"),
		os.Getenv("DBName"),
		os.Getenv("DBPort"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	// err = db.AutoMigrate(&models.Student)
	db.AutoMigrate(models.Student{})

	return Database{
		DB: db,
	}
}
