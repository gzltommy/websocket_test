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
	conn, _, err := client.Dial("ws://192.168.24.147:80", nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//mMsg := metaverse.MicroMessage{
	//	MessageId: 0,
	//	RequireId: 0,
	//	Guid:      0,
	//	Guids:     nil,
	//	SpaceId:   0,
	//	ConnId:    "",
	//	Key:       "",
	//	Body:      nil,
	//	BodyMsgId: 0,
	//}
	//
	//body := metaverse.CS_Login{Token: "aaaaa"}
	//
	//data, _ := proto.Marshal(msg)

	err = conn.WriteMessage(websocket.BinaryMessage, []byte("hello word"))
	if err != nil {
		panic(err)
	}
	mt, data, err := conn.ReadMessage()
	if err != nil {
		panic(err)
	}
	fmt.Println(mt, data)
}
