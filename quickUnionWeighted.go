package dyncon

type QuickUnionWeightedUF struct {
	Elements []int
	TreeSize []int
}

func initQuickWeightedUF(size int) *QuickUnionWeightedUF {
	quwUF := QuickUnionWeightedUF{Elements: make([]int, size),
		TreeSize: make([]int, size)}
	for i := range quwUF.Elements {
		quwUF.Elements[i] = i
		quwUF.TreeSize[i] = 1
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
	if quwUF.TreeSize[i] < quwUF.TreeSize[j] {
		quwUF.Elements[i] = j
		quwUF.TreeSize[j] = quwUF.TreeSize[j] + quwUF.TreeSize[i]
	} else {
		quwUF.Elements[j] = i
		quwUF.TreeSize[i] = quwUF.TreeSize[i] + quwUF.TreeSize[j]

	}

}
