package captcha

import "testing"

func Test_SumNext(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1122",
			args{nums: []int{1, 1, 2, 2}},
			3,
		},
		{
			"1111",
			args{nums: []int{1, 1, 1, 1}},
			4,
		},
		{
			"1234",
			args{nums: []int{1, 2, 3, 4}},
			0,
		},
		{
			"91212129",
			args{nums: []int{9, 1, 2, 1, 2, 1, 2, 9}},
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumNext(tt.args.nums...); got != tt.want {
				t.Errorf("sumNext() = %v, want %v", got, tt.want)
			}
		})
	}
}
