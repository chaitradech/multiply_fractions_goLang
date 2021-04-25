package fractions


// To be treated as a class, to represent fractions
type Fraction struct {
	Numerator, Denominator int
}

//Function to convert each of the fraction on to it's simplified form
func Reduce(input Fraction) (res1, res2 int) {
	gcd := ComputeGCD(input.Numerator,input.Denominator)
	res1 = input.Numerator/ gcd
	res2 = input.Denominator/ gcd
	return
}

//Function to calculate the GCD 
func ComputeGCD( numerator ,denominator int) int {
	if (denominator == 0) {
		return numerator
	}
	return ComputeGCD(denominator, numerator % denominator)
}

//Function to Multiply the simplified inputs
func MultiplyFractions( input1 Fraction, input2 Fraction) Fraction {
	output1 := input1.Numerator *  input2.Numerator
	output2 := input1.Denominator * input2.Denominator
	
	ans := Fraction{Numerator: output1 , Denominator: output2}
	ans.Numerator, ans.Denominator = Reduce(ans)
	return ans 
}

