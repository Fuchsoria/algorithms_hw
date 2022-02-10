package primes

var Cases = []struct {
	name   string
	inputN int64
	output int
}{
	{
		name:   "N 10",
		inputN: 10,
		output: 4,
	},

	{
		name:   "N 1",
		inputN: 1,
		output: 0,
	},
	{
		name:   "N 2",
		inputN: 2,
		output: 1,
	},
	{
		name:   "N 3",
		inputN: 3,
		output: 2,
	},
	{
		name:   "N 4",
		inputN: 4,
		output: 2,
	},
	{
		name:   "N 5",
		inputN: 5,
		output: 3,
	},
	{
		name:   "N 100",
		inputN: 100,
		output: 25,
	},
	{
		name:   "N 1k",
		inputN: 1_000,
		output: 168,
	},
	{
		name:   "N 10k",
		inputN: 10_000,
		output: 1229,
	},
	{
		name:   "N 100k",
		inputN: 100_000,
		output: 9592,
	},
	// {
	// 	name:   "N 1kk",
	// 	inputN: 1_000_000,
	// 	output: 78498,
	// },
	// {
	// 	name:   "N 10kk",
	// 	inputN: 10_000_000,
	// 	output: 664579,
	// },
	// {
	// 	name:   "N 100kk",
	// 	inputN: 100_000_000,
	// 	output: 5761455,
	// },
	// {
	// 	name:   "N 1kkk",
	// 	inputN: 1_000_000_000,
	// 	output: 50847534,
	// },
	// {
	// 	name:   "N 10kkk",
	// 	inputN: 123_456_789,
	// 	output: 7027260,
	// },
}
