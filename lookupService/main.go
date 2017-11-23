package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	phoneRegistry := make(map[string]string)

	router := gin.Default()

	// This handler will match /user/john but will not match neither /user/ or /user
	router.GET("/lookup/:name", lookup(phoneRegistry))
	router.POST("/register/:name/:phoneNumber", register(phoneRegistry))
	router.Run(":8080")
}

func lookup(phoneRegistry map[string]string) func(*gin.Context) {
	return func(c *gin.Context) {
		phoneNumber, ok := phoneRegistry[c.Param("name")]

		if ok {
			c.JSON(http.StatusOK, gin.H{
				"phoneNumber": phoneNumber,
			})
		} else {
			c.Status(http.StatusNotFound)
		}
	}
}

func register(phoneRegistry map[string]string) func(*gin.Context) {
	return func(c *gin.Context) {
		phoneRegistry[c.Param("name")] = c.Param("phoneNumber")

		c.Status(http.StatusOK)
	}
}
