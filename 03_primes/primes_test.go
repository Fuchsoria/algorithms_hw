package primes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrimes(t *testing.T) {
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.output, CountPrimes(IsPrime, tt.inputN))
		})
	}
}

func TestPrimesSqrt(t *testing.T) {
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.output, CountPrimes(IsPrimeSqrt, tt.inputN))
		})
	}
}
