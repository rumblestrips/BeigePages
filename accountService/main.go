package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	accountRegistry := make(map[string]string)

	router := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	port = fmt.Sprintf(":%s", port)

	// This handler will match /user/john but will not match neither /user/ or /user
	router.POST("/accountService/:name/:phoneNumber", register(accountRegistry))
	router.Run(port)
}

func register(accountRegistry map[string]string) func(*gin.Context) {
	return func(c *gin.Context) {
		accountRegistry[c.Param("name")] = c.Param("phoneNumber")

		c.Status(http.StatusOK)
	}
}
