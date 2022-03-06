package main

type RecentCounter struct {
	q []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		q: []int{},
	}
}

func (rc *RecentCounter) Ping(t int) int {

	rc.q = append(rc.q, t)

	lo, n := rc.getLeftMostValidIndex(t), len(rc.q)

	if n < 6000 {
		return n - lo + 1
	}

	for rc.q[0] < t-3000 {
		rc.q = rc.q[1:]
	}

	return len(rc.q)
}

func (rc *RecentCounter) getLeftMostValidIndex(t int) int {
	lo, hi := 0, len(rc.q)-1
	mid, res := -1, -1

	for lo <= hi {
		mid = lo + (hi-lo)/2

		if rc.q[mid] >= t-3000 {
			res = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return res
}
