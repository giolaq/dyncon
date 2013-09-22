package quickunion

import "fmt"

func quickUnion(N int) []int {
	id := make([]int, N)
	for i := range id {
		id[i] = i
	}
	return id
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

func union(id []int, p, q int) {
	i := root(id, p)
	j := root(id, q)
	id[i] = j

}
