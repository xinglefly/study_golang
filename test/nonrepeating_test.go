package test

import "testing"

func TestSubstr(t *testing.T) {

	tests := []struct {
		s   string
		ans int
	}{
		{"abcabcbb", 3},
		{"abcabcabcd", 4},
		{"", 0},
		{"bbbbbb", 1},
		{"abcabcchelowrdacbxingle", 14},
		{"马丛是不是个傻子马丛是", 7},
	}

	for _, tt := range tests {
		if actural := lengthOfNonRepeatingSubStr(tt.s); actural != tt.ans {
			t.Errorf("input s %s, got %d, expected %d", tt.s, actural, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "abcabcchelowrdacbxingle"
	ans := 14

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("input s %s, got %d, expected %d", s, actual, ans)
		}
	}
}
