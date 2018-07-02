package mid

import (
	"testing"

	"github.com/go-ego/ego"
	"github.com/stretchr/testify/assert"
)

func TestCreateDefaultRouter(t *testing.T) {
	router := ego.Default()
	assert.Len(t, 2, router.Handlers)
}

func TestCreateClassicRouter(t *testing.T) {
	router := ego.Classic()
	assert.Len(t, 2, router.Handlers)
}
