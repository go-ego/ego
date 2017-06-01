package mid

import (
	"github.com/go-ego/ego"
	// "github.com/go-ego/ego/mid/logger"
)

// Classic returns an Engine instance with the Logger and Recovery middleware already attached.
func Classic() *ego.Engine {
	engine := ego.New()
	engine.Use(Logger(), Recovery())
	// engine.Use(logger.Logger(), logger.Recovery())
	return engine
}

// Default returns an Engine instance with the Logger and Recovery middleware already attached.
func Default() *ego.Engine {
	engine := ego.New()
	engine.Use(Logger(), Recovery())
	// engine.Use(logger.Logger(), logger.Recovery())
	return engine
}
