package searching

import "github.com/ayberktandogan/introduction-to-algorithms-go/common"

func Run() {
	A := common.RandomIntArr(common.LEN, common.MAX)

	o, el := LinearSearch(A, 3)
	common.PrintResult("Linear Search", o, *el)
}
