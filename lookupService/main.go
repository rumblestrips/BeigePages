package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var phoneRegistry map[string]string
var numLookups int
var numRegistrations int

func main() {
	phoneRegistry = make(map[string]string)
	numLookups = 0
	numRegistrations = 0

	router := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	port = fmt.Sprintf(":%s", port)

	// This handler will match /user/john but will not match neither /user/ or /user
	router.GET("/lookup/:name", lookup)
	router.POST("/register/:name/:phoneNumber", register)
	router.GET("/metrics", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("num_lookups %d\nnum_registrations %d", numLookups, numRegistrations))
	})
	router.Run(port)
}

func lookup(c *gin.Context) {
	numLookups++
	phoneNumber, ok := phoneRegistry[c.Param("name")]

	if ok {
		c.JSON(http.StatusOK, gin.H{
			"phoneNumber": phoneNumber,
		})
	} else {
		c.Status(http.StatusNotFound)
	}
}

func register(c *gin.Context) {
	numRegistrations++
	phoneRegistry[c.Param("name")] = c.Param("phoneNumber")

	c.Status(http.StatusOK)
}
