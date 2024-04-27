package sort

func recurseQuickSort(arr []int, lo, hi int) {
	if (hi - lo) < 1 {
		return
	}

	pivot := arr[hi]
	idx := lo - 1

	for i := lo; i < hi; i++ {
		curr := arr[i]
		if curr <= pivot {
			idx++
			arr[i] = arr[idx]
			arr[idx] = curr
		}
	}

	idx++
	arr[hi] = arr[idx]
	arr[idx] = pivot

	lo1 := lo
	hi1 := idx - 1
	recurseQuickSort(arr, lo1, hi1)

	lo2 := idx + 1
	hi2 := hi
	recurseQuickSort(arr, lo2, hi2)
}

func QuickSort(arr []int) {
	hi := len(arr) - 1
	lo := 0

	recurseQuickSort(arr, lo, hi)
}
