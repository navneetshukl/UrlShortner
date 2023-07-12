package handlers

import (

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(200, "home.html", gin.H{})
}

func LongToShort(c *gin.Context) {
	url:=c.PostForm("url")
	c.JSON(200,gin.H{
		"Long Url":url,
	})
}
