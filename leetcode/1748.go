package main

import (
	"sort"
)

func sumOfUnique(nums []int) int {
	sum := 0

	if len(nums) == 0 {
		return sum
	}

	sort.Ints(nums)

	prev := 0
	cnt := 0

	for _, val := range nums {
		if cnt == 0 {
			prev = val
			cnt = 1
		} else if prev == val {
			cnt++
		} else {
			if cnt == 1 {
				sum += prev
			}

			prev = val
			cnt = 1
		}
	}

	if cnt == 1 {
		sum += prev
	}

	return sum
}
