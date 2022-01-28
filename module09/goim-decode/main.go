package main

import (
	"encoding/binary"
	"fmt"
)

//https://github.com/Terry-Mao/goim/blob/557e33b9ca7da0c765675fdcb50a6133814b1d0a/api/protocol/protocol.go

func decoder(data []byte) {
	if len(data) < 16 {
		fmt.Println("Invalid package")
		return
	}

	packageLen := binary.BigEndian.Uint32(data[:4])
	fmt.Printf("packageLen:%v\n", packageLen)

	headerLen := binary.BigEndian.Uint16(data[4:6])
	fmt.Printf("headerLen:%v\n", headerLen)

	version := binary.BigEndian.Uint16(data[6:8])
	fmt.Printf("version:%v\n", version)

	operation := binary.BigEndian.Uint32(data[8:12])
	fmt.Printf("operation:%v\n", operation)

	sequence := binary.BigEndian.Uint32(data[12:16])
	fmt.Printf("sequence:%v\n", sequence)

	body := string(data[16:])
	fmt.Printf("body:%v\n", body)
}
