package sorting

func MergeSort(A []int) []int {
	if len(A) < 2 {
		return A
	}

	f := MergeSort(A[:(len(A) / 2)])
	s := MergeSort(A[len(A)/2:])

	return merge(f, s)
}

func merge(f []int, s []int) []int {
	i := 0
	j := 0

	fin := make([]int, 0)

	for i < len(f) && j < len(s) {
		if f[i] < s[j] {
			fin = append(fin, f[i])
			i++
		} else {
			fin = append(fin, s[j])
			j++
		}
	}

	for i < len(f) {
		fin = append(fin, f[i])
		i++
	}

	for j < len(s) {
		fin = append(fin, s[j])
		j++
	}

	return fin
}
