package limit

import "github.com/andoma-go/gin"

type Handler func(c *gin.Context, remoteAddr, CIDRs string) bool

type Limit struct {
	handler Handler
}

type Option func(*Limit)

// WithHandler
func WithHandler(handler Handler) Option {
	return func(limit *Limit) {
		limit.handler = handler
	}
}
