package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"fmt"
)

type BindFile struct {
	File  *multipart.FileHeader `form:"file" binding:"required"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/transcribe", func(c *gin.Context) {
		var bindFile BindFile

		// Bind file
		if err := c.ShouldBind(&bindFile); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
			return
		}

		file, err := bindFile.File.Open()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		
		defer file.Close()

		text, err := audio_transcribe(file)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"result": text,
			})
		}
	})

	return r
}

func main() {
	init_model()
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

