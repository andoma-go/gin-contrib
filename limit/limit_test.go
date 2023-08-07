package limit

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andoma-go/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter(CIDRs string) *gin.Engine {
	// no debug mode
	gin.SetMode(gin.ReleaseMode)

	// create a default
	r := gin.Default()

	// our middle-ware
	r.Use(New(CIDRs))

	// routes
	r.GET("/", testGET)

	return r
}

func TestAllowAccessSource(t *testing.T) {
	r := setupRouter("127.0.0.1/32")

	// prepare
	ExpectedResponseStatus := 200

	// run
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = "127.0.0.1:80"
	r.ServeHTTP(w, req)

	// check
	assert.Equal(t, ExpectedResponseStatus, w.Code)
}

func TestNotAllowAccessSource(t *testing.T) {
	r := setupRouter("172.18.0.0/16")

	// prepare
	ExpectedResponseStatus := 403

	// run
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = "127.0.0.1:80"
	r.ServeHTTP(w, req)

	// check
	assert.Equal(t, ExpectedResponseStatus, w.Code)
}

func TestAllowAccessFromManySource(t *testing.T) {
	r := setupRouter("172.18.0.0/16, 127.0.0.1/32, ::1/128")

	// prepare
	ExpectedResponseStatus := 200

	// run
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = "127.0.0.1:80"
	r.ServeHTTP(w, req)

	// check
	assert.Equal(t, ExpectedResponseStatus, w.Code)
}

func TestNotAllowAccessFromManySource(t *testing.T) {
	r := setupRouter("172.18.0.0/16, 127.0.0.1/32, ::1/128")

	// prepare
	ExpectedResponseStatus := 403

	// run
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = "192.168.1.12:80"
	r.ServeHTTP(w, req)

	// check
	assert.Equal(t, ExpectedResponseStatus, w.Code)
}

func testGET(c *gin.Context) {
	c.String(200, "pong")
}
