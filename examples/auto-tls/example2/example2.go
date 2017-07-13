package main

import (
	"log"

	"github.com/go-ego/autotls"
	"github.com/go-ego/ego"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	r := ego.Default()

	// Ping handler
	r.GET("/ping", func(c *ego.Context) {
		c.String(200, "pong")
	})

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}

	log.Fatal(autotls.RunWithManager(r, &m))
}
