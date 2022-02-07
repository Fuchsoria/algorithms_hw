package lucky

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestN(t *testing.T) {
	cases := []struct {
		name   string
		input  int
		output int
	}{
		{
			name:   "N 1",
			input:  1,
			output: 10,
		},
		{
			name:   "N 2",
			input:  2,
			output: 670,
		},
		{
			name:   "N 3",
			input:  3,
			output: 55252,
		},
		{
			name:   "N 4",
			input:  4,
			output: 4816030,
		},
	}

	lucky := Lucky{}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			require.Equal(t, tt.output, lucky.GetLuckyCount(tt.input))
		})
	}
}
