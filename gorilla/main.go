package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const readBufferSize = 32 << 10 // 读消息缓冲区32K

func main() {
	var upGrader = websocket.Upgrader{
		ReadBufferSize:  readBufferSize,
		WriteBufferSize: readBufferSize,
		// 默认允许跨域访问
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	} // use default options

	var conn *websocket.Conn
	var err error
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err = upGrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer conn.Close()

		ticker := time.NewTimer(time.Minute * 3)
		for {
			select {
			case <-ticker.C:
				// 定时发心跳
				if err := conn.WriteControl(
					websocket.PingMessage,
					nil,
					time.Now().Add(time.Second)); err != nil {
					return
				}
			default:
				err = conn.SetReadDeadline(time.Now().Add(time.Minute * 3))
				msgType, msg, err := conn.ReadMessage()
				if err != nil {
					fmt.Printf("----ws.ReadMessage-fail--- %v,%v,%v \n", msgType, msg, err)
					return
				}

				fmt.Println("-------收到消息：------", msgType, string(msg))

				time.Sleep(time.Second)
				if err = conn.WriteMessage(msgType, msg); err != nil {
					fmt.Println("----ws.ReadMessage---fail----", err)
					return
				}
			}
		}
	})

	go http.ListenAndServe(":8199", nil)

	time.Sleep(time.Hour * 1)
}
