package binaryx

import (
	"math"
	"time"
)

// question 2.1-5
func AddBinaryIntegers(A []int, B []int, n int) ([]int, *time.Duration) {
	sum := 0

	start := time.Now()
	for i := n - 2; i >= 0; i-- {
		sum += A[i] * int(math.Pow(2, float64(n-2-i)))
		sum += B[i] * int(math.Pow(2, float64(n-2-i)))
	}

	c := make([]int, n)
	for j := n - 1; j >= 0; j-- {
		bit := sum / int(math.Pow(2, float64(j)))
		sum = sum % int(math.Pow(2, float64(j)))
		c[j] = bit
	}
	el := time.Since(start)

	return c, &el
}
