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
		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("-------read:", err)
				break
			}
			fmt.Println("==========", mt, message)
			//websocket.BinaryMessage
			err = conn.WriteMessage(mt, message)
			if err != nil {
				log.Println("+++++++++++write:", err)
				break
			}
		}
	})
	go http.ListenAndServe(":8199", nil)

	time.Sleep(time.Second * 5)
	conn.Close()
	fmt.Println("===guan bi==")
	time.Sleep(time.Second * 5)
	conn.Close()
	fmt.Println("===guan bi==")
	time.Sleep(time.Second * 10000)

}
