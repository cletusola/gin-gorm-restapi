package main 

import ( 

	"github.com/gin-gonic/gin"
	"github.com/cletusola/gin-gorm-api/config"
	"github.com/cletusola/gin-gorm-api/routes"
	
)


// main function
func main(){

	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8080")

}