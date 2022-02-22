package main

import (
	"fmt"
	"sort"
)

type sint32 []int32

func (a sint32) Len() int {
	return len(a)
}

func (a sint32) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a sint32) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a sint32) String() (s string) {
	for _, ele := range a {
		s += fmt.Sprintf("%d,", ele)
	}

	return
}

func main() {
	arr := []int32{5, 20, -10, 30, 0, 3, 2, -10}

	sarr := sint32(arr)

	sort.Sort(sarr)

	fmt.Printf("%s\n", sarr)
}
