package dyncon

type QuickUnionUF struct {
	Elements []int
}

func initQuickUnionUF(size int) *QuickUnionUF {
	quUF := QuickUnionUF{Elements: make([]int, size)}
	for i := range quUF.Elements {
		quUF.Elements[i] = i
	}
	return &quUF
}

func (quUF QuickUnionUF) root(i int) int {
	for i != quUF.Elements[i] {
		i = quUF.Elements[i]
	}
	return i
}

func (quUF QuickUnionUF) connected(p, q int) bool {
	return quUF.root(p) == quUF.root(q)
}

func (quUF *QuickUnionUF) union(p, q int) {
	i := quUF.root(p)
	j := quUF.root(q)
	quUF.Elements[i] = j

}
