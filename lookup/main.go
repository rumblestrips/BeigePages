package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// This handler will match /user/john but will not match neither /user/ or /user
	router.GET("/lookup/:name", func(c *gin.Context) {
		//name := c.Param("name")
		phoneNumber := "0123456"
		c.JSON(http.StatusOK, gin.H{
			"phoneNumber": phoneNumber,
		})
	})

	router.Run(":8080")
}
