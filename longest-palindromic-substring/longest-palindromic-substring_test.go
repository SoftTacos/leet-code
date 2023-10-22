package longest

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {

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
			name: "a",
			s:    "a",
			want: "a",
		},
		{
			name: "doo",
			s:    "doo",
			want: "oo",
		},
		{
			name: "doodood",
			s:    "doodood",
			want: "doodood",
		},
		{
			name: "ccc",
			s:    "ccc",
			want: "ccc",
		},
		{
			name: "asdfcccasdf",
			s:    "asdfcccasdf",
			want: "ccc",
		},
		{
			name: "ood",
			s:    "ood",
			want: "oo",
		},
		{
			name: "dood",
			s:    "dood",
			want: "dood",
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
			name: "asdfasdf",
			s:    "asdfasdf",
			want: "a",
		},
		{
			name: "12345",
			s:    "543212345asa543212345",
			want: "543212345asa543212345",
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
