package searching

import (
	"time"

	"github.com/ayberktandogan/introduction-to-algorithms-go/common"
)

// question 2.1-4
func LinearSearch(A [common.LEN]int, i int) (*int, *time.Duration) {
	var res *int = nil

	start := time.Now()
	for idx := 0; idx < len(A); idx++ {
		if A[idx] == i {
			res = &i
			break
		}
	}
	el := time.Since(start)

	return res, &el
}
