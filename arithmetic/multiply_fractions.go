package fractions

// Fraction signifies a fraction value.
type Fraction struct {
	Numerator   int64
	Denominator int64
}

// reduce simplifies each of the fraction value on to it's reduce form.
func reduce(input Fraction) (res1, res2 int64) {
	gcd := computeGCD(input.Numerator, input.Denominator)
	res1 = input.Numerator / gcd
	res2 = input.Denominator / gcd
	return
}

// computeGCD returns the GCD using the recursion technique.
func computeGCD(numerator, denominator int64) int64 {
	if denominator == 0 {
		return numerator
	}
	return computeGCD(denominator, numerator%denominator)
}

// MultiplyFractions performs multiplication of two fraction values.
func MultiplyFractions(input1 Fraction, input2 Fraction) Fraction {
	output1 := input1.Numerator * input2.Numerator
	output2 := input1.Denominator * input2.Denominator

	ans := Fraction{Numerator: output1, Denominator: output2}
	ans.Numerator, ans.Denominator = reduce(ans)
	return ans
}
