package sort

func recurseQuickSort(arr []int, lo, hi, p int) {
	if (hi - lo) < 1 {
		return
	}

	pValue := arr[p]
	idx := lo - 1

	for i := lo; i <= hi; i++ {
		if i == p {
			continue
		}

		curr := arr[i]
		if curr <= pValue {
			idx++
			arr[i] = arr[idx]
			arr[idx] = curr

			if idx == p {
				p = i
			}
		}
	}

	idx++
	arr[p] = arr[idx]
	arr[idx] = pValue

	lo1 := lo
	hi1 := idx - 1
	p1 := lo1 + (hi1-lo1)/2
	recurseQuickSort(arr, lo1, hi1, p1)

	lo2 := idx + 1
	hi2 := hi
	p2 := lo2 + (hi2-lo2)/2
	recurseQuickSort(arr, lo2, hi2, p2)
}

func QuickSort(arr []int) {
	hi := len(arr) - 1
	lo := 0
	p := hi / 2

	recurseQuickSort(arr, lo, hi, p)
}
