package main

import (
	"crypto/tls"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

func main() {
	client := ghttp.NewWebSocketClient()
	client.HandshakeTimeout = time.Second    // 设置超时时间
	client.Proxy = http.ProxyFromEnvironment // 设置代理
	client.TLSClientConfig = &tls.Config{}   // 设置 tls 配置
	conn, _, err := client.Dial("ws://127.0.0.1:9501", nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	err = conn.WriteMessage(websocket.TextMessage, []byte("hello word"))
	if err != nil {
		panic(err)
	}
	mt, data, err := conn.ReadMessage()
	if err != nil {
		panic(err)
	}
	fmt.Println(mt, data)
}
