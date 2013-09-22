package dynCon

type QuickUnionWeightedUF struct {
	Elements   []int
	TreeHeight []int
}

func initQuickWeightedUF(size int) *QuickUnionWeightedUF {
	quwUF := QuickUnionWeightedUF{Elements: make([]int, size),
		TreeHeight: make([]int, size)}
	for i := range quwUF.Elements {
		quwUF.Elements[i] = i
		quwUF.TreeHeight[i] = 0
	}
	return &quwUF
}

func (quwUF QuickUnionWeightedUF) root(i int) int {
	for i != quwUF.Elements[i] {
		i = quwUF.Elements[i]
	}
	return i
}

func (quwUF QuickUnionWeightedUF) connected(p, q int) bool {
	return quwUF.root(p) == quwUF.root(q)
}

func (quwUF *QuickUnionWeightedUF) union(p, q int) {
	i := quwUF.root(p)
	j := quwUF.root(q)
	if quwUF.TreeHeight[i] < quwUF.TreeHeight[j] {
		quwUF.Elements[i] = j
	} else if quwUF.TreeHeight[i] > quwUF.TreeHeight[j] {
		quwUF.Elements[j] = i
	} else {
		quwUF.Elements[j] = i
		quwUF.TreeHeight[i]++
	}

}
