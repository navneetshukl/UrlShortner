package main

import (
	"urlshortner/database"
	"urlshortner/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.MigrateDatabase()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/",handlers.Home)
	r.POST("/submit",handlers.LongToShort)
	r.GET("/short/:link",handlers.Short)
	r.Run()
}