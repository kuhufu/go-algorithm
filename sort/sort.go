package sort

func BubSort(a []int) {
	bubSort(a, 0, len(a))
}

func InsertSort(a []int) {
	insertSort(a, 0, len(a))
}

func SelectSort(a []int) {
	selectSort(a, 0, len(a))
}

func SelectSort2(a []int) {
	selectSort2(a, 0, len(a))
}

func SelectSort3(a []int) {
	selectSort3(a, 0, len(a))
}

func QuickSort(a []int) {
	quickSort(a, 0, len(a)-1)
}

func Quick3Way(a []int) {
	quick3Way(a, 0, len(a)-1)
}

func MergeSort(a []int) {
	MergeSortFinal(a)
}

func MergeSortHalfCopy(a []int) {
	mergeSortHalfCopy(a, 0, len(a)-1)
}

func MergeSortFullCopy(a []int) {
	mergeSortFullCopy(a, 0, len(a)-1)
}

func MergeSortWithP(a []int) {
	mergeSortHalfCopyWithP(a, 0, len(a)-1)
}

func MergeSortFinal(a []int) {
	aux := make([]int, len(a))
	mergeSortFinal(a, aux, 0, len(a)-1)
}
