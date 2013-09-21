package main

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

func main() {
	id, sz := weightedQuickUnionByHeight(10)
	union(id, sz, 5, 0)
	union(id, sz, 5, 6)
	union(id, sz, 1, 2)
	union(id, sz, 8, 3)
	union(id, sz, 4, 9)
	fmt.Printf("connected 1,2 ? = %v \n", connected(id, 1, 2))
	fmt.Printf("connected 1,7 ? = %v \n", connected(id, 1, 7))
}
