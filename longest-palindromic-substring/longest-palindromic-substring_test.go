package longest

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	total := 1000
	charset := "12345asa54321"
	reallyLong := make([]byte, 500)
	reallyLong = append(reallyLong, []byte(charset)...)
	reallyLong = append(reallyLong, make([]byte, total-len(reallyLong))...)
	//12345asa54321
	//12345asa54321
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "",
			s:    "",
			want: "",
		},
		{
			name: "doo",
			s:    "doo",
			want: "oo",
		},
		{
			name: "dad",
			s:    "dad",
			want: "dad",
		},
		{
			name: "doodle",
			s:    "asdfasdfdoodle",
			want: "dood",
		},
		{
			name: "long",
			s:    string(reallyLong),
			want: "12345asa54321",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.s); got != tt.want {
				t.Errorf("longestPalindrome() = `%v`, want `%v`", got, tt.want)
			}
		})
	}
}
