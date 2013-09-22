package weightedquickunionbyheight

import "fmt"

func weightedQuickUnionByHeight(N int) ([]int, []int) {
	id := make([]int, N)
	ht := make([]int, N)

	for i := range id {
		id[i] = i
		ht[i] = 0
	}
	return id, ht
}

func root(id []int, i int) int {
	for i != id[i] {
		i = id[i]
	}
	return i
}

func connected(id []int, p, q int) bool {
	return root(id, p) == root(id, q)
}

func union(id, ht []int, p, q int) {
	i := root(id, p)
	j := root(id, q)
	if ht[i] < ht[j] {
		id[i] = j
	} else if ht[i] > ht[j] {
		id[j] = i
	} else {
		id[j] = i
		ht[i]++
	}

}
