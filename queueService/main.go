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
	router.POST("/enqueue/:name/:phoneNumber/:postCode/:email", enqueue)
	router.Run(port)
}

func enqueue(c *gin.Context) {
	name := c.Param("name")
	phoneNumber := c.Param("phoneNumber")
	postCode := c.Param("postCode")
	email := c.Param("email")

	statusCode := push(name, phoneNumber, postCode, email)

	c.Status(statusCode)

}

func push(name string, phoneNumber string postCode string, email string) int {
	accountServiceStatusCode := pushToAccountService(name, phoneNumber, postCode, email)
	lookupServiceStatusCode := pushToLookupService(name, phoneNumber)
	if accountServiceStatusCode != http.StatusOK && lookupServiceStatusCode != http.StatusOK {
		return http.StatusBadRequest
	}
	return http.StatusOK
}

func pushToLookupService(name string, number string) int {
	lookupServiceURL := os.Getenv("LOOKUP_SERVICE_URL")
	if lookupServiceURL == "" {
		lookupServiceURL = "http://localhost:8082"
	}
	url := fmt.Sprintf("%s/register/%s/%s", lookupServiceURL, name, number)
	resp, err := http.Post(url, "text/plain", strings.NewReader(""))
	if err != nil {
		fmt.Println("Guru Meditation: calling lookupService")
	}
	return resp.StatusCode
}

func pushToAccountService(name string, phoneNumber string, postCode string, email string) int {
	accountServiceURL := os.Getenv("ACCOUNT_SERVICE_URL")
	if accountServiceURL == "" {
		accountServiceURL = "http://localhost:8080"
	}
	url := fmt.Sprintf("%s/accountService/%s/%s/%s/%s", accountServiceURL, name, phoneNumber, postCode, email)
	resp, err := http.Post(url, "text/plain", strings.NewReader(""))
	if err != nil {
		fmt.Println("Guru Meditation: calling accountService")
	}
	return resp.StatusCode
}
