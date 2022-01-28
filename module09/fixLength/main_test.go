package main

import (
	"log"
	"net"
	"testing"
)

func TestFixLength(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:5566")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for i := 0; i < 10; i++ {
		msg := []byte("巴拉巴拉巴拉巴拉巴拉巴拉巴拉巴拉巴拉巴拉巴拉巴拉巴拉巴拉巴拉巴拉巴拉巴拉巴拉巴拉巴拉巴拉")
		fixLength := dataLength - len(msg)
		msg = append(msg, make([]byte, fixLength)...)

		conn.Write(msg)
	}
}
