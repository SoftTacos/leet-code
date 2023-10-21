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
			name: "doo",
			s:    "doo",
			want: "oo",
		},
		{
			name: "ood",
			s:    "ood",
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
