package main

import (
	"reflect"
	"testing"
)

func Test_transformInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"1122",
			args{input: "1122"},
			[]int{1, 1, 2, 2},
		},
		{
			"1111",
			args{input: "1111"},
			[]int{1, 1, 1, 1},
		},
		{
			"1234",
			args{input: "1234"},
			[]int{1, 2, 3, 4},
		},
		{
			"91212129",
			args{input: "91212129"},
			[]int{9, 1, 2, 1, 2, 1, 2, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transformInput(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transformInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
