package regularexpressionmatching

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nothing",
			args: args{
				p: "",
				s: "",
			},
			want: true,
		},
		{
			name: "simple match",
			args: args{
				p: "simple",
				s: "simple",
			},
			want: true,
		},
		{
			name: "simple doesn't match string",
			args: args{
				p: "simple",
				s: "simples",
			},
			want: false,
		},
		{
			name: "simple doesnt match pattern",
			args: args{
				p: "simples",
				s: "simple",
			},
			want: false,
		},
		{
			name: "wildcard",
			args: args{
				p: "simple.",
				s: "simples",
			},
			want: true,
		},
		{
			name: "repeating single character",
			args: args{
				p: "s*imples",
				s: "ssssimples",
			},
			want: true,
		},
		{
			name: "repeating missing single character",
			args: args{
				p: "bs*imples",
				s: "bimples",
			},
			want: true,
		},
		{
			name: "bad rep single character",
			args: args{
				p: "s*imples",
				s: "ssssimpless",
			},
			want: false,
		},
		{
			name: "post repeat not match",
			args: args{
				p: "s*mples",
				s: "ssssimples",
			},
			want: false,
		},
		{
			name: "rep. mid",
			args: args{
				p: "s.*s",
				s: "simples",
			},
			want: true,
		},
		{
			name: "rep. at end",
			args: args{
				p: "s.*",
				s: "simple",
			},
			want: true,
		},
		{
			name: "rep. at start",
			args: args{
				p: ".*s",
				s: "simples",
			},
			want: true,
		},
		{
			name: "bad rep. at start",
			args: args{
				p: ".*s",
				s: "simplest",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

// [1,2,3,4,5]

// [1,9,1,1,9,1]
