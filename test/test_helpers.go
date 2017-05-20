package test

import (
	"net/http"

	"github.com/go-ego/ego"
)

func CreateTestContext(w http.ResponseWriter) (c *ego.Context, r *ego.Engine) {
	r = ego.New()
	c = r.allocateContext()
	c.reset()
	c.writermem.reset(w)
	return
}
