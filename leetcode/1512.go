package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	pairs := 0

	var freq map[int]int = make(map[int]int)

	for _, val := range nums {

		if cnt, has := freq[val]; has {
			pairs += cnt
			freq[val] += 1
		} else {
			freq[val] = 1
		}
	}

	return pairs
}

func main() {
	var nums = []int{1, 2, 3, 1, 1, 3}

	fmt.Println(numIdenticalPairs(nums))
}
