# Ego Default Server

This is API experiment for Ego.

```go
package main

import (
	"github.com/go-ego/ego"
	"github.com/go-ego/ego/egoS"
)

func main() {
	egoS.GET("/", func(c *gin.Context) { c.String(200, "Hello World") })
	egoS.Run()
}
```
