package config 

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/cletusola/gin-gorm-api/models"
)

// setting DB as a global variable
var DB *gorm.DB

// db connection function
func Connect(){

	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"), &gorm.Config{})

	if err != nil {
		panic(err)
	}	

	// auto migrate database
	db.AutoMigrate(&models.User{})
	DB = db
}
