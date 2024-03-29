package main

import (
	"net/http"

	"github.com/andoma-go/gin"
	"github.com/andoma-go/gin-contrib/pprof"
)

func main() {
	router := gin.Default()
	adminGroup := router.Group("/admin", func(c *gin.Context) {
		if c.Request.Header.Get("Authorization") != "foobar" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	})
	pprof.RouteRegister(adminGroup, "pprof")
	router.Run(":8080")
}
