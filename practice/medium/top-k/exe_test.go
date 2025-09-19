package top_k

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			want: []int{1, 2},
		},
		{
			name: "test2",
			args: args{
				nums: []int{1},
				k:    1,
			},
			want: []int{1},
		},
		{
			name: "test3",
			args: args{
				nums: []int{4, 1, -1, 2, -1, 2, 3},
				k:    2,
			},
			want: []int{-1, 2},
		},
		{
			name: "test4",
			args: args{
				nums: []int{1, 2},
				k:    2,
			},
			want: []int{1, 2},
		},
		{
			name: "test5",
			args: args{
				nums: []int{5, 3, 1, 1, 1, 3, 73, 1},
				k:    2,
			},
			want: []int{1, 3},
		},
		{
			name: "test6",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				k:    10,
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "test7",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4},
				k:    3,
			},
			want: []int{1, 4, 3},
		},
		{
			name: "test8",
			args: args{
				nums: []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4},
				k:    2,
			},
			want: []int{4, 3},
		},
		{
			name: "test9",
			args: args{
				nums: []int{1, 1, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 5},
				k:    4,
			},
			want: []int{1, 5, 2, 3},
		},
		{
			name: "test10",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5},
				k:    5,
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				require.ElementsMatchf(t, tt.want, got, "topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
