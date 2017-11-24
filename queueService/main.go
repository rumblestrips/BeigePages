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
		port = "8082"
	}
	port = fmt.Sprintf(":%s", port)

	// This handler will match /user/john but will not match neither /user/ or /user
	router.POST("/enqueue/:name/:phoneNumber", enqueue)
	router.Run(port)
}

func enqueue(c *gin.Context) {
	name := c.Param("name")
	phoneNumber := c.Param("phoneNumber")

	statusCode := push(name, phoneNumber)

	c.Status(statusCode)

}

func push(name string, phoneNumber string) int {
	accountServiceStatusCode := pushToAccountService(name, phoneNumber)
	lookupServiceStatusCode := pushToLookupService(name, phoneNumber)
	if accountServiceStatusCode != http.StatusOK && lookupServiceStatusCode != http.StatusOK {
		return http.StatusBadRequest
	}
	return http.StatusOK
}

func pushToLookupService(name string, number string) int {
	url := fmt.Sprintf("http://localhost:8081/lookupService/%s/%s", name, number)
	resp, err := http.Post(url, "text/plain", strings.NewReader(""))
	if err != nil {
		fmt.Println("Guru Meditation: hfdhjksfhdkjhfdkjashfdsa")
	}
	return resp.StatusCode
}

func pushToAccountService(name string, number string) int {
	url := fmt.Sprintf("http://localhost:8080/accountService/%s/%s", name, number)
	resp, err := http.Post(url, "text/plain", strings.NewReader(""))
	if err != nil {
		fmt.Println("Guru Meditation: hfdhjksfhdkjhfdkjashfdsa")
	}
	return resp.StatusCode
}
