package controllers

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // configure upgrader

func AllChatsController(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Error(err)
		return
	}
	// ensure connection close when function returns
	defer ws.Close()
	fmt.Println(ws.LocalAddr())
	fmt.Println(ws.RemoteAddr())
}
