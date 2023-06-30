package firstmissingpositive

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	tests := []struct {
		name        string
		nums        []int
		wantMissing int
	}{
		{
			name:        "zero",
			nums:        []int{0},
			wantMissing: 1,
		},
		{
			name:        "negatives",
			nums:        []int{0, -1, -5, 2, 1, -5, 11},
			wantMissing: 3,
		},
		{
			name:        "perfect array",
			nums:        []int{1, 2, 3, 4, 5},
			wantMissing: 6,
		},
		{
			name:        "all negatives",
			nums:        []int{-1, -2, -3},
			wantMissing: 1,
		},
		{
			name:        "idk",
			nums:        []int{7, 8, 9, 11, 12},
			wantMissing: 1,
		},
		{
			name:        "streak",
			nums:        []int{1, 2, 3, 11, 12},
			wantMissing: 4,
		},
		{
			name:        "len",
			nums:        []int{3, 4, -1, 1},
			wantMissing: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMissing := firstMissingPositive(tt.nums); gotMissing != tt.wantMissing {
				t.Errorf("firstMissingPositive() = %v, want %v", gotMissing, tt.wantMissing)
			}
		})
	}
}

/*

[0,-1,-5,2,1,-5,11]
[-8,-8,8,2,1,8,11]


*/
