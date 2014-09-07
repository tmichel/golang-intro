// to run this: go test ./table_driven_test.go
package adder

import (
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	var tests = []struct {
		a, b     int
		expected int
		message  string
	}{
		{2, 3, 5, "adding positive numbers"},
		{-2, -3, -5, "adding negative numbers"},
		{1, 0, 1, "adding zero does not change the value"},
	}

	for _, tt := range tests {
		res := Add(tt.a, tt.b)
		if tt.expected != res {
			t.Fatalf("in '%s' expected: %d, got: %d", tt.message, tt.expected, res)
		}
	}
}
