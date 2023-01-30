package controller

import (
	//"context"
	"fmt"
	"log"
	"net/http"
	//"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rupesh40/CarRental-Go/database"
	helper "github.com/rupesh40/CarRental-Go/helpers"
	"github.com/rupesh40/CarRental-Go/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

var validate = validator.New()


func HashPassword(password string) string{
	bytes,err:= bcrypt.GenerateFromPassword([]byte(password),14)
	if err!=nil{
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword, providedPassword string) (bool,string){
	err :=bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true 
	msg := ""
	if err!=nil{
		msg = fmt.Sprintf("email or password is incorrect")
		check = false 
	}
	return check, msg 
}

func SignUp()gin.HandlerFunc{
return func(c *gin.Context){
	var db =  database.GetDB()
	//var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user models.User

	if err := c.BindJSON(&user);err !=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return 
	}
	if validationErr := validate.Struct(user); validationErr != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error" : validationErr.Error()})
		return 
	}
	result := db.Where("email = ?", user.Email).Find(&user)
	if result.RowsAffected>0{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"this email already exist "})
		return 
	}
	result = db.Where("phone = ?", user.Phone).Find(&user)
	if result.RowsAffected >0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"this phone number already exist "})
		return 
	}
	password := HashPassword(*user.Password)
	user.Password = &password
	user.CreatedAt, _ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
	id := uuid.New()
	user.UserId = id.String()
	token,refreshToken , _ :=helper.GenerateAllTokens(*user.Email,*user.FirstName,*user.LastName,*user.UserType,*&user.UserId)
	user.Token = &token
	user.RefreshToken = &refreshToken

	err := db.Create(&user).Error
	if err != nil {
		msg := fmt.Sprintf("User item was not created")
		c.JSON(http.StatusInternalServerError,gin.H{"error" :msg })
		return 
	}
	
	c.JSON(http.StatusOK, gin.H{"user Insterted ":user})

}
}

func Login() gin.HandlerFunc{
	return func(c *gin.Context){
		var db =  database.GetDB()
		//var ctx, cancel = context.WithTimeout(context.Background(),100* time.Second)
		var user models.User
		var foundUser models.User 
		
		if err := c.BindJSON(&user);err !=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return 
		}
		result:=db.Where("email = ?", user.Email).First(&foundUser)
		//defer cancel()
		
		if result.Error!= nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"email or password is incorrect"})
			return 
		}
		if foundUser.Email ==nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error" : "user not found, please check your Email!"})
			return 
		}
		passwordIsValid,msg := VerifyPassword(*user.Password,*foundUser.Password)
		//defer cancel()
		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
			return 
		}
		token, refreshToken, _ :=helper.GenerateAllTokens(*foundUser.Email, *foundUser.FirstName, *foundUser.LastName, *foundUser.UserType, *&foundUser.UserId)
		// if err !=nil{
		// 	c.JSON(http.StatusInternalServerError,gin.H{"error": err.Err} )
		// 	return 
		// }
		helper.UpdateAllTokens(token,refreshToken,foundUser)
		var newfoundUser models.User
		result =db.Where("email = ?", user.Email).First(&newfoundUser)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"error": result.Error})
			return 
		}
		c.JSON(http.StatusOK, newfoundUser)
	}
}
func Logout()gin.HandlerFunc{
	return func(c *gin.Context){
		var user models.User
		email,b := c.Get("email")
		first_name, b := c.Get("first_name")
		last_name, b  := c.Get("last_name")
		user_type,b := c.Get("user_type")
		user_id,b := c.Get("uid")
		if !b{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"error while getting keys"})
		}
		
		user.UserId = user_id.(string)
		
		token, refreshToken, _ :=helper.GenerateAllTokens(email.(string), first_name.(string), last_name.(string), user_type.(string), user_id.(string))
		fmt.Println()

		helper.UpdateAllTokens(token,refreshToken,user)
		
		
		c.JSON(http.StatusOK,gin.H{"log out success, new token":token})
	}
}
func GetUser()gin.HandlerFunc{
	return func(c *gin.Context){
		var db =  database.GetDB()
		var user models.User 
		userId := c.Param("user_id")

		if err :=helper.MatchUserTypeToUid(c,userId);err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error()})
		}

		result:=db.Where("user_id = ?", userId).First(&user)
		//defer cancel()
		
		if result.Error!= nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"user not found try again"})
			return 
		}
		c.JSON(http.StatusOK,user)
	}
}
