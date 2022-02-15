package bits

import "testing"

func BenchmarkKingMoves(b *testing.B) {
	chess := Chess{}

	for _, bm := range CasesKing {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = chess.GetKingBitboardMoves(bm.input)
			}
		})
	}
}

func BenchmarkHorseMoves(b *testing.B) {
	chess := Chess{}

	for _, bm := range CasesKing {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = chess.GetHorseBitboardMoves(bm.input)
			}
		})
	}
}

func BenchmarkCount(b *testing.B) {
	chess := Chess{}

	for _, bm := range CasesKing {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = chess.Popcnt(chess.GetKingBitboardMoves(bm.input))
			}
		})
	}

	for _, bm := range CasesHorse {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = chess.Popcnt(chess.GetHorseBitboardMoves(bm.input))
			}
		})
	}
}

func BenchmarkOptCount(b *testing.B) {
	chess := Chess{}

	for _, bm := range CasesKing {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = chess.PopcntOpt(chess.GetKingBitboardMoves(bm.input))
			}
		})
	}

	for _, bm := range CasesHorse {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = chess.PopcntOpt(chess.GetHorseBitboardMoves(bm.input))
			}
		})
	}
}
