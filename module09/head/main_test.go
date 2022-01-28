package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"testing"
)

func TestHead(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:5588")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for i := 0; i < 10; i++ {
		msg := "阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴"
		data, err := encode(msg)
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.Write(data)
	}

}
func encode(message string) ([]byte, error) {
	// 读取信息的长度，转换成4个字节
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)

	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}

	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}
