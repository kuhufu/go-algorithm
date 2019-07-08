package sort

import (
	"math/rand"
)

func quickSort(a []int, lo, hi int) {
	if hi <= lo+15 {
		selectSort(a, lo, hi+1)
		return
	}

	swap(a, lo, rand.Intn(hi-lo) + lo)
	v := a[lo]

	i := lo
	j := hi

	for {
		for i <= j && a[i] <= v {
			i++
		}
		for i <= j && a[j] >= v {
			j--
		}
		if i > j {
			break
		}
		swap(a, i, j)
	}

	swap(a, lo, i-1)
	quickSort(a, lo, i-2)
	quickSort(a, i, hi)
}

func quick3Way(a []int, lo, hi int) {
	if hi <= lo+15 {
		bubSort(a, lo, hi+1)
		return
	}

	lt := lo
	i := lo + 1
	gt := hi

	swap(a, lo, rand.Intn(hi-lo) + lo)
	v := a[lo]

	for i <= gt {
		if a[i] < v {
			swap(a, lt, i)
			lt++
			i++
		} else if a[i] > v {
			swap(a, i, gt)
			gt--
		} else {
			i++
		}
	}
	quick3Way(a, lo, lt-1)
	quick3Way(a, gt+1, hi)
}
