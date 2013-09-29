package dyncon

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

func TestSqrt(t *testing.T) {
	file, err := os.Open("mediumUF.txt")
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
