package limit

import (
	"log"
	"net"
	"strings"

	"github.com/andoma-go/gin"
)

// CIDR is a middleware that check given CIDR rules and return 403 Forbidden when user is not coming from allowed source.
// CIDRs accepts a list of CIDRs, separated by comma. (e.g. 127.0.0.1/32, ::1/128 )
func New(CIDRs string, opts ...Option) gin.HandlerFunc {

	limit := &Limit{
		handler: func(c *gin.Context, remoteAddr, CIDRs string) {
			log.Println("[LIMIT] Request from [" + remoteAddr + "] is not allow to access `" + c.Request.RequestURI + "`, only allow from: [" + CIDRs + "]")
			c.AbortWithStatus(403)
		},
	}

	for _, opt := range opts {
		opt(limit)
	}

	return func(c *gin.Context) {
		remoteAddr := c.ClientIP()

		// parse it into IP type
		remoteIP := net.ParseIP(remoteAddr)

		// split CIDRs by comma, and we are going check them one by one
		cidrSlices := strings.Split(CIDRs, ",")

		// under of CIDR we were in
		var matchCount uint

		// go over each CIDR and do the tests
		for _, cidr := range cidrSlices {
			// remove unwanted spaces
			cidr = strings.TrimSpace(cidr)

			// try to parse the CIDR
			_, cidrIPNet, err := net.ParseCIDR(cidr)
			if err != nil {
				_ = c.AbortWithError(500, err)
				return
			}

			// This is the core of this middleware, it ask current CIDR network range to test if current IP is in
			if cidrIPNet.Contains(remoteIP) {
				matchCount = matchCount + 1
			}
		}

		// if no CIDR ranges contains our IP
		if matchCount == 0 && limit.handler != nil {
			limit.handler(c, remoteAddr, CIDRs)
		}
	}
}
