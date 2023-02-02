package routes 


import(
	"github.com/gin-gonic/gin"
	"github.com/cletusola/gin-gorm-api/controller"
)

// routes for all actions 
func UserRoute(router *gin.Engine){
	
	router.GET("/", controller.GetUsers)
	router.POST("/", controller.CreateUser)
	router.PUT("/:id", controller.UpdateUser)
	router.DELETE("/:id", controller.DeleteUser)
}