package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/rupesh40/CarRental-Go/controllers"
	"github.com/rupesh40/CarRental-Go/middleware"
)
func Routes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	
	incomingRoutes.GET("/users/:user_id",controller.GetUser())
	incomingRoutes.GET("users/logout",controller.Logout())

	incomingRoutes.POST("cars/AddCar", controller.AddCar())
	incomingRoutes.GET("cars/bookCar/:car_id", controller.Book())
	incomingRoutes.GET("cars/returnCar/:car_id", controller.Return())
	incomingRoutes.GET("cars/", controller.GetAll())
}
