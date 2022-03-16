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
	client.HandshakeTimeout = time.Second                                        // 设置超时时间
	client.Proxy = http.ProxyFromEnvironment                                     // 设置代理
	client.TLSClientConfig = &tls.Config{RootCAs: nil, InsecureSkipVerify: true} // 设置 tls 配置
	conn, r, err := client.Dial("wss://api.secondlive.world/ws/", nil)
	if err != nil {
		fmt.Printf("===== Dial === err:%v, r:%v \n", err, r)
		return
	}
	defer conn.Close()
	err = conn.WriteMessage(websocket.TextMessage, []byte("hello word"))
	if err != nil {
		fmt.Println("=====1111===", err)
		return
	}
	mt, data, err := conn.ReadMessage()
	if err != nil {
		fmt.Println("=====2222===", err)
		return
	}
	fmt.Println("==222===", mt, data)
}
