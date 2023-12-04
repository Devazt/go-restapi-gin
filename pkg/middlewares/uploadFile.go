package middlewares

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		src, err := file.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		defer src.Close()

		tempFile, err := ioutil.TempFile("uploads", "image-*.*")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		defer tempFile.Close()

		if _, err = io.Copy(tempFile, src); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		data := tempFile.Name()
		filename := data[8:]

		c.Set("dataFile", filename)
		next(c)
	}
}
