package sorting

import (
	"time"

	"github.com/ayberktandogan/introduction-to-algorithms-go/common"
)

func InsertionSort(a [common.LEN]int, n int) ([]int, *time.Duration) {
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

	return A[:], &el
}

// question 2.1-3
func ReverseInsertionSort(a [common.LEN]int, n int) ([]int, *time.Duration) {
	A := a

	start := time.Now()
	for i := 1; i < n; i++ {
		key := A[i]
		j := i - 1

		for j > -1 && A[j] < key {
			A[j+1] = A[j]
			j -= 1
		}

		A[j+1] = key
	}
	el := time.Since(start)

	return A[:], &el
}
