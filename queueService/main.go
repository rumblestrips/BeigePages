package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}
	port = fmt.Sprintf(":%s", port)

	// This handler will match /user/john but will not match neither /user/ or /user
	router.POST("/enqueue/:name/:phoneNumber", enqueue)
	router.Run(port)
}

func enqueue(c *gin.Context) {
	name := c.Param("name")
	number := c.Param("phoneNumber")

	push(name, number)

	c.Status(http.StatusOK)

}

func push(string, string) {
	// Add hardcoded push calls here.
}
