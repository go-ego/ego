package mid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDefaultRouter(t *testing.T) {
	router := Default()
	assert.Len(t, router.Handlers, 2)
}

func TestCreateClassicRouter(t *testing.T) {
	router := Classic()
	assert.Len(t, router.Handlers, 2)
}
