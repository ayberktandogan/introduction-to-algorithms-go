package sorting

import (
	"time"

	"github.com/ayberktandogan/introduction-to-algorithms-go/common"
)

func Run() {
	A := common.RandomIntArr(common.LEN, common.MAX)

	o, el := InsertionSort(A, len(A))
	common.PrintResult("Insertion Sort", o, *el)

	o, el = ReverseInsertionSort(A, len(A))
	common.PrintResult("Reverse Insertion Sort", o, *el)

	start := time.Now()
	cpA := MergeSort(A[:])
	elp := time.Since(start)
	common.PrintResult("Merge Sort", cpA, elp)
}
