package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// This handler will match /user/john but will not match neither /user/ or /user
	router.POST("/register/:name/:phoneNumber", register)
	router.Run(":8083")
}

func register(c *gin.Context) {

	name := c.Param("name")
	number := c.Param("phoneNumber")

	statusCode := push(name, number)
	c.Status(statusCode)
}

func push(name string, number string) int {
	url := fmt.Sprintf("http://localhost:8082/enqueue/%s/%s", name, number)
	resp, err := http.Post(url, "text/plain", strings.NewReader(""))
	if err != nil {
		fmt.Println("Guru Meditation: hfdhjksfhdkjhfdkjashfdsa")
	}
	return resp.StatusCode
}
