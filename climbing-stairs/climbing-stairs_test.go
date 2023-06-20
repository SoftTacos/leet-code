package climbingstairs

import "testing"

func Test_climbStairs(t *testing.T) {
	tests := []struct {
		name             string
		n                int
		wantCombinations int
	}{
		{
			name:             "1",
			n:                1,
			wantCombinations: 1,
		},
		{
			name:             "2",
			n:                2,
			wantCombinations: 2,
		},
		{
			name:             "3",
			n:                3,
			wantCombinations: 3,
		},
		{
			name:             "4",
			n:                4,
			wantCombinations: 5,
		},
		{
			name:             "42",
			n:                42,
			wantCombinations: 433494437,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCombinations := climbStairs(tt.n); gotCombinations != tt.wantCombinations {
				t.Errorf("climbStairs() = %v, want %v", gotCombinations, tt.wantCombinations)
			}
		})
	}
}
