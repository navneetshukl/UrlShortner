package database

import (
	"fmt"
	"log"
	"os"
	"urlshortner/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase()(*gorm.DB,error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DB_CONNECTION_STRING")

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database ",err)
		return nil,err
	}
	fmt.Println("Connected to Database")
	return DB,nil
}

func MigrateDatabase() {
	DB,err:=ConnectToDatabase()
	if err!=nil{
		log.Fatal("There is error connecting to database ",err)
		return
	}
	DB.AutoMigrate(&models.URL{})
}

func FindURL(short string)(models.URL){
	var url models.URL
	DB,_:=ConnectToDatabase()
	result:=DB.Where("short_url=?", short).Find(&url)
	if result.Error != nil {
		panic("Failed to query the database")
	}

	return url
}

func InsertURL(longurl,shorturl string)(error){
	DB,_:=ConnectToDatabase()
	url:=models.URL{
		LongURL: longurl,
		ShortURL: shorturl,
	}
	err:=DB.Create(&url)
	if err.Error!=nil{
		return err.Error
	}
	return nil
}
