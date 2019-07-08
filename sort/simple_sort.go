package sort

func bubSort(a []int, begin, end int) {
	for ; end >= 0; end-- {
		for j := begin + 1; j < end; j++ {
			if less(a[j], a[j-1]) {
				swap(a, j, j-1)
			}
		}
	}
}

func insertSort(a []int, begin, end int) {
	for i := begin + 1; i < end; i++ {
		for j := i; j > begin; j-- {
			if less(a[j], a[j-1]) {
				swap(a, j, j-1)
			}
		}
	}
}

func selectSort2(a []int, begin, end int) {
	for i := begin; i < end; i++ {
		min := i
		amin := a[min]
		for j := i + 1; j < end; j++ {
			if less(a[j], amin) {
				min = j
				amin = a[j]
			}
		}
		swap(a, i, min)
	}
}

func selectSort3(a []int, begin, end int) {
	for end = end - 1; end >= begin; end-- {
		max := end
		amax := a[max]
		for j := begin; j < end; j++ {
			if less(amax, a[j]) {
				max = j
				amax = a[j]
			}
		}
		swap(a, end, max)
	}
}

func selectSort(a []int, begin, end int) {
	for i := begin; i < end; i++ {
		for j := i + 1; j < end; j++ {
			if less(a[j], a[i]) {
				swap(a, i, j)
			}
		}
	}
}
