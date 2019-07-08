package sort

func mergeSortHalfCopy(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := (lo + hi) / 2
	mergeSortHalfCopy(a, lo, mid)
	mergeSortHalfCopy(a, mid+1, hi)

	halfCopyMerge(a, lo, mid, hi)
}

func mergeSortFullCopy(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := (lo + hi) / 2
	mergeSortHalfCopy(a, lo, mid)
	mergeSortHalfCopy(a, mid+1, hi)

	fullCopyMerge(a, lo, mid, hi)
}

func halfCopyMerge(a []int, lo, mid, hi int) {
	//只复制前半段，不用全复制
	aux := make([]int, mid+1-lo)
	copy(aux, a[lo:mid+1])
	//a2 := a[mid+1 : hi+1]

	//         k       j
	//         l     m     h
	//   a  [][][][][][][][]
	//
	//         i
	//   aux  [][][][]

	i := 0
	j := mid + 1
	for k := lo; k <= hi; k++ {
		if i >= len(aux) {
			a[k] = a[j]
			j++
		} else if j >= hi+1 {
			a[k] = aux[i]
			i++
		} else if a[j] < aux[i] {
			a[k] = a[j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}
	}
}

func fullCopyMerge(a []int, lo, mid, hi int) {
	aux := make([]int, hi-lo+1)
	copy(aux, a[lo:hi+1])

	//         k
	//         l     m     h
	//   a  [][][][][][][][]
	//
	//         i     m j
	//   aux  [][][][][][][]

	i := 0
	auxMid := i + mid -lo
	j := auxMid + 1
	for k := lo; k <= hi; k++ {
		if i > auxMid {
			a[k] = aux[j]
			j++
		} else if j >= len(aux) {
			a[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			a[k] = aux[j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}
	}
}

func mergeSortHalfCopyWithP(a []int, lo, hi int) {
	if hi <= lo + 15 {
		bubSort(a, lo, hi+1)
		return
	}
	mid := (lo + hi) / 2
	mergeSortHalfCopyWithP(a, lo, mid)
	mergeSortHalfCopyWithP(a, mid+1, hi)

	halfCopyMerge(a, lo, mid, hi)
}

func mergeSortFinal(a []int, aux []int, lo int, hi int) {
	if hi <= lo + 15 {
		selectSort(a, lo, hi+1)
		return
	}
	mid := (lo + hi) / 2
	mergeSortFinal(a, aux, lo, mid)
	mergeSortFinal(a, aux, mid+1, hi)

	mergeFinal(a, aux, lo, mid, hi)
}

func mergeFinal(a []int, aux []int, lo int, mid int, hi int) {

	copy(aux[lo:hi+1], a[lo:hi+1]) //go提供的复制函数速度更快

	//          k       j
	//          l     m     h
	//   a   [][][][][][][][][][]
	//
	//          i     m     h
	//   aux [][][][][][][][][][]

	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			a[k] = aux[j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}
	}
}
