package main

func maxNumberOfBalloons(text string) int {
	mp := map[rune]int{
		'b': 1,
		'a': 1,
		'l': 2,
		'o': 2,
		'n': 1,
	}

	freq := map[rune]int{}

	for _, ele := range text {
		freq[ele]++
	}

	cnt := len(text)

	for ch, val := range mp {
		if freq[ch]/val < cnt {
			cnt = freq[ch] / val
		}
	}

	return cnt

}
