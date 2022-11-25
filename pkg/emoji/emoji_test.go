package emoji

import "testing"

func Test_translateSymbol(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "translate: exclamation",
			args: args{s: "!"},
			want: "exclamation",
		},
		{
			name: "translate: question",
			args: args{s: "?"},
			want: "question",
		},
		{
			name: "translate: asterisk",
			args: args{s: "*"},
			want: "asterisk",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateSymbol(tt.args.s); got != tt.want {
				t.Errorf("translateSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}
