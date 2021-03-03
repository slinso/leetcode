package p0010

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
			name: "ex1",
			args: args{
				s: "aa",
				p: "a",
			},
			want: false,
		},
		{
			name: "ex2",
			args: args{
				s: "aa",
				p: "a*",
			},
			want: true,
		},
		{
			name: "ex3",
			args: args{
				s: "ab",
				p: ".*",
			},
			want: true,
		},
		{
			name: "ex4",
			args: args{
				s: "aab",
				p: "c*a*b",
			},
			want: true,
		},
		{
			name: "ex5",
			args: args{
				s: "mississippi",
				p: "mis*is*p*.",
			},
			want: false,
		},
		{
			name: "ex6",
			args: args{
				s: "abc",
				p: "a.c",
			},
			want: true,
		},
		{
			name: "ex7",
			args: args{
				s: "ab",
				p: ".*c",
			},
			want: false,
		},
		{
			name: "ex8",
			args: args{
				s: "aaa",
				p: "a*a",
			},
			want: true,
		},
		{
			name: "ex9",
			args: args{
				s: "a",
				p: "ab*a",
			},
			want: false,
		},
		{
			name: "ex10",
			args: args{
				s: "aaa",
				p: "ab*a*c*a",
			},
			want: true,
		},
		{
			name: "ex11",
			args: args{
				s: "ab",
				p: ".*..",
			},
			want: true,
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
