package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.MaxMultipartMemory = 8 << 20
	engine.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "sb")
		}
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, file.Filename)
	})
	engine.Run()
}
