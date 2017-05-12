package mid

import (
	"github.com/go-ego/ego"
	"github.com/go-ego/ego/mid/logger"
)

func Classic() *ego.Engine {
	engine := ego.New()
	// engine.Use(Logger(), Recovery())
	engine.Use(logger.Logger(), logger.Recovery())
	return engine
}

func Default() *ego.Engine {
	engine := ego.New()
	// engine.Use(Logger(), Recovery())
	engine.Use(logger.Logger(), logger.Recovery())
	return engine
}
