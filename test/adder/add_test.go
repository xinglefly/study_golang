package adder

import "testing"

//表格驱动测试  go test .
//代码覆盖率    go test -coverprofile=c.out
//			  go tool cover -html=c.out
//代码性能      go test -bench .
func TestAdder(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{1, 2, 3},
		{3, 6, 9},
		{2, 4, 6},
		{1, 8, 9},
		{-10, 1, -9},
	}

	for _, tt := range tests {
		if actual := add(tt.a, tt.b); actual != tt.c {
			t.Errorf("adder(%d,%d); got %d; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}
