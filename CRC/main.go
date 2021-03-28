package main

import "fmt"

const (
	n          = 6
	k          = 4
	POLYNOMIAL = 0X9800 // 生成多项式 (1001 1000 0000 0000)
	DATAFRAME  = 0X002B // 信息位    (0000 0000 0010 1011)
)

func main() {
	ch := make(chan uint16, 1)
	// 信息位左移4位: 101011 0000
	dataFrame := fillZero(DATAFRAME, k)
	// 生成多项式除以信息位，得到FCS序列 0100
	fcs := crcDivide(POLYNOMIAL, dataFrame)
	// 生成消息帧: 101011 0100 发送到管道
	ch <- getDataFrame(dataFrame, fcs)
	// 从管道接收并解析消息帧
	res := decodeDataFrame(POLYNOMIAL, <-ch)
	fmt.Printf("crc result is: %d", res)
}

func fillZero(data uint16, size int) uint16 {
	return data << size
}

func getDataFrame(data, fcs uint16) uint16 {
	return data | (fcs & 0X0F)
}

func crcDivide(check, data uint16) uint16 {
	i := uint16(0)
	remainder := data
	for i = 0; i < 16; i++ {
		// 如果最高位为1，则进行模2除
		if remainder&0X8000 == 1 {
			remainder ^= check
		}
		// 左移一位
		remainder = remainder << 1
	}
	// 返回余数的低k位
	return remainder & 0X000F
}

func decodeDataFrame(check, data uint16) uint16 {
	return crcDivide(check, data)
}
