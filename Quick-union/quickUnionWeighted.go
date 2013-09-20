package main
import "fmt"

func quickUnionWeighted(N int)  ([]int, []int){
	id := make([]int, N) 
	sz := make([]int, N)  
	 
	for i := range id {
		    id[i] = i;
			sz[i] = 1;
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
	return root(id,p) == root(id, q)
}

func union(id, sz []int, p, q int) {
	i := root(id, p)
	j := root(id, q)
	if sz[i] < sz [j] {
		id[i] = j
		sz[j] = sz[j] + sz[i]
	} else {
		id[j] = i
		sz[i] = sz[i] + sz[j]
	}
	
}

func main() {
	id, sz := quickUnionWeighted(10)
    union(id,sz,5,0)
	union(id,sz,5,6)
	union(id,sz,1,2)
	union(id,sz,8,3)
	union(id,sz,4,9)
	fmt.Printf ("connected 1,2 ? = %v \n", connected(id,1,2) )
	fmt.Printf ("connected 1,7 ? = %v \n", connected(id,1,7) )
}