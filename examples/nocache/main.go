package main

import (
	"github.com/andoma-go/gin-contrib/nocache"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(nocache.NoCache())
	router.Run(":8080")
}
