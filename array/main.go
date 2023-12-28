package array

import "github.com/ayberktandogan/introduction-to-algorithms-go/common"

func Run() {
	A := common.RandomIntArr(common.LEN, common.MAX)

	o, el := SumArray(A)
	common.PrintResult("Sum Array", o, *el)
}
