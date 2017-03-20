package main

import (
	"fmt"
	"runtime"

	"github.com/go-ego/ego"
)

func main() {
	ConfigRuntime()
	StartWorkers()
	StartEgo()
}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func StartWorkers() {
	go statsWorker()
}

func StartEgo() {
	ego.SetMode(ego.ReleaseMode)

	router := ego.New()
	router.Use(rateLimit, ego.Recovery())
	router.LoadHTMLGlob("resources/*.templ.html")
	router.Static("/static", "resources/static")
	router.GET("/", index)
	router.GET("/room/:roomid", roomGET)
	router.POST("/room-post/:roomid", roomPOST)
	router.GET("/stream/:roomid", streamRoom)

	router.Run(":80")
}
