package main

import (
	"log"
	"net/http"
	"time"

	"github.com/andoma-go/gin"
	"github.com/andoma-go/gin-contrib/timeout"
)

func emptySuccessResponse(c *gin.Context) {
	time.Sleep(200 * time.Microsecond)
	c.String(http.StatusOK, "")
}

func main() {
	r := gin.New()

	r.GET("/", timeout.New(
		timeout.WithTimeout(100*time.Microsecond),
		timeout.WithHandler(emptySuccessResponse),
	))

	// Listen and Server in 0.0.0.0:8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
