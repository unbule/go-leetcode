package greedy

import "testing"

func Test_findContentChildren(t *testing.T) {
	tests := []struct {
		name    string
		greedy  []int
		cookies []int
		expect  int
	}{
		{"case1", []int{1, 2, 3}, []int{1, 1}, 1},
		{"case2", []int{1, 2}, []int{1, 2, 3}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if real := findContentChildren(tt.greedy, tt.cookies); real == tt.expect {
				t.Log(tt.name, "ok")
			} else {
				t.Fatal(tt.name, "expect:", tt.expect, "real", real)
			}
		})
	}
}
