package main

func countPoints(rings string) int {
	cnt, n := 0, len(rings)

	mp := make(map[int]map[byte]interface{})

	for i := 0; i < n; i += 2 {
		color, rod := rings[i], int(rings[i+1]-'0')

		if mp[rod] == nil {
			mp[rod] = make(map[byte]interface{})
		}

		mp[rod][color] = struct{}{}
	}

	for _, val := range mp {
		if len(val) == 3 {
			cnt++
		}
	}

	return cnt
}
