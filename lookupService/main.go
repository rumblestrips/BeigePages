package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gin-gonic/gin"
)

var phoneRegistry map[string]string
var numLookups prometheus.Counter
var numRegistrations prometheus.Counter

func main() {
	phoneRegistry = make(map[string]string)
	numLookups = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "beigepages",
		Subsystem: "lookup",
		Name:      "num_lookups",
		Help:      "Number of lookups",
	})
	numRegistrations = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "beigepages",
		Subsystem: "lookup",
		Name:      "num_registrations",
		Help:      "Number of registrations",
	})
	prometheus.MustRegister(numLookups)
	prometheus.MustRegister(numRegistrations)

	router := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	port = fmt.Sprintf(":%s", port)

	// This handler will match /user/john but will not match neither /user/ or /user
	router.GET("/lookup/:name", lookup)
	router.POST("/register/:name/:phoneNumber", register)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.Run(port)
}

func lookup(c *gin.Context) {
	numLookups.Inc()
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
	numRegistrations.Inc()
	phoneRegistry[c.Param("name")] = c.Param("phoneNumber")

	c.Status(http.StatusOK)
}
