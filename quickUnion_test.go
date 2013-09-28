package dyncon

import "testing"

func TestSqrt(t *testing.T) {
	const size = 10
	qf := initQuickFindUF(size)
	qf.union(1, 3)
	qf.union(1, 5)
	qf.union(1, 2)
	qf.union(7, 3)
	qf.union(4, 9)

	if x := qf.connected(1, 7); x != true {
		t.Errorf("Error")
	}
}
