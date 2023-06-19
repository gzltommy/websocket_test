package main

import (
	"crypto/tls"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
	"time"
)

func main() {
	client := ghttp.NewWebSocketClient()
	client.HandshakeTimeout = time.Second    // 设置超时时间
	client.Proxy = http.ProxyFromEnvironment // 设置代理
	client.TLSClientConfig = &tls.Config{}   // 设置 tls 配置
	conn, _, err := client.Dial("ws://127.0.0.1:8199/ws?user_id=100", nil)
	if err != nil {
		fmt.Println("---", err)
		return
	}
	defer conn.Close()

	fmt.Println("okokokokokokok")

	//err = conn.WriteMessage(websocket.BinaryMessage, []byte("hello word"))
	//if err != nil {
	//	panic(err)
	//}
	//mt, data, err := conn.ReadMessage()
	//if err != nil {
	//	panic(err)
	//}
	time.Sleep(time.Hour)
}
