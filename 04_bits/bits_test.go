package bits

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKing(t *testing.T) {
	chess := Chess{}

	for _, tt := range CasesKing {
		t.Run(tt.name, func(t *testing.T) {
			mask := chess.GetKingBitboardMoves(tt.input)
			count := chess.Popcnt(mask)
			countOpt := chess.PopcntOpt(mask)

			require.Equal(t, tt.output, int(mask))
			require.Equal(t, tt.outputCount, count)
			require.Equal(t, tt.outputCount, countOpt)
		})
	}
}

func TestHorse(t *testing.T) {
	chess := Chess{}

	for _, tt := range CasesHorse {
		t.Run(tt.name, func(t *testing.T) {
			mask := chess.GetHorseBitboardMoves(tt.input)
			count := chess.Popcnt(mask)
			countOpt := chess.PopcntOpt(mask)

			require.Equal(t, tt.output, int(mask))
			require.Equal(t, tt.outputCount, count)
			require.Equal(t, tt.outputCount, countOpt)
		})
	}
}
