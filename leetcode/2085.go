package main

func countWords(words1 []string, words2 []string) int {
	cnt := 0

	mp := make(map[string]int)

	for _, val := range words1 {
		mp[val]++
	}

	for _, val := range words2 {
		if ele, has := mp[val]; has {
			if ele < 2 {
				mp[val]--
			}
		}
	}

	for _, val := range mp {
		if val == 0 {
			cnt++
		}
	}

	return cnt
}
