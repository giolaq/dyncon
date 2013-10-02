package dyncon

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

type union func(int, int)

func TestQuickFind(t *testing.T) {
	file, err := os.Open("tinyUF.txt")
	if err != nil {
		t.Errorf("Error opening File")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())

	qf := initQuickFind(size)

	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		y, _ := strconv.Atoi(scanner.Text())
		qf.Union(x, y)
	}

	if conn := qf.Connected(6, 1); conn != true {
		t.Errorf("Error")
	}

}

func TestQuickUnion(t *testing.T) {
	file, err := os.Open("tinyUF.txt")
	if err != nil {
		t.Errorf("Error opening File")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())

	uf := initQuickUnionUF(size)

	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		y, _ := strconv.Atoi(scanner.Text())
		uf.union(x, y)
	}

	if conn := uf.connected(6, 1); conn != true {
		t.Errorf("Error")
	}
}
