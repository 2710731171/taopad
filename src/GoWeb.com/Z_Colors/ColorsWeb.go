package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type myForm struct {
	Colors []string `form:"colors[]"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("htmlWeb/*")
	r.GET("/colors", func(c *gin.Context) {
		c.HTML(http.StatusOK, "colors.html", gin.H{
			"title": "选择颜色",
		})
	})
	r.POST("/colors", func(c *gin.Context) {
		var fakeForm myForm
		c.ShouldBind(&fakeForm)
		c.JSON(200, gin.H{"color": fakeForm.Colors})
	})
	r.Run()
}
