package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"time"
)

func main() {
	//注意这里要使用证书中包含的主机名称
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	conn, err := tls.Dial("tcp", "localhost:8888", tlsConfig)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()
	log.Println("Client Connect To ", conn.RemoteAddr())
	status := conn.ConnectionState()
	fmt.Printf("%#v\n", status)
	buf := make([]byte, 1024)
	ticker := time.NewTicker(1 * time.Millisecond * 500)
	for {
		select {
		case <-ticker.C:
			{
				_, err = io.WriteString(conn, "hello")
				if err != nil {
					log.Fatalln(err.Error())
				}
				len, err := conn.Read(buf)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("Receive From Server:", string(buf[:len]))
				}
			}
		}
	}

}
