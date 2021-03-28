package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, "hello")
	sum := checkSum(buffer.Bytes())
	fmt.Printf("check sum is:%d", sum)
}

func checkSum(data []byte) uint16 {
	var (
		sum    uint32
		length = len(data)
		index  int
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index])
	}
	sum += sum >> 16
	return uint16(^sum)
}
