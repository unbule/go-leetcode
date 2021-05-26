package greedy

import (
	"strings"
	"testing"
)

func Test_reverseParentheses(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect string
	}{
		{"testcase1", "(abcd)", "dcba"},
		{"testcase2", "(u(love)i)", "iloveu"},
		{"testcase3", "(ed(et(oc))el)", "leetcode"},
		{"testcase4", "a(bcdefghijkl(mno)p)q", "apmnolkjihgfedcbq"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reverseParentheses(tt.input)
			if strings.Compare(result, tt.expect) != 0 {
				t.Fatal("expect:", tt.expect, "result", result)
			}
		})
	}
}
