package binaryx

import "github.com/ayberktandogan/introduction-to-algorithms-go/common"

const LEN = 5

func Run() {
	A := common.RandomBinaryArr(5)
	B := common.RandomBinaryArr(5)

	o, el := AddBinaryIntegers(A, B, LEN+1)
	common.PrintResult("Add Binary Integers", o, *el)
}
