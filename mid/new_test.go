package mid

import (
	"testing"

	"github.com/go-ego/ego"
	"github.com/stretchr/testify/assert"
)

func TestCreateDefaultRouter(t *testing.T) {
	router := ego.Default()
	assert.Len(t, router.Handlers, 2)
}

func TestCreateClassicRouter(t *testing.T) {
	router := ego.Classic()
	assert.Len(t, router.Handlers, 2)
}
