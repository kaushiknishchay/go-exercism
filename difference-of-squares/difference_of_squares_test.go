package diffsquares

import "testing"

var tests = []struct{ n, sqOfSum, sumOfSq uint64 }{
	{5, 225, 55},
	{10, 3025, 385},
	{100, 25502500, 338350},
	{132400, 3037305594056233536, 773655506235400},
}

func TestSquareOfSum(t *testing.T) {
	for _, test := range tests {
		if s := SquareOfSum(test.n); s != test.sqOfSum {
			t.Fatalf("SquareOfSum(%d) = %d, want %d", test.n, s, test.sqOfSum)
		}
	}
}

func TestSumOfSquares(t *testing.T) {
	for _, test := range tests {
		if s := SumOfSquares(test.n); s != test.sumOfSq {
			t.Fatalf("SumOfSquares(%d) = %d, want %d", test.n, s, test.sumOfSq)
		}
	}
}

func TestDifference(t *testing.T) {
	for _, test := range tests {
		want := test.sqOfSum - test.sumOfSq
		if s := Difference(test.n); s != want {
			t.Fatalf("Difference(%d) = %d, want %d", test.n, s, want)
		}
	}
}

// Benchmark functions on just a single number (100, from the original PE problem)
// to avoid overhead of iterating over tests.
func BenchmarkSquareOfSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SquareOfSum(100)
	}
}

func BenchmarkSumOfSquares(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumOfSquares(100)
	}
}

func BenchmarkDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Difference(100)
	}
}
