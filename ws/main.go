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

		for {
			//msgType, msg, err := ws.ReadMessage()
			//if err != nil {
			//	fmt.Println("----ws.ReadMessage----", err)
			//	//return
			//}

			time.Sleep(time.Second)
			if err = ws.WriteMessage(1, []byte("666666")); err != nil {
				fmt.Println("----ws.ReadMessage-------", err)
			}
		}
	})
	s.SetServerRoot(gfile.MainPkgPath())
	s.SetPort(8199)
	go s.Run()

	time.Sleep(time.Second * 10)
	fmt.Println("======= close ======")
	ws.Close()

	fmt.Println("======= close ===222===")
	ws.Close()
	time.Sleep(time.Second * 1000)
}
