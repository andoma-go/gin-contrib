package main

import (
	"github.com/gin-gonic/gin"

	"github.com/andoma-go/gin-contrib/nocache"
)

func main() {
	router := gin.Default()
	router.Use(nocache.NoCache())
	router.Run(":8080")
}
