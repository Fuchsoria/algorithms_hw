package exponentiation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestN(t *testing.T) {
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.output, Calculate(tt.inputA, tt.inputN))
		})
	}
}
