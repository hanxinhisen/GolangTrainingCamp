package main

import (
	"log"
	"net"
	"testing"
)

func TestSplit(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:5577")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for i := 0; i < 10; i++ {
		msg := []byte("阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴阿巴\n")
		conn.Write(msg)
	}
}
