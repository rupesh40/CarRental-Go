package controller

import (
	//"context"
	"fmt"
	//"log"
	"net/http"
	//"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rupesh40/go-car-rental/V1/database"
	"github.com/rupesh40/go-car-rental/V1/models"
	"github.com/google/uuid"
)

func Book()gin.HandlerFunc{
	return func(c *gin.Context){
		var db = database.GetDB()
		var car models.Car
		carID := c.Params.ByName("car_id")
		if carID == ""{
			c.JSON(400,gin.H{"error":"Car ID cannot be empty"})
			return
		}
		result:=db.Where("id = ?", carID).First(&car)
		if result.Error!= nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Incorrect carId, try again"})
			return 
		}
		if car.Status =="Available"{
			db.Model(&car).Update("status", "Rented")
			c.JSON(200,car)

		}else{
			c.JSON(400,gin.H{
				"message" : "This Vehicle is not currently available",
			})
		}
	
	}
}
func Return()gin.HandlerFunc{
	return func(c *gin.Context){
		var db = database.GetDB()
		var car models.Car
		carID := c.Params.ByName("car_id")
		if carID == ""{
			c.JSON(400,gin.H{"error":"Car ID cannot be empty"})
			return
		}
		result:=db.Where("id = ?", carID).First(&car)
		if result.Error!= nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Incorrect carId, try again"})
			return 
		}
		if car.Status =="Rented"{
			db.Model(&car).Updates(map[string]interface{}{"status": "Available", "last_used_date": time.Now().String()})			
			
			c.JSON(200,car)

		}else{
			c.JSON(400,gin.H{
				"message" : "Car for this ID is not rented,check car_ID, please report the issue",
			})
		}
	}
}
func AddCar()gin.HandlerFunc{
	 return func(c *gin.Context){
		var db = database.GetDB()

		var car models.Car

		if err:= c.BindJSON(&car);err !=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		}	
		if validationErr := validate.Struct(car); validationErr != nil {
			c.JSON(http.StatusBadRequest,gin.H{"error" : validationErr.Error()})
			return 
		}
		id := uuid.New()
		car.ID = id.String()
		err := db.Create(&car).Error
		if err != nil {
			msg := fmt.Sprintf("Car item was not created")
			c.JSON(http.StatusInternalServerError,gin.H{"error" :msg })
			return 
		}
	
		c.JSON(http.StatusOK, gin.H{"car Insterted ": car})
	}
}

func GetAll()gin.HandlerFunc{
	 return func(c *gin.Context){
		var db =  database.GetDB()
		var cars []models.Car 

		result := db.Find(&cars)
		
		if result.Error!= nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"user not found try again"})
			return 
		}
		c.JSON(http.StatusOK,cars)
	}
}