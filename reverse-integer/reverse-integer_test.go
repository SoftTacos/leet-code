package reverseinteger

import "testing"

func Test_reverse(t *testing.T) {
	tests := []struct {
		name  string
		x     int
		wantR int
	}{
		{
			name:  "0",
			x:     0,
			wantR: 0,
		},
		{
			name:  "1",
			x:     1,
			wantR: 1,
		},
		{
			name:  "1234",
			x:     1234,
			wantR: 4321,
		},
		{
			name:  "-1234",
			x:     -1234,
			wantR: -4321,
		},
		{
			name:  "-100001234",
			x:     -100001234,
			wantR: -432100001,
		},
		{
			name:  "max+1",
			x:     8463847412,
			wantR: 0,
		},
		{
			name:  "min-1",
			x:     -9463847412,
			wantR: 0,
		},
		//
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := reverse(tt.x); gotR != tt.wantR {
				t.Errorf("reverse() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
