package a2019

import "testing"

type d7test struct {
	input  []int
	phase  []int
	output int
}

var d7tests = []d7test{
	{
		input:  []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
		phase:  []int{0, 1, 2, 3, 4},
		output: 43210,
	},
	{
		input: []int{33, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23,
			101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
		phase:  []int{0, 1, 2, 3, 4},
		output: 54321,
	},
	{
		input: []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33,
			1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
		phase:  []int{0, 1, 2, 3, 4},
		output: 65210,
	},
	{
		input: []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26,
			27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
		phase:  []int{5, 6, 7, 8, 9},
		output: 139629729,
	},
	{
		input: []int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54,
			-5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4,
			53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10},
		phase:  []int{5, 6, 7, 8, 9},
		output: 18216,
	},
}

func TestBaseCalc(t *testing.T) {
	res := calcMax([]int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}, []int{4, 3, 2, 1, 0})
	if res != 43210 {
		t.Errorf("Bad res: %v", res)
	}
}

func TestBasicPhase(t *testing.T) {
	for _, te := range d7tests {
		mv := FindMax(te.input, te.phase)
		if mv != te.output {
			t.Errorf("Better max: %v (should have been %v)", mv, te.output)
		}
	}
}
