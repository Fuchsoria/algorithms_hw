package fibonacci

var Cases = []struct {
	name   string
	inputN float64
	output float64
}{
	{
		name:   "N 0",
		inputN: 0,
		output: 0,
	},
	{
		name:   "N 1",
		inputN: 1,
		output: 1,
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
		output: 3,
	},
	{
		name:   "N 5",
		inputN: 5,
		output: 5,
	},
	{
		name:   "N 6",
		inputN: 10,
		output: 55,
	},
	{
		name:   "N 7",
		inputN: 100,
		output: 354224848179261915075,
	},
}
