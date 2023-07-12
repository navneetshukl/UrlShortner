package main

import (
	"fmt"
	"urlshortner/database"
	"urlshortner/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")

	r.GET("/json",func(c *gin.Context){
		c.JSON(200,gin.H{
			"FirstName":"Navneet",
			"LastName":"Shukla",
		})
	})
	r.GET("/",handlers.Home)
	r.POST("/submit",handlers.LongToShort)
	DB,_:=database.ConnectToDatabase()
	database.MigrateDatabase()
	fmt.Println(DB)

	r.Run()
}