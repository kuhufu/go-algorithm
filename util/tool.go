package util

import (
	"math/rand"
	"time"
)

func Shuffle(a []int) {
	N := len(a)
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i:= 0; i < N; i++ {
		i := r.Intn(N)
		j := r.Intn(N)
		swap(a, i, j)
	}
}

func RandIntArray(num int) []int {
	a := make([]int, num)
	//r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < num; i++ {
		a[i] = rand.Intn(num)
	}
	//Shuffle(a)
	return a
}

func RandIntArrayUseNewSeed(num int) []int {
	a := make([]int, num)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < num; i++ {
		a[i] = r.Intn(num)
	}
	//Shuffle(a)
	return a
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func IsSorted(a []int) bool {
	for i := 1; i < len(a)-1; i++ {
		if  compare(a[i], a[i-1]) == -1{
			return false
		}
	}

	return true
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