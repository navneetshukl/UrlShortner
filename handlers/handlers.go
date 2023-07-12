package handlers

import (
	"fmt"
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
	url := c.PostForm("url")

	short := uuid.New().String()
	short = "localhost:8080/short/" + short

	c.JSON(200, gin.H{
		"Original_URL": url,
		"Shorten_URL":  short,
	})

	err := database.InsertURL(url, short)
	if err != nil {
		log.Fatal("Error in Inserting into database")
	}

}

func Short(c *gin.Context) {
	link := c.Param("link")
	//localhost:8080/short/8bac290a-2928-426a-8a16-43aced9d800c
	link = "localhost:8080/short/" + link

	var url models.URL

	ans := database.GetFromRedis(link)
	fmt.Println("Answer from Redis is ", ans)
	if ans!=""{
		fmt.Println("redirecting from redis")
		c.Redirect(http.StatusMovedPermanently, ans)
	} else{
	url = database.FindURL(link)
	fmt.Println("Redirecting from PostgreSQL database")
	database.InsertIntoRedis(url.ShortURL, url.LongURL)
	c.Redirect(http.StatusMovedPermanently, url.LongURL)
	}

}
