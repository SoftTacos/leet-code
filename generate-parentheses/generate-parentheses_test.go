package generateparentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateParenthesis(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		wantAns []string
	}{
		{
			name:    "1",
			n:       1,
			wantAns: []string{"()"},
		},
		{
			name:    "2",
			n:       2,
			wantAns: []string{"(())", "()()"},
		},
		{
			name:    "3",
			n:       3,
			wantAns: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAns := generateParenthesis(tt.n)
			assert.EqualValues(t, tt.wantAns, gotAns)
		})
	}
}
