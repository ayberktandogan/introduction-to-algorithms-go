package searching

import "github.com/ayberktandogan/introduction-to-algorithms-go/common"

func Run() {
	A := common.RandomIntArr(common.LEN, common.MAX)

	LinearSearch(A, 3)
}
