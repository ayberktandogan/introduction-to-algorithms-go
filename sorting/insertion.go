package sorting

import (
	"time"
)

func InsertionSort(a [LEN]int, n int) {
	A := a

	start := time.Now()
	for i := 1; i < n; i++ {
		key := A[i]
		j := i - 1

		for j > -1 && A[j] > key {
			A[j+1] = A[j]
			j -= 1
		}

		A[j+1] = key
	}
	el := time.Since(start)

	printResult(A, el)
}
