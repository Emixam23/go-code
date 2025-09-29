package valid_sudoku

import (
	"strings"
	"testing"
)

func Test_isValidSudoku(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				board: [][]byte{
					[]byte(strings.Join([]string{"1", "2", ".", ".", "3", ".", ".", ".", "."}, "")),
					[]byte(strings.Join([]string{"4", ".", ".", "5", ".", ".", ".", ".", "."}, "")),
					[]byte(strings.Join([]string{".", "9", "8", ".", ".", ".", ".", ".", "3"}, "")),
					[]byte(strings.Join([]string{"5", ".", ".", ".", "6", ".", ".", ".", "4"}, "")),
					[]byte(strings.Join([]string{".", ".", ".", "8", ".", "3", ".", ".", "5"}, "")),
					[]byte(strings.Join([]string{"7", ".", ".", ".", "2", ".", ".", ".", "6"}, "")),
					[]byte(strings.Join([]string{".", ".", ".", ".", ".", ".", "2", ".", "."}, "")),
					[]byte(strings.Join([]string{".", ".", ".", "4", "1", "9", ".", ".", "8"}, "")),
					[]byte(strings.Join([]string{".", ".", ".", ".", "8", ".", ".", "7", "9"}, "")),
				},
			},
			want: true,
		},
		{
			name: "test2 - same row",
			args: args{
				board: [][]byte{
					[]byte(strings.Join([]string{"1", "2", ".", ".", "3", "2", ".", ".", "."}, "")),
					[]byte(strings.Join([]string{"4", ".", ".", "5", ".", ".", ".", ".", "."}, "")),
					[]byte(strings.Join([]string{".", "9", "8", ".", ".", ".", ".", ".", "3"}, "")),
					[]byte(strings.Join([]string{"5", ".", ".", ".", "6", ".", ".", ".", "4"}, "")),
					[]byte(strings.Join([]string{".", ".", ".", "8", ".", "3", ".", ".", "5"}, "")),
					[]byte(strings.Join([]string{"7", ".", ".", ".", "2", ".", ".", ".", "6"}, "")),
					[]byte(strings.Join([]string{".", ".", ".", ".", ".", ".", "2", ".", "."}, "")),
					[]byte(strings.Join([]string{".", ".", ".", "4", "1", "9", ".", ".", "8"}, "")),
					[]byte(strings.Join([]string{".", ".", ".", ".", "8", ".", ".", "7", "9"}, "")),
				},
			},
			want: false,
		},
		{
			name: "test3 - same column",
			args: args{
				board: [][]byte{
					[]byte(strings.Join([]string{"1", "2", ".", ".", "3", ".", ".", ".", "."}, "")),
					[]byte(strings.Join([]string{"4", ".", ".", "5", ".", ".", ".", "7", "."}, "")),
					[]byte(strings.Join([]string{".", "9", "8", ".", ".", ".", ".", ".", "3"}, "")),
					[]byte(strings.Join([]string{"5", ".", ".", ".", "6", ".", ".", ".", "4"}, "")),
					[]byte(strings.Join([]string{".", ".", ".", "8", ".", "3", ".", ".", "5"}, "")),
					[]byte(strings.Join([]string{"7", ".", ".", ".", "2", ".", ".", ".", "6"}, "")),
					[]byte(strings.Join([]string{".", ".", ".", ".", ".", ".", "2", ".", "."}, "")),
					[]byte(strings.Join([]string{".", ".", ".", "4", "1", "9", ".", ".", "8"}, "")),
					[]byte(strings.Join([]string{".", ".", ".", ".", "8", ".", ".", "7", "9"}, "")),
				},
			},
			want: false,
		},
		{
			name: "test4 - same big case",
			args: args{
				board: [][]byte{
					[]byte(strings.Join([]string{"1", "2", ".", ".", "3", ".", ".", ".", "."}, "")),
					[]byte(strings.Join([]string{"4", ".", ".", "5", ".", ".", ".", ".", "."}, "")),
					[]byte(strings.Join([]string{".", "9", "8", ".", ".", ".", ".", ".", "3"}, "")),
					[]byte(strings.Join([]string{"5", ".", ".", ".", "6", ".", ".", ".", "4"}, "")),
					[]byte(strings.Join([]string{".", ".", ".", "8", ".", "3", ".", ".", "5"}, "")),
					[]byte(strings.Join([]string{"7", ".", ".", ".", "2", ".", ".", ".", "6"}, "")),
					[]byte(strings.Join([]string{".", ".", ".", ".", ".", ".", "8", ".", "."}, "")),
					[]byte(strings.Join([]string{".", ".", ".", "4", "1", "9", ".", ".", "8"}, "")),
					[]byte(strings.Join([]string{".", ".", ".", ".", "8", ".", ".", "7", "9"}, "")),
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSudoku(tt.args.board); got != tt.want {
				t.Errorf("isValidSudoku() = %v, want %v", got, tt.want)
			}
		})
	}
}
