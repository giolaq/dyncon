package main

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

func main() {
	id := quickUnion(10)
	union(id, 5, 0)
	union(id, 5, 6)
	union(id, 1, 2)
	union(id, 8, 3)
	union(id, 4, 9)
	fmt.Printf("connected 1,2 ? = %v \n", connected(id, 1, 2))
	fmt.Printf("connected 1,7 ? = %v \n", connected(id, 1, 7))
}
