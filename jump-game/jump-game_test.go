package jumpgame

import "testing"

func Test_canJump(t *testing.T) {
	tests := []struct {
		name        string
		nums        []int
		wantCanJump bool
	}{
		{
			name:        "[0]",
			nums:        []int{0},
			wantCanJump: true,
		},
		{
			name:        "tiny can't",
			nums:        []int{0, 1},
			wantCanJump: false,
		},
		{
			name:        "simple",
			nums:        []int{1, 1},
			wantCanJump: true,
		},
		{
			name:        "123",
			nums:        []int{1, 2, 1},
			wantCanJump: true,
		},
		{
			name:        "middle1",
			nums:        []int{1, 0, 1},
			wantCanJump: false,
		},
		{
			name:        "middle2",
			nums:        []int{3, 2, 1, 0, 4},
			wantCanJump: false,
		},
		{
			name:        "first",
			nums:        []int{7, 0, 0, 0, 0, 0, 1},
			wantCanJump: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCanJump := canJump(tt.nums); gotCanJump != tt.wantCanJump {
				t.Errorf("canJump() = %v, want %v", gotCanJump, tt.wantCanJump)
			}
		})
	}
}
