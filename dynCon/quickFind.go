package quickfind


type QuickFindUF struct {
	Elements []int	
}

func initQuickFindUF(size int) *QuickFindUF {
	qfUF := QuickFindUF{ Elements: make([]int, size) }
	for i := range qfUF.Elements {
		qfUF.Elements[i] = i
	}
	return &qfUF
}

func (qfUF QuickFindUF) connected(p, q int) bool {
	return qfUF.Elements[p] == qfUF.Elements[q]

}

func (qfUF *QuickFindUF) union(p, q int) {
	pid := qfUF.Elements[p]
	qid := qfUF.Elements[q]
	for i := range qfUF.Elements {
		if qfUF.Elements[i] == pid {
			qfUF.Elements[i] = qid
		}
	}

}
