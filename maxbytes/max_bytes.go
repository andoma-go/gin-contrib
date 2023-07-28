package maxbytes

import (
	"net/http"

	"github.com/andoma-go/gin"
)

// MaxBytes returns a middleware that limit reading of http request body.
func MaxBytes(n int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		if n > 0 && c.Request.ContentLength > n {
			c.AbortWithStatus(http.StatusRequestEntityTooLarge)
		}
	}
}
