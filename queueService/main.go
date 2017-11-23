package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// This handler will match /user/john but will not match neither /user/ or /user
	router.POST("/enqueue/:name/:phoneNumber", enqueue)
	router.Run(":8080")
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
