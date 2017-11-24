package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	accountRegistry := make(map[string]string)

	router := gin.Default()

	// This handler will match /user/john but will not match neither /user/ or /user
	router.POST("/accountService/:name/:phoneNumber", register(accountRegistry))
	router.Run(":8080")
}

func register(accountRegistry map[string]string) func(*gin.Context) {
	return func(c *gin.Context) {
		accountRegistry[c.Param("name")] = c.Param("phoneNumber")

		c.Status(http.StatusOK)
	}
}
