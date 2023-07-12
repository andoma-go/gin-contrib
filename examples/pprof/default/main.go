package main

import (
	"github.com/andoma-go/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	pprof.Register(router)
	router.Run(":8080")
}
