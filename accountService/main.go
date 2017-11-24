package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type account struct {
	name        string
	phoneNumber string
	postCode    string
	email       string
}

func main() {
	accountRegistry := make(map[string]account)

	router := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	port = fmt.Sprintf(":%s", port)

	// This handler will match /user/john but will not match neither /user/ or /user
	router.POST("/accountService/:name/:phoneNumber/:postCode/:email", register(accountRegistry))
	router.GET("/accountService/:name", lookup(accountRegistry))
	router.Run(port)
}

func register(accountRegistry map[string]account) func(*gin.Context) {
	return func(c *gin.Context) {
		newAccount := account{
			name:        c.Param("name"),
			phoneNumber: c.Param("phoneNumber"),
			postCode:    c.Param("postCode"),
			email:       c.Param("email"),
		}

		accountRegistry[newAccount.name] = newAccount

		c.Status(http.StatusOK)
	}
}

func lookup(accountRegistry map[string]account) func(*gin.Context) {
	return func(c *gin.Context) {
		acc, ok := accountRegistry[c.Param("name")]

		if ok {
			c.JSON(http.StatusOK, acc)
		} else {
			c.Status(http.StatusNotFound)
		}
	}
}
