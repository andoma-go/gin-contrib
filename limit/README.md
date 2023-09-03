# Gin Access Limit Middleware

Copied from [bu/gin-access-limit](https://github.com/bu/gin-access-limit)

## Usage

```go

package main

import (
    gin "github.com/andoma-go/gin"
    "github.com/andoma-go/gin-contrib/limit"
)

func main() {
    // create a Gin engine
    r := gin.Default()

    // this API is only accessible from Docker containers
    r.Use(limit.New("172.18.0.0/16"))

    // if need to specify serveral range of allowed sources, use comma to concatenate them
    // r.Use(limit.New("172.18.0.0/16, 127.0.0.1/32"))

    // routes
    r.GET("/", func (c *gin.Context) {
        c.String(200, "pong")
    })

    // listen to request
    r.Run(":8080")
}

```
