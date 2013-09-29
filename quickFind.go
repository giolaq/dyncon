package dyncon

type Interface interface {
	Union(p, q int)
	Connected(p, q int) bool
	Find(p int) int
	
}


type QuickFind struct {
	Elements []int
}

func initQuickFind(size int) *QuickFind {
	qfUF := QuickFind{Elements: make([]int, size)}
	for i := range qfUF.Elements {
		qfUF.Elements[i] = i
	}
	return &qfUF
}

func (qfUF QuickFind) Connected(p, q int) bool {
	return qfUF.Elements[p] == qfUF.Elements[q]

}

func (qfUF *QuickFind) Union(p, q int) {
	pid := qfUF.Elements[p]
	qid := qfUF.Elements[q]
	for i := range qfUF.Elements {
		if qfUF.Elements[i] == pid {
			qfUF.Elements[i] = qid
		}
	}

}

func (qf *QuickFind) Find(p int) int {
	return qf.Elements[p]
}
