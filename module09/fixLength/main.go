package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:5566")
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handler(conn)
	}
}

// 发送方，接受放事先约定数据长度,发送的数据只能在长度范围内，不足长度则进行补齐
const dataLength = 1024

func handler(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		m := make([]byte, dataLength)
		n, err := reader.Read(m)
		if err == io.EOF {
			fmt.Println(err)
			return
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("内容为:", string(m[:n]))

	}
}
