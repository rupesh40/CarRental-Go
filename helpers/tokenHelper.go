package helper 

import (
	//"context"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/rupesh40/go-car-rental/V1/database"
	"github.com/rupesh40/go-car-rental/V1/models"
	jwt "github.com/dgrijalva/jwt-go"

)
type SignedDetails struct{
	Email				string 
	First_name			string
	Last_name 			string
	Uid 				string
	User_type 			string
	jwt.StandardClaims	 
}


var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email,firstName,lastName,userType,uid string)(signedToken,signedRefreshToken string, err error){
	claims := &SignedDetails{
		Email :     email,
		First_name: firstName,
		Last_name:  lastName,
		Uid: 		uid,
		User_type:  userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt : time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	refreshClaims :=  &SignedDetails{
		StandardClaims : jwt.StandardClaims{
			ExpiresAt : time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err !=nil {
		log.Panic(err)
		return  
	}
	return token, refreshToken, err
}
func ValidateToken(signedToken string)(claims *SignedDetails,msg string){
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token)(interface{},error){
			return []byte(SECRET_KEY), nil 
		},
	)
	if err != nil{
		msg = err.Error()
		return 
	}
	claims,ok := token.Claims.(*SignedDetails)
	if !ok{
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return 
	}
	if claims.ExpiresAt < time.Now().Local().Unix(){
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}
	return claims, msg
}
func UpdateAllTokens(signedToken, signedRefreshToken string, user models.User){
	var db =  database.GetDB()
	Token :=  &signedToken
	Refresh_token := &signedRefreshToken
	Updated_at, _ := time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
	db.Model(&user).Where("user_id = ?", user.UserId).Update("token",Token) 
	db.Model(&user).Where("user_id = ?", user.UserId).Update("refresh_token",Refresh_token)
	db.Model(&user).Where("user_id = ?", user.UserId).Update("updated_at", Updated_at)  
	
	//db.Save(&user)
	return 
}