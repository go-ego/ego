package hello

import (
	"net/http"

	"github.com/go-ego/ego"
)

// This function's name is a must. App engine uses it to drive the requests properly.
func init() {
	// Starts a new ego instance with no middle-ware
	r := ego.New()

	// Define your handlers
	r.GET("/", func(c *ego.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.GET("/ping", func(c *ego.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Handle all requests using net/http
	http.Handle("/", r)
}
