package sorting

import (
	"fmt"
	"math/rand"
	"time"
)

const LEN = 100000
const MAX = 100000

func Run() {
	A := randomIntArr(LEN, MAX)

	InsertionSort(A, len(A))
}

func randomIntArr(len int, max int) [LEN]int {
	var arr [LEN]int

	for idx := range arr {
		arr[idx] = rand.Intn(max)
	}

	return arr
}

func printResult(o [LEN]int, el time.Duration) {
	for i := 0; i < len(o)-1; i++ {
		if o[i] > o[i+1] {
			panic("the list is not sorted!!!")
		}
	}

	fmt.Printf("This process roughly took %s ", el.String())
}
