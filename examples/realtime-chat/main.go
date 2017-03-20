package main

import (
	"fmt"
	"io"
	"math/rand"

	"github.com/go-ego/ego"
)

func main() {
	router := ego.Default()
	router.SetHTMLTemplate(html)

	router.GET("/room/:roomid", roomGET)
	router.POST("/room/:roomid", roomPOST)
	router.DELETE("/room/:roomid", roomDELETE)
	router.GET("/stream/:roomid", stream)

	router.Run(":3000")
}

func stream(c *ego.Context) {
	roomid := c.Param("roomid")
	listener := openListener(roomid)
	defer closeListener(roomid, listener)

	c.Stream(func(w io.Writer) bool {
		c.SSEvent("message", <-listener)
		return true
	})
}

func roomGET(c *ego.Context) {
	roomid := c.Param("roomid")
	userid := fmt.Sprint(rand.Int31())
	c.HTML(200, "chat_room", ego.H{
		"roomid": roomid,
		"userid": userid,
	})
}

func roomPOST(c *ego.Context) {
	roomid := c.Param("roomid")
	userid := c.PostForm("user")
	message := c.PostForm("message")
	room(roomid).Submit(userid + ": " + message)

	c.JSON(200, ego.H{
		"status":  "success",
		"message": message,
	})
}

func roomDELETE(c *ego.Context) {
	roomid := c.Param("roomid")
	deleteBroadcast(roomid)
}
