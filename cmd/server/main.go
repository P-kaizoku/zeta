package main

import (
	"net/http"

	"github.com/P-kaizoku/zeta/internal/transport/ws"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // fine for local dev, NOT for prod — flag for Phase 8
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	hub := ws.NewHub()
	go hub.Run()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "server is running"})
	})

	r.GET("/ws", func(ctx *gin.Context) {
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			return
		}

		client := &ws.Client{}

	})
}
