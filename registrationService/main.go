package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}
	port = fmt.Sprintf(":%s", port)

	// This handler will match /user/john but will not match neither /user/ or /user
	router.POST("/register/:name/:phoneNumber/:postCode/:email", register)
	router.Run(port)
}

func register(c *gin.Context) {

	name := c.Param("name")
	phoneNumber := c.Param("phoneNumber")
	postCode := c.Param("postCode")
	email := c.Param("email")

	statusCode := push(name, phoneNumber, postCode, email)
	c.Status(statusCode)
}

func push(name string, phoneNumber string, postCode string, email string) int {
	queueServiceURL := os.Getenv("QUEUE_SERVICE_URL")
	if queueServiceURL == "" {
		queueServiceURL = "http://localhost:8082"
	}

	url := fmt.Sprintf("%s/enqueue/%s/%s/%s/%s", queueServiceURL, name, phoneNumber, email)
	resp, err := http.Post(url, "text/plain", strings.NewReader(""))
	if err != nil {
		fmt.Println("Guru Meditation: registrationService")
	}
	return resp.StatusCode
}
