package dynCon

type QuickUnionWeightedByHeightUF struct {
	Elements   []int
	TreeHeight []int
}

func initQuickWeightedByHeightUF(size int) *QuickUnionWeightedByHeightUF {
	quwUF := QuickUnionWeightedByHeightUF{Elements: make([]int, size),
		TreeHeight: make([]int, size)}
	for i := range quwUF.Elements {
		quwUF.Elements[i] = i
		quwUF.TreeHeight[i] = 0
	}
	return &quwUF
}

func (quwUF QuickUnionWeightedByHeightUF) root(i int) int {
	for i != quwUF.Elements[i] {
		i = quwUF.Elements[i]
	}
	return i
}

func (quwUF QuickUnionWeightedByHeightUF) connected(p, q int) bool {
	return quwUF.root(p) == quwUF.root(q)
}

func (quwUF *QuickUnionWeightedByHeightUF) union(p, q int) {
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
