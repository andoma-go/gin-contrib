# Gin Static Middleware

Static middleware, support both local files and embed filesystem.

Copied from [gin-contrib/static](https://github.com/gin-contrib/static) and [soulteary/gin-static](https://github.com/soulteary/gin-static)

## Quick Start

### Download and Import

Download and install it:

```bash
go get github.com/andoma-go/gin-contrib/static
```

Import it in your code:

```go
import "github.com/andoma-go/gin-contrib/static"
```

## Example

See the [example](../examples/static/)

### Serve Local Files

```go
package main

import (
	"log"

	"github.com/andoma-go/gin-contrib/static"
	"github.com/andoma-go/gin"
)

func main() {
	r := gin.Default()

	// if Allow DirectoryIndex
	// r.Use(static.Serve("/", static.LocalFile("./public", true)))
	// set prefix
	// r.Use(static.Serve("/static", static.LocalFile("./public", true)))

	r.Use(static.Serve("/", static.LocalFile("./public", false)))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "test")
	})

	// Listen and Server in 0.0.0.0:8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
```

### Serve Embed folder

```go
package main

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/andoma-go/gin-contrib/static"
	"github.com/andoma-go/gin"
)

//go:embed public
var EmbedFS embed.FS

func main() {
	r := gin.Default()

	// method 1: use as Gin Router
	// trim embedfs path `public/page`, and use it as url path `/`
	// r.GET("/", static.ServeEmbed("public/page", EmbedFS))

	// method 2: use as middleware
	// trim embedfs path `public/page`, the embedfs path start with `/`
	// r.Use(static.ServeEmbed("public/page", EmbedFS))

	// method 2.1: use as middleware
	// trim embedfs path `public/page`, the embedfs path start with `/public/page`
	// r.Use(static.ServeEmbed("", EmbedFS))

	// method 3: use as manual
	// trim embedfs path `public/page`, the embedfs path start with `/public/page`
	// staticFiles, err := static.EmbedFolder(EmbedFS, "public/page")
	// if err != nil {
	// 	log.Fatalln("initialization of embed folder failed:", err)
	// } else {
	// 	r.Use(static.Serve("/", staticFiles))
	// }

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "test")
	})

	r.NoRoute(func(c *gin.Context) {
		fmt.Printf("%s doesn't exists, redirect on /\n", c.Request.URL.Path)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
```
