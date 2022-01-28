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
			"1212",
			args{nums: []int{1, 2, 1, 2}},
			6,
		},
		{
			"1221",
			args{nums: []int{1, 2, 2, 1}},
			0,
		},
		{
			"123425",
			args{nums: []int{1, 2, 3, 4, 2, 5}},
			4,
		},
		{
			"123123",
			args{nums: []int{1, 2, 3, 1, 2, 3}},
			12,
		},
		{
			"12131415",
			args{nums: []int{1, 2, 1, 3, 1, 4, 1, 5}},
			4,
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
