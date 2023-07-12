package handlers

import (
	"log"
	"net/http"
	"urlshortner/database"
	"urlshortner/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Home(c *gin.Context) {
	c.HTML(200, "home.html", gin.H{})
}

func LongToShort(c *gin.Context) {
	url:=c.PostForm("url")
	c.JSON(200,gin.H{
		"Long Url":url,
	})

	short:=uuid.New().String()
	short = "localhost:8080/short/"+ short

	/*Url:=models.URL{
		LongURL: url,
		ShortURL: short,
	}*/
	err:=database.InsertURL(url,short)
	if err!=nil {
		log.Fatal("Error in Inserting into database")
	}
	
}

func Short(c *gin.Context){
	link:=c.Param("link")
	//localhost:8080/short/8bac290a-2928-426a-8a16-43aced9d800c
	link="localhost:8080/short/"+link

	var url models.URL

	url=database.FindURL(link)
	c.Redirect(http.StatusMovedPermanently,url.LongURL)

}
