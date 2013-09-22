package quickunionweighted

import "fmt"

func quickUnionWeighted(N int) ([]int, []int) {
	id := make([]int, N)
	sz := make([]int, N)

	for i := range id {
		id[i] = i
		sz[i] = 1
	}
	return id, sz
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

func union(id, sz []int, p, q int) {
	i := root(id, p)
	j := root(id, q)
	if sz[i] < sz[j] {
		id[i] = j
		sz[j] = sz[j] + sz[i]
	} else {
		id[j] = i
		sz[i] = sz[i] + sz[j]
	}

}

