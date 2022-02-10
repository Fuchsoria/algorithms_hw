package exponentiation

var Cases = []struct {
	name   string
	inputA float64
	inputN float64
	output float64
}{
	{
		name:   "N 10",
		inputA: 2,
		inputN: 10,
		output: 1024,
	},
	{
		name:   "N 0",
		inputA: 123456789,
		inputN: 0,
		output: 1,
	},
	{
		name:   "N 10k",
		inputA: 1.0001,
		inputN: 10_000,
		output: 2.7181459268248984,
	},
	{
		name:   "N 100k",
		inputA: 1.00001,
		inputN: 100_000,
		output: 2.718268237192295,
	},
	{
		name:   "N 1kk",
		inputA: 1.000001,
		inputN: 1_000_000,
		output: 2.7182804690959363,
	},
	{
		name:   "N 10kk",
		inputA: 1.0000001,
		inputN: 10_000_000,
		output: 2.7182816941320103,
	},
	{
		name:   "N 100kk",
		inputA: 1.00000001,
		inputN: 100_000_000,
		output: 2.71828179834636,
	},
	{
		name:   "N 1kkk",
		inputA: 1.000000001,
		inputN: 1_000_000_000,
		output: 2.7182820520118995,
	},
	{
		name:   "N 10kkk",
		inputA: 1.0000000001,
		inputN: 10_000_000_000,
		output: 2.7182820532331595,
	},
}
