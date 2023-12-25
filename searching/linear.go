package searching

import (
	"time"

	"github.com/ayberktandogan/introduction-to-algorithms-go/common"
)

// question 2.1-4
func LinearSearch(A [common.LEN]int, i int) *int {
	var res *int = nil

	start := time.Now()
	for idx := 0; i < len(A); i++ {
		if A[idx] == i {
			res = &idx
			break
		}
	}
	el := time.Since(start)

	common.PrintResult(A, el)

	return res
}
