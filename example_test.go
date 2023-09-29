package driven_test

import (
	"fmt"
	"testing"
)

func ExampleRelu() {
	// Output:
}

var reluIntCases = []struct {
	val      int
	expected int
}{
	{-1, 0},
	{0, 0},
	{1, 1},
}

func TestReluInt(t *testing.T) {
	for _, tc := range reluIntCases {
		name := fmt.Sprintf("%v", tc.val)
		t.Run(name, func(t *testing.T) {
			out := Relu(tc.val)
			if out != tc.expected {
				t.Fatalf("expected: %#v, got: %#v", tc.expected, out)
			}
		})
	}
}
