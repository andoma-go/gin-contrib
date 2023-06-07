package main

import (
	"fmt"

	"github.com/andoma-go/gin-contrib/traceid"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(traceid.TraceId())
	router.GET("/", func(c *gin.Context) {
		fmt.Println(traceid.FromTraceId(c.Request.Context()))
		fmt.Println(traceid.GetTraceId(c))
	})
	router.Run(":8080")
}
