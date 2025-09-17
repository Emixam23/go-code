package twoSum

import "testing"

func Test_hasDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: false,
		},
		{
			name: "test2",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 3},
			},
			want: true,
		},
		{
			name: "test3",
			args: args{
				nums: []int{},
			},
			want: false,
		},
		{
			name: "test4",
			args: args{
				nums: []int{1},
			},
			want: false,
		},
		{
			name: "test5",
			args: args{
				nums: []int{1, 1, 1, 1, 1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("hasDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
