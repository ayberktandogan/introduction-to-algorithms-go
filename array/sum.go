package array

import (
	"time"

	"github.com/ayberktandogan/introduction-to-algorithms-go/common"
)

// question 2.1-2
func SumArray(A [common.LEN]int) (int, *time.Duration) {
	sum := 0

	start := time.Now()
	for i := 0; i < len(A); i++ {
		sum += A[i]
	}
	el := time.Since(start)

	return sum, &el
}
