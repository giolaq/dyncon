package main
import "fmt"

func quickFindUF(N int) []int{
	id := make([]int, N)  
	for i := range id {
		    id[i] = i;
	    }
	return id
}

func connected(id []int, p int, q int) bool {
	return id[p] == id[q]
	
}

func union(id []int, p int, q int) {
	pid := id[p]
	qid := id[q]
	for i := 0; i< len(id); i++ {
		if  id[i] == pid  {
			id[i] = qid
		} 
	}
	
}
func main() {
	id := quickFindUF(10)
    union(id, 5,0)
	union(id, 5,6)
	union(id, 1,2)
	union(id, 8,3)
	union(id, 4,9)
	fmt.Printf ("connected 1,2 ? = %v \n", connected(id, 1,2) )
	fmt.Printf ("connected 1,7 ? = %v \n", connected(id, 1,7) )
}