package sort

import (
	"math/rand"
)

func quickSort(arr []int, lo, hi int) {
	if hi <= lo+15 {
		selectSort(arr, lo, hi+1)
		return
	}

	swap(arr, lo, rand.Intn(hi-lo)+lo)
	v := arr[lo]

	i := lo
	j := hi

	for {
		for i <= j && arr[i] <= v {
			i++
		}
		for i <= j && arr[j] > v {
			j--
		}
		if i > j {
			break
		}
		swap(arr, i, j)
	}

	swap(arr, lo, j)
	quickSort(arr, lo, j-1)
	quickSort(arr, j+1, hi)
}

func quick3Way(a []int, lo, hi int) {
	if hi <= lo+15 {
		bubSort(a, lo, hi+1)
		return
	}

	lt := lo
	i := lo + 1
	gt := hi

	swap(a, lo, rand.Intn(hi-lo)+lo)
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
