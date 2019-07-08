package sort

func less(i, j int) bool {
	return compare(i, j) == -1
}

func compare(i, j int) int {
	if i > j {
		return 1
	}
	if i == j {
		return 0
	} else {
		return -1
	}
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}
