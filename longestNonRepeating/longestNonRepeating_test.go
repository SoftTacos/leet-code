package longestnonrepeating

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "",
			want: 0,
		},
		{
			name: "asdfghjk",
			want: 8,
		},
		{
			name: "qwertqwertqwer",
			want: 5,
		},
		{
			name: "abcabcbb",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.name); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
