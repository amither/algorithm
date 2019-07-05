package probability

import "testing"

func Test_gen_random_without_repeating(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"test1", args{10, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen_random_without_repeating(tt.args.m, tt.args.n)
		})
	}
}
