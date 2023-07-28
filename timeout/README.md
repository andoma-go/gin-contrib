# Timeout

Timeout wraps a handler and aborts the process of the handler if the timeout is reached.

Copied from [gin-contrib/timeout](https://github.com/gin-contrib/timeout)

## Example

```go
package main

import (
  "log"
  "net/http"
  "time"

  "github.com/andoma-go/gin-contrib/timeout"
  "github.com/andoma-go/gin"
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

```

### custom error response

Add new error response func:

```go
func testResponse(c *gin.Context) {
  c.String(http.StatusRequestTimeout, "test response")
}
```

Add `WithResponse` option.

```go
  r.GET("/", timeout.New(
    timeout.WithTimeout(100*time.Microsecond),
    timeout.WithHandler(emptySuccessResponse),
    timeout.WithResponse(testResponse),
  ))
```

### custom middleware

```go
package main

import (
  "log"
  "net/http"
  "time"

  "github.com/andoma-go/gin-contrib/timeout"
  "github.com/andoma-go/gin"
)

func testResponse(c *gin.Context) {
  c.String(http.StatusRequestTimeout, "timeout")
}

func timeoutMiddleware() gin.HandlerFunc {
  return timeout.New(
    timeout.WithTimeout(500*time.Millisecond),
    timeout.WithHandler(func(c *gin.Context) {
      c.Next()
    }),
    timeout.WithResponse(testResponse),
  )
}

func main() {
  r := gin.New()
  r.Use(timeoutMiddleware())
  r.GET("/slow", func(c *gin.Context) {
    time.Sleep(800 * time.Millisecond)
    c.Status(http.StatusOK)
  })
  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```
