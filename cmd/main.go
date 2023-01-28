package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/rupesh40/go-car-rental/V1/routes"
	"github.com/rupesh40/go-car-rental/V1/database"

)

func main(){
	database.Init()
	port:= os.Getenv("SERVER_PORT")
	if port ==""{
		port = "3000"
	}
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.Routes(router)
	routes.SetUpGraphqlRoutes(router)

	router.Run(":"+ port)

}