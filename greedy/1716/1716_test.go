package greedy

import "testing"

func Test_totalMoney(t *testing.T) {
	tests := []struct {
		name   string
		n      int
		expect int
	}{
		{"case1", 4, 10},
		{"case2", 10, 37},
		{"case3", 20, 96},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if real := totalMoney(tt.n); tt.expect != real {
				t.Fatal("real:", real, "expect:", tt.expect)
			}
		})
	}
}
