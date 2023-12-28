package common

import (
	"fmt"
	"math/rand"
	"time"
)

const LEN = 100000
const MAX = 100000

func PrintResult(prName string, o any, el time.Duration) {
	fmt.Printf("This process [%s] roughly took %s \n", prName, el.String())
}

func RandomIntArr(len int, max int) [LEN]int {
	var arr [LEN]int

	for idx := range arr {
		arr[idx] = rand.Intn(max)
	}

	return arr
}

func RandomBinaryArr(len int) []int {
	arr := make([]int, len)

	for idx := range arr {
		arr[idx] = rand.Intn(2)
	}

	return arr
}
