package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gorilla/websocket"
	"time"
)

func main() {
	s := g.Server()
	var (
		ws  *ghttp.WebSocket
		err error
	)
	s.BindHandler("/ws", func(r *ghttp.Request) {
		ws, err = r.WebSocket()
		if err != nil {
			r.Exit()
		}

		fmt.Println("=========连接来了========")

		ticker := time.NewTimer(time.Second * 3)
		for {
			select {
			case <-ticker.C:
				if err := ws.WriteControl(
					websocket.PingMessage,
					nil,
					time.Now().Add(time.Second)); err != nil {
					return
				}
			default:
				ws.SetReadDeadline(time.Now().Add(time.Minute * 3))
				msgType, msg, err := ws.ReadMessage()
				if err != nil {
					fmt.Printf("----ws.ReadMessage---- %v,%v,%v \n", msgType, msg, err)
					return
				}

				fmt.Println("-------xh------", string(msg))

				time.Sleep(time.Second)
				if err = ws.WriteMessage(msgType, msg); err != nil {
					fmt.Println("----ws.ReadMessage-------", err)
				}
			}
		}
	})
	s.SetServerRoot(gfile.MainPkgPath())
	s.SetPort(8199)
	s.Run()
}
