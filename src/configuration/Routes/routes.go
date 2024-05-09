package routes

import (
	controller "github.com/GabrielBruno7/go-crud-study/src/configuration/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(route *gin.RouterGroup) {
	route.GET("/findUserById/:id", controller.FindUserById)
	route.GET("/findUserByEmail/:email", controller.FindUserByEmail)
	route.POST("/user", controller.CreateUser)
	route.PUT("/user/:id", controller.UpdateUser)
	route.DELETE("/user/:id", controller.DeleteUser)
}