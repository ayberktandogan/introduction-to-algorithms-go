package sorting

import (
	"github.com/ayberktandogan/introduction-to-algorithms-go/common"
)

func Run() {
	A := common.RandomIntArr(common.LEN, common.MAX)

	InsertionSort(A, len(A))
	ReverseInsertionSort(A, len(A))
}
