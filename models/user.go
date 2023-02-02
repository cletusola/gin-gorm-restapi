package models 

import(
	"gorm.io/gorm"
)

// user schema 
type User struct {
	gorm.Model
	Id 			int 		`json:"ID" gorm:"primary_key"`
	Name 		string 		`json:"name"`
	Email 		string 		`json:"email"`
	Password 	string 		`json:"password"`
}