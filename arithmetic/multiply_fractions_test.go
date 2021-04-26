package fractions

import (
	"testing"
)

//TestComputeGCD tests the functionality of ComputeGCD.

func TestComputeGCD(t *testing.T) {
	got := ComputeGCD(6, 10)
	assertEqual(t, got, 2)
}

func assertEqual(t *testing.T, got int64, expected int64) {
	if got == expected {
		return
	} else {
		t.Errorf("computeGCD(6,10) = %d ; want 2", expected)
	}
}

//TestMultiplyFractions1 inputs a=1/2 and b=1/3, and expectes o/p res=1/6.
func TestMultiplyFractions1(t *testing.T) {
	a := Fraction{Numerator: 1, Denominator: 2}
	b := Fraction{Numerator: 1, Denominator: 3}

	res := MultiplyFractions(a, b)

	if res.Numerator == 1 && res.Denominator == 6 {
		return
	} else {
		t.Errorf("Didn't match expected result for MultiplyFractions(%d,%d) = %d ;", a, b, res)
	}
}

//TestMultiplyFractions2 inputs a=6/10 and b=4/10, and expectes o/p res=6/25.
func TestMultiplyFractions2(t *testing.T) {
	a := Fraction{Numerator: 6, Denominator: 10}
	b := Fraction{Numerator: 4, Denominator: 10}

	res := MultiplyFractions(a, b)

	if res.Numerator == 6 && res.Denominator == 25 {
		return
	} else {
		t.Errorf("Didn't match expected result for MultiplyFractions(%d,%d) = %d ;", a, b, res)
	}
}

//TestMultiplyFractions3 inputs a=6/5 and b=4/10, and expectes o/p res=12/25.
func TestMultiplyFractions3(t *testing.T) {
	a := Fraction{Numerator: 6, Denominator: 5}
	b := Fraction{Numerator: 4, Denominator: 10}

	res := MultiplyFractions(a, b)

	if res.Numerator == 12 && res.Denominator == 25 {
		return
	} else {
		t.Errorf("Didn't match expected result for MultiplyFractions(%d,%d) = %d ;", a, b, res)
	}
}

//TestMultiplyFractions4 inputs a=12/10 and b=10/12, and expectes o/p : 1/1.
func TestMultiplyFractions4(t *testing.T) {
	a := Fraction{Numerator: 12, Denominator: 10}
	b := Fraction{Numerator: 10, Denominator: 12}

	res := MultiplyFractions(a, b)

	if res.Numerator == 1 && res.Denominator == 1 {
		return
	} else {
		t.Errorf("Didn't match expected result for MultiplyFractions(%d,%d) = %d ;", a, b, res)
	}
}

//TestMultiplyFractions5 inputs a=3/6 and b=5/9, and expectes o/p res=5/18.
func TestMultiplyFractions5(t *testing.T) {
	a := Fraction{Numerator: 3, Denominator: 6}
	b := Fraction{Numerator: 5, Denominator: 9}

	res := MultiplyFractions(a, b)

	if res.Numerator == 5 && res.Denominator == 18 {
		return
	} else {
		t.Errorf("Didn't match expected result for MultiplyFractions(%d,%d) = %d ;", a, b, res)
	}
}

//TestMultiplyFractions6 inputs a=7/3 and b=5/10, and expectes o/p res=21/50.
func TestMultiplyFractions6(t *testing.T) {
	a := Fraction{Numerator: 7, Denominator: 5}
	b := Fraction{Numerator: 3, Denominator: 10}

	res := MultiplyFractions(a, b)

	if res.Numerator == 21 && res.Denominator == 50 {
		return
	} else {
		t.Errorf("Didn't match expected result for MultiplyFractions(%d,%d) = %d ;", a, b, res)
	}
}

//TestMultiplyFractions7 inputs a=89/10 and b=14/17, and expectes o/p res=623/85.
func TestMultiplyFractions7(t *testing.T) {
	a := Fraction{Numerator: 89, Denominator: 10}
	b := Fraction{Numerator: 14, Denominator: 17}

	res := MultiplyFractions(a, b)

	if res.Numerator == 623 && res.Denominator == 85 {
		return
	} else {
		t.Errorf("Didn't match expected result for MultiplyFractions(%d,%d) = %d ;", a, b, res)
	}
}

//TestMultiplyFractions8 inputs a=22/11 and b=1/3, and expectes o/p res=2/3.
func TestMultiplyFractions8(t *testing.T) {
	a := Fraction{Numerator: 22, Denominator: 11}
	b := Fraction{Numerator: 1, Denominator: 3}

	res := MultiplyFractions(a, b)

	if res.Numerator == 2 && res.Denominator == 3 {
		return
	} else {
		t.Errorf("Didn't match expected result for MultiplyFractions(%d,%d) = %d ;", a, b, res)
	}
}

//TestMultiplyFractions9 inputs a=55/23 and b=21/18, and expectes o/p res=1265/378.
func TestMultiplyFractions9(t *testing.T) {
	a := Fraction{Numerator: 55, Denominator: 21}
	b := Fraction{Numerator: 23, Denominator: 18}

	res := MultiplyFractions(a, b)

	if res.Numerator == 1265 && res.Denominator == 378 {
		return
	} else {
		t.Errorf("Didn't match expected result for MultiplyFractions(%d,%d) = %d ;", a, b, res)
	}
}

//TestMultiplyFractions10 inputs a=19/17 and b=13/47, and expectes o/p res=323/611.
func TestMultiplyFractions10(t *testing.T) {
	a := Fraction{Numerator: 19, Denominator: 13}
	b := Fraction{Numerator: 17, Denominator: 47}

	res := MultiplyFractions(a, b)

	if res.Numerator == 323 && res.Denominator == 611 {
		return
	} else {
		t.Errorf("Didn't match expected result for MultiplyFractions(%d,%d) = %d ;", a, b, res)
	}
}

//TestMultiplyFractionsTogether tests multiple inputs one after other for optimisation.
func TestMultiplyFractionsTogether(t *testing.T) {
	var tests = []struct {
		a      Fraction
		b      Fraction
		result Fraction
	}{
		{Fraction{7, 3}, Fraction{5, 10}, Fraction{21, 50}},
		{Fraction{89, 10}, Fraction{14, 17}, Fraction{623, 85}},
		{Fraction{55, 23}, Fraction{21, 18}, Fraction{1265, 378}},
		{Fraction{19, 17}, Fraction{13, 47}, Fraction{323, 611}},
		{Fraction{1, 2}, Fraction{1, 2}, Fraction{1, 2}},
		{Fraction{1, 2}, Fraction{1, 3}, Fraction{1, 6}},
		{Fraction{6, 10}, Fraction{4, 10}, Fraction{6, 25}},
		{Fraction{6, 5}, Fraction{4, 10}, Fraction{12, 25}},
		{Fraction{3, 6}, Fraction{5, 9}, Fraction{5, 18}},
	}

	for _, test := range tests {
		total := MultiplyFractions(test.a, test.b)
		if total.Numerator != test.result.Numerator && total.Denominator != test.result.Denominator {
			return
		} else {
			t.Errorf("Didn't match expected result for MultiplyFractions(%d,%d) = Got: %d , Expected: %d;", test.a, test.b, test.result, total)
		}
	}
}
