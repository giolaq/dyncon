package dynCon

type QuickUnionWeightedCompressedPathUF struct {
	Elements []int
	TreeSize []int
}

func initQuickWeightedUF(size int) *QuickUnionWeightedCompressedPathUF {
	quwCUF := QuickUnionWeightedCompressedPathUF{Elements: make([]int, size),
		TreeSize: make([]int, size)}
	for i := range quwCUF.Elements {
		quwCUF.Elements[i] = i
		quwCUF.TreeSize[i] = 1
	}
	return &quwCUF
}

func (quwCUF QuickUnionWeightedCompressedPathUF) root(i int) int {
	for i != quwCUF.Elements[i] {
		i = quwCUF.Elements[i]
	}
	return i
}

func (quwCUF QuickUnionWeightedCompressedPathUF) connected(p, q int) bool {
	return quwCUF.root(p) == quwCUF.root(q)
}

func (quwCUF *QuickUnionWeightedCompressedPathUF) union(p, q int) {
	i := quwCUF.root(p)
	j := quwCUF.root(q)
	if quwCUF.TreeSize[i] < quwCUF.TreeSize[j] {
		quwCUF.Elements[i] = j
		quwCUF.TreeSize[j] = quwCUF.TreeSize[j] + quwCUF.TreeSize[i]
	} else {
		quwCUF.Elements[j] = i
		quwCUF.TreeSize[i] = quwCUF.TreeSize[i] + quwCUF.TreeSize[j]

	}

}
