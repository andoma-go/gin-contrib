# pprof

gin pprof middleware

> Package pprof serves via its HTTP server runtime profiling data in the format expected by the pprof visualization tool.

Copied from [gin-contrib/pprof](https://github.com/gin-contrib/pprof)

## Usage

### Start using it

Import it in your code:

```go
import "github.com/andoma-go/gin-contrib/pprof"
```

### Example

```go
package main

import (
  "github.com/andoma-go/gin-contrib/pprof"
  "github.com/andoma-go/gin"
)

func main() {
  router := gin.Default()
  pprof.Register(router)
  router.Run(":8080")
}
```

### change default path prefix

```go
func main() {
  router := gin.Default()
  // default is "debug/pprof"
  pprof.Register(router, "dev/pprof")
  router.Run(":8080")
}
```

### custom router group

```go
package main

import (
  "net/http"

  "github.com/andoma-go/gin-contrib/pprof"
  "github.com/andoma-go/gin"
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

```

### Use the pprof tool

Then use the pprof tool to look at the heap profile:

```bash
go tool pprof http://localhost:8080/debug/pprof/heap
```

Or to look at a 30-second CPU profile:

```bash
go tool pprof http://localhost:8080/debug/pprof/profile
```

Or to look at the goroutine blocking profile, after calling runtime.SetBlockProfileRate in your program:

```bash
go tool pprof http://localhost:8080/debug/pprof/block
```

Or to collect a 5-second execution trace:

```bash
wget http://localhost:8080/debug/pprof/trace?seconds=5
```
