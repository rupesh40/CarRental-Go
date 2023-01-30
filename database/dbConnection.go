package database 

import (
	"fmt"
	"log"
	//"time"
	"os"
	//"context"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/rupesh40/CarRental-Go/models"
	

)
var (
	db  *gorm.DB
	err error
)
type Config struct{
	Host	string 
	Port 	string 
	Password string 
	User	string  
	DBName  string 
	SSLMode string

} 
func Init() {
	con := GetEnv()
	db, err = gorm.Open(postgres.Open(con),&gorm.Config{DisableForeignKeyConstraintWhenMigrating: true,})
	if err != nil {
		log.Fatal("coudl not establish db connection", err)
		panic(err) 
	}
	fmt.Println()
	autoMigration()
}

func GetDB() *gorm.DB {
	return db
}
func GetEnv() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err)
	}
	config := &Config{
		Host 	:   os.Getenv("DB_HOST"),
		Port 	: 	"5432",
		Password: 	os.Getenv("DB_PASS"),
		User 	: 	os.Getenv("DB_USER"),
		SSLMode : 	"disable",
		DBName 	: 	os.Getenv("DB_NAME"),
	}
	dsn:= fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port,config.User , config.Password,config.DBName, config.SSLMode,
	)
	return dsn
}
func autoMigration() {
	var models = []interface{}{&models.User{},&models.Car{}}
	err := db.AutoMigrate(models...)
	if err!= nil{
		log.Panic(err.Error())
	}
	log.Print("Migration is Succesful")

}
 