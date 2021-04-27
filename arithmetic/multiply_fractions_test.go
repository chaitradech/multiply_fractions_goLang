package fractions

import (
	"testing"
)

// TestComputeGCD tests the functionality of computeGCD method.

func TestComputeGCD(t *testing.T) {
	var tests = []struct {
		a int64
		b int64
		n int64
	}{
		{1, 1, 1},
		{2, 1, 1},
		{3, 2, 1},
		{4, 3, 1},
		{5, 5, 5},
		{6, 8, 2},
		{7, 13, 1},
	}
	for _, test := range tests {
		got := computeGCD(test.a, test.b)
		if got == test.n {
			return
		} else {
			t.Errorf("Didn't match expected result for computeGCD(%d,%d) = Got: %d , Expected: %d;", test.a, test.b, test.n, got)
		}
	}
}

// TestReduce tests the functionality of reduce method.

func TestReduce(t *testing.T) {
	var tests = []struct {
		a Fraction
		x int64
		y int64
	}{
		{Fraction{5, 5}, 1, 1},
		{Fraction{13, 73}, 13, 73},
		{Fraction{6, 8}, 2, 1},
	}
	for _, test := range tests {
		got1, got2 := reduce(test.a)
		if got1 == test.x && got2 == test.y {
			return
		} else {
			t.Errorf("Didn't match expected result for reduce(%d) = Got: %d, %d, expected: %d, %d;", test.a, got1, got2, test.x, test.y)
		}
	}
}

// TestMultiplyFractions tests the functionality of MultiplyFractions method.
func TestMultiplyFractions(t *testing.T) {
	var tests = []struct {
		a      Fraction
		b      Fraction
		result Fraction
	}{
		{Fraction{19, 17}, Fraction{13, 47}, Fraction{247, 799}},
		{Fraction{55, 23}, Fraction{21, 18}, Fraction{385, 138}},
		{Fraction{89, 10}, Fraction{14, 17}, Fraction{623, 85}},
		{Fraction{3, 6}, Fraction{5, 9}, Fraction{5, 18}},
		{Fraction{6, 5}, Fraction{4, 10}, Fraction{12, 25}},
		{Fraction{6, 10}, Fraction{4, 10}, Fraction{6, 25}},
		{Fraction{7, 3}, Fraction{5, 10}, Fraction{7, 6}},
		{Fraction{1, 2}, Fraction{1, 2}, Fraction{1, 4}},
		{Fraction{1, 2}, Fraction{1, 3}, Fraction{1, 6}},
	}

	for _, test := range tests {
		total := MultiplyFractions(test.a, test.b)
		if total.Numerator == test.result.Numerator && total.Denominator == test.result.Denominator {
			return
		} else {
			t.Errorf("Didn't match expected result for MultiplyFractions(%d,%d) = Got: %d , Expected: %d;", test.a, test.b, test.result, total)
		}
	}
}
