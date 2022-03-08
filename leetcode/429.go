package main

type Node struct {
	Val      int
	Children []*Node
}

func LevelOrder(root *Node) [][]int {

	type Pair struct {
		Val   *Node
		Level int
	}

	q := []*Pair{}

	res := [][]int{}

	if root == nil {
		return res
	} else {
		q = append(q, &Pair{Val: root, Level: 0})
	}

	prevLevel := -1

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr.Level == prevLevel {
			res[prevLevel] = append(res[prevLevel], curr.Val.Val)
		} else {
			res = append(res, []int{curr.Val.Val})
			prevLevel = curr.Level
		}

		for _, kid := range curr.Val.Children {
			if kid == nil {
				continue
			}
			q = append(q, &Pair{
				Val:   kid,
				Level: 1 + curr.Level,
			})
		}
	}

	return res
}
