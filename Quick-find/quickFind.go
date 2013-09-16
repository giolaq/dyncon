package main
import "fmt"

type IntSlice []int

func (is *IntSlice) QuickFindUF(N int) {
    slice := *is
	slice = make([]int, N)  
	for i := range slice {
		    slice[i] = i;
	    }
		
	*is = slice
  
}


func (is *IntSlice) connected(p int, q int) bool {
	return (*is)[p] == (*is)[q]
	
}

func (is *IntSlice) union(p int, q int) {
	pid := (*is)[p]
	qid := (*is)[q]
	for i := 0; i< len(*is); i++ {
		if  (*is)[i] == pid  {
			(*is)[i] = qid
		} 
	}
	
}
func main() {
	var intS IntSlice
	intS.QuickFindUF(10)

    intS.union(5,0)
	intS.union(5,6)
	intS.union(1,2)
	intS.union(8,3)
	intS.union(4,9)
	fmt.Printf ("connected 1,2 ? = %v \n", intS.connected(1,2) )
	fmt.Printf ("connected 1,7 ? = %v \n", intS.connected(1,7) )
}