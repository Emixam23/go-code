package encode_and_decode

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution_Encode(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				strs: []string{"hello", "world"},
			},
		},
		{
			name: "test2",
			args: args{
				strs: []string{"", "world"},
			},
		},
		{
			name: "test2",
			args: args{
				strs: []string{"neet", "code", "love", "you"},
			},
		},
		{
			name: "test3",
			args: args{
				strs: []string{"we", "say", ":", "yes", "!@#$%^&*()"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Solution{}
			encoded := s.Encode(tt.args.strs)

			decoded := s.Decode(encoded)
			require.Equal(t, tt.args.strs, decoded)
		})
	}
}
