package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// This handler will match /user/john but will not match neither /user/ or /user
	router.POST("/register/:name/:phoneNumber", register)
	router.Run(":8083")
}

func register(c *gin.Context) {

	//name := c.Param("name")
	//number := c.Param("phoneNumber")
	//url := fmt.Sprintf("http://localhost:8080/enqueue/%s/%s", name, number)
	//http.Post(url)

	c.Status(http.StatusOK)
}
