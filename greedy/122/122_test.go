package greedy

import "testing"

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		expect int
		real   int
	}{
		{"case1", []int{7, 1, 5, 3, 6, 4}, 7, -1},
		{"case2", []int{1, 2, 3, 4, 5}, 4, -1},
		{"case3", []int{7, 6, 4, 3, 1}, 0, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.real = maxProfit(tt.prices); tt.expect != tt.real {
				t.Fatal("expect:", tt.expect, "real:", tt.real)
			}
		})
	}
}
