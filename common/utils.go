package common

import (
	"fmt"
	"math/rand"
	"time"
)

const LEN = 100000
const MAX = 100000

func PrintResult(o [LEN]int, el time.Duration) {
	fmt.Printf("This process roughly took %s ", el.String())
}

func RandomIntArr(len int, max int) [LEN]int {
	var arr [LEN]int

	for idx := range arr {
		arr[idx] = rand.Intn(max)
	}

	return arr
}
