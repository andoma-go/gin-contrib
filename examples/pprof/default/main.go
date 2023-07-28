package main

import (
	"github.com/andoma-go/gin"
	"github.com/andoma-go/gin-contrib/pprof"
)

func main() {
	router := gin.Default()
	pprof.Register(router)
	router.Run(":8080")
}
