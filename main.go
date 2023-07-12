package main

import (
	"fmt"
	"urlshortner/database"
	"urlshortner/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	DB,_:=database.ConnectToDatabase()
	database.MigrateDatabase()
	fmt.Println(DB)

	r.LoadHTMLGlob("templates/*.html")
	r.GET("/",handlers.Home)
	r.POST("/submit",handlers.LongToShort)
	r.GET("/short/:link",handlers.Short)
	r.Run()
}