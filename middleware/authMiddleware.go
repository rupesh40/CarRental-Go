package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	helper "github.com/rupesh40/go-car-rental/V1/helpers"
	"github.com/rupesh40/go-car-rental/V1/database"
	"github.com/rupesh40/go-car-rental/V1/models"

)

func Authenticate() gin.HandlerFunc{

	return func(c * gin.Context){

		clientToken := c.Request.Header.Get("token")
		if clientToken ==""{
			c.JSON(http.StatusInternalServerError,gin.H{"error":fmt.Sprintf("No Autherization header provided / please login again")})
			c.Abort()
			return 
		}
		claims, err := helper.ValidateToken(clientToken)
		if err !=""{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort() 
		}
		var db =  database.GetDB()
		var foundUser models.User
		result:=db.Where("user_id = ?", claims.Uid).First(&foundUser)
			if result.Error!= nil{
				c.JSON(http.StatusInternalServerError, gin.H{"error":"User does not exits, invalid token"})
				c.Abort()
			}
		if clientToken != *foundUser.Token{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"User log out , token expired,"})
			c.Abort()
		}
		 	
		c.Set("email",claims.Email)
		c.Set("first_name", claims.First_name )
		c.Set("last_name", claims.Last_name)
		c.Set("uid",claims.Uid) 
		c.Set("user_type",claims.User_type)
		c.Next()
		
	}
}