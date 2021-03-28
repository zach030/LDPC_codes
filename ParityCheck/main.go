package main

import "fmt"

//奇校验：原始码流+校验位 总共有奇数个1
//偶校验：原始码流+校验位 总共有偶数个1
const RawBitSize = 7
const TestBitSize = 8

func main() {
	// sample of ASCII
	// low seven bit is valid
	var correct = uint8(65)      // correct: 100 0001 ; case1: 100 0011 ; case2: 100 0010; case3: 100 1010
	testCases := []uint8{67, 66,74}
	oddAns := oddCheck(correct, testCases)
	evenAns := evenCheck(correct, testCases)
	fmt.Printf("odd-check answer is:%t, even-check answer is:%t", oddAns, evenAns)
}

func oddCheck(correct uint8, cases []uint8) []bool {
	ret := make([]bool, 0)
	// 获得正确中的1的个数
	correctOneBitNum := getOneBitNum(correct,RawBitSize)
	// 经过奇数校验后得到的正确校验码流
	correctCase := getOddBitStream(correctOneBitNum, correct)
	for _, c := range cases {
		// 获得测试校验码流，并对比
		ret = append(ret, getOneBitNum(getTestBitStream(correctCase, c),TestBitSize)%2 == 0)
	}
	return ret
}

func evenCheck(correct uint8, cases []uint8) []bool {
	ret := make([]bool, 0)
	// 获得正确中的1的个数
	correctOneBitNum := getOneBitNum(correct,RawBitSize)
	// 经过奇数校验后得到的正确校验码流
	correctCase := getEvenBitStream(correctOneBitNum, correct)
	for _, c := range cases {
		// 获得测试校验码流，并对比
		ret = append(ret, getOneBitNum(getTestBitStream(correctCase, c),TestBitSize)%2 != 0)
	}
	return ret
}

// 获取待校验数中1的个数
func getOneBitNum(num uint8, size int) int {
	hasNum := 0
	andNum := uint8(1)
	for i := 0; i < size; i++ {
		if num&andNum != 0 {
			hasNum++
		}
		andNum <<= 1
	}
	return hasNum
}

// 获取奇校验正确码流
func getOddBitStream(rawBitNum int, num uint8) uint8 {
	ret := uint8(0)
	if rawBitNum%2 == 0 {
		ret = 0X80 | num
	} else {
		ret = num
	}
	return ret
}

// 获取偶校验正确码流
func getEvenBitStream(rawBitNum int, num uint8) uint8 {
	ret := uint8(0)
	if rawBitNum%2 == 0 {
		ret = num
	} else {
		ret = 0X80 | num
	}
	return ret
}

// 获取测试的码流
func getTestBitStream(correct, c uint8) uint8 {
	return correct & 0X80 | c
}
