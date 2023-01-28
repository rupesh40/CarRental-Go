package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/rupesh40/go-car-rental/V1/controllers"
)
func AuthRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("user/signup", controller.SignUp())
	incomingRoutes.POST("user/login", controller.Login())
}
