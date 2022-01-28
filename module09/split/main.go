package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:5577")
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

// 发送方，接受放事先分隔符，发送方在数据末尾添加指定分隔符
const dataLength = 1024

func handler(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadSlice('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("内容为:", string(msg))

	}
}
