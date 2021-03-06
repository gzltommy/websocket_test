package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
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

		for {
			msgType, msg, err := ws.ReadMessage()
			if err != nil {
				fmt.Println("----ws.ReadMessage----", err)
				return
			}

			fmt.Println("-------xh------")

			time.Sleep(time.Second)
			if err = ws.WriteMessage(msgType, msg); err != nil {
				fmt.Println("----ws.ReadMessage-------", err)
			}
		}
	})
	s.SetServerRoot(gfile.MainPkgPath())
	s.SetPort(8199)
	s.Run()
}
