package testfractions

import (
	"testing"
	. "fractions/multiply_fractions"
)

//Testing ComputeGCD function

func TestComputeGCD(t *testing.T) {
	got := ComputeGCD(6,10)
	assertEqual(t, got, 2)
  }

  func assertEqual(t *testing.T, got int, expected int) {
	if got == expected {
	  return
	  }  else {
		t.Errorf("computeGCD(6,10) = %d ; want 2", expected)
	  }
  }

  //Testing with inputs a=1/2 and b=1/3, expected o/p res=1/6
func TestMultiplyFractions1(t *testing.T) {
	a := Fraction{Numerator: 1, Denominator: 2}
	b := Fraction{Numerator: 1,Denominator: 3}

	
	res := MultiplyFractions(a, b)

	if (res.Numerator == 1 && res.Denominator ==6) {
		return
	}else {
		t.Errorf("Didn't match expected result for MultiplyFractions(%d,%d) = %d ;", a, b, res)
	  }
}

//Testing with inputs a=6/10 and b=4/10, expected o/p res=6/25
  func TestMultiplyFractions2(t *testing.T) {
	a := Fraction{Numerator: 6, Denominator: 10}
	b := Fraction{Numerator: 4,Denominator: 10}

	
	res := MultiplyFractions(a, b)

	if (res.Numerator == 6 && res.Denominator ==25) {
		return
	}else {
		t.Errorf("Didn't match expected result for MultiplyFractions(%d,%d) = %d ;", a, b, res)
	  }
}

//Testing with inputs a=6/5 and b=4/10, expected o/p res=12/25
func TestMultiplyFractions3(t *testing.T) {
	a := Fraction{Numerator: 6, Denominator: 5}
	b := Fraction{Numerator: 4,Denominator: 10}

	
	res := MultiplyFractions(a, b)

	if (res.Numerator == 12 && res.Denominator ==25) {
		return
	}else {
		t.Errorf("Didn't match expected result for MultiplyFractions(%d,%d) = %d ;", a, b, res)
	  }
}

//Testing with inputs a=12/10 and b=10/12, expected o/p : 1/1
func TestMultiplyFractions4(t *testing.T) {
	a := Fraction{Numerator: 12, Denominator: 10}
	b := Fraction{Numerator: 10,Denominator: 12}

	
	res := MultiplyFractions(a, b)

	if (res.Numerator == 1 && res.Denominator ==1) {
		return
	}else {
		t.Errorf("Didn't match expected result for MultiplyFractions(%d,%d) = %d ;", a, b, res)
	  }
}
//Testing with inputs a=3/6 and b=5/9, expected o/p res=5/18
func TestMultiplyFractions5(t *testing.T) {
	a := Fraction{Numerator: 3, Denominator: 6}
	b := Fraction{Numerator: 5,Denominator: 9}

	
	res := MultiplyFractions(a, b)

	if (res.Numerator == 5 && res.Denominator ==18) {
		return
	}else {
		t.Errorf("Didn't match expected result for MultiplyFractions(%d,%d) = %d ;", a, b, res)
	  }
}
//Testing with inputs a=7/3 and b=5/10, expected o/p res=21/50
func TestMultiplyFractions6(t *testing.T) {
	a := Fraction{Numerator: 7, Denominator: 5}
	b := Fraction{Numerator: 3,Denominator: 10}

	
	res := MultiplyFractions(a, b)

	if (res.Numerator == 21 && res.Denominator ==50) {
		return
	}else {
		t.Errorf("Didn't match expected result for MultiplyFractions(%d,%d) = %d ;", a, b, res)
	  }
}
//Testing with inputs a=89/10 and b=14/17, expected o/p res=623/85
func TestMultiplyFractions7(t *testing.T) {
	a := Fraction{Numerator: 89, Denominator: 10}
	b := Fraction{Numerator: 14,Denominator: 17}

	
	res := MultiplyFractions(a, b)

	if (res.Numerator == 623 && res.Denominator ==85) {
		return
	}else {
		t.Errorf("Didn't match expected result for MultiplyFractions(%d,%d) = %d ;", a, b, res)
	  }
}
//Testing with inputs a=22/11 and b=1/3, expected o/p res=2/3
func TestMultiplyFractions8(t *testing.T) {
	a := Fraction{Numerator: 22, Denominator: 11}
	b := Fraction{Numerator: 1,Denominator: 3}

	
	res := MultiplyFractions(a, b)

	if (res.Numerator == 2 && res.Denominator ==3) {
		return
	}else {
		t.Errorf("Didn't match expected result for MultiplyFractions(%d,%d) = %d ;", a, b, res)
	  }
}
//Testing with inputs a=55/23 and b=21/18, expected o/p res=1265/378
func TestMultiplyFractions9(t *testing.T) {
	a := Fraction{Numerator: 55, Denominator: 21}
	b := Fraction{Numerator: 23,Denominator: 18}

	
	res := MultiplyFractions(a, b)

	if (res.Numerator == 1265 && res.Denominator ==378) {
		return
	}else {
		t.Errorf("Didn't match expected result for MultiplyFractions(%d,%d) = %d ;", a, b, res)
	  }
}
//Testing with inputs a=19/17 and b=13/47, expected o/p res=323/611
func TestMultiplyFractions10(t *testing.T) {
	a := Fraction{Numerator: 19, Denominator: 13}
	b := Fraction{Numerator: 17,Denominator: 47}

	
	res := MultiplyFractions(a, b)

	if (res.Numerator == 323 && res.Denominator ==611) {
		return
	}else {
		t.Errorf("Didn't match expected result for MultiplyFractions(%d,%d) = %d ;", a, b, res)
	  }
}
