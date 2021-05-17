package main

import "testing"

func Test_sumForMathSeq(t *testing.T) {
	type args struct {
		from  int
		count int
		step  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "one",
			args: args{
				from:  1,
				count: 3,
				step:  1,
			},
			want: 6,
		},
		{
			name: "two",
			args: args{
				from:  1,
				count: 3,
				step:  2,
			},
			want: 9,
		},
		{
			name: "threee",
			args: args{
				from:  1,
				count: 3,
				step:  3,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumForMathSeq(tt.args.from, tt.args.count, tt.args.step); got != tt.want {
				t.Errorf("sumForMathSeq() = %v, want %v", got, tt.want)
			}
		})
	}
}
