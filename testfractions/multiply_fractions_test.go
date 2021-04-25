package testfractions

import (
	"testing"
	. "fractions/multiply_fractions"
)

//Testing
func TestComputeGCD(t *testing.T) {
	got := ComputeGCD(6,10)
	if got != 2 {
		t.Errorf("computeGCD(6,10) = %d ; want 2", got)
	}
}

