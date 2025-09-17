package is_anagram

import "testing"

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				s: "rat",
				t: "car",
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				s: "a",
				t: "ab",
			},
			want: false,
		},
		{
			name: "test4",
			args: args{
				s: "listen",
				t: "silent",
			},
			want: true,
		},
		{
			name: "test5",
			args: args{
				s: "",
				t: "",
			},
			want: true,
		},
		{
			name: "test6_same_letters_different_order",
			args: args{
				s: "abcde",
				t: "edcba",
			},
			want: true,
		},
		{
			name: "test7_different_length",
			args: args{
				s: "abc",
				t: "abcd",
			},
			want: false,
		},
		{
			name: "test8_repeated_letters",
			args: args{
				s: "aabbcc",
				t: "baccab",
			},
			want: true,
		},
		{
			name: "test9_one_empty_one_nonempty",
			args: args{
				s: "",
				t: "a",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram(%s, %s) = %v, want %v", tt.args.s, tt.args.t, got, tt.want)
			}
		})
	}
}
