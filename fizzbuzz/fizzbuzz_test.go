package fizzbuzz

import "testing"

func TestRun(t *testing.T) {
	type args struct {
		n uint32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"in 1 out 1",
			args{n: 1},
			"1",
		},
		{
			"in 3 out Fizz",
			args{n: 3},
			"Fizz",
		},
		{
			"in 5 out Buzz",
			args{n: 5},
			"Buzz",
		},
		{
			"in 15 out FizzBuzz",
			args{n: 15},
			"FizzBuzz",
		},
		{
			"in 75 out FizzBuzz",
			args{n: 75},
			"FizzBuzz",
		},
		{
			"in 99 out Fizz",
			args{n: 99},
			"FizzBuzz",
		},
		{
			"in 100 out Buzz",
			args{n: 100},
			"Buzz",
		},
		{
			"in 101 out 101",
			args{n: 101},
			"101",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Run(tt.args.n); got != tt.want {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
