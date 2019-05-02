package bigbinomial

import (
	"fmt"
	"math/big"

	"github.com/vsivsi/bigfloat" // was: "github.com/ALTree/bigfloat"
)

// PMF returns a function that calculates the probability ρ Binomial
// Probability Mass Function for n trials, for any value of
// k: 0 <= k <= n
//
// The binomial PMF calculates, for n independant binary trials each with
// success rate ρ, the probability that k out of n will be successful. For
// example: if you flip a fair coin 100 times (ρ=0.5, n=100), the probability
// of flipping heads exactly 50 times (k=50) is PMF(ρ, n, k).
//
// This package implements a function PMF(ρ, n) that returns a function pmf(k).
//
// Returns an error when called with out of range values for ρ (which must be
// > 0.0 and < 1.0) or n (which must be > 1).
func PMF(ρ float64, n int64) (func(k int64) float64, error) {

	if ρ < 0.0 || ρ > 1.0 {
		return nil, fmt.Errorf("Parameter ρ must be between 0.0 and 1.0, ρ = %g", ρ)
	}

	if n <= 0 {
		return nil, fmt.Errorf("Parameter n must be greater than 0, n = %d", n)
	}

	return func(k int64) float64 {

		if k < 0 || k > n {
			return 0.0
		}

		b := (&big.Int{}).Binomial(n, k)
		bits := uint(b.BitLen()) + 64
		z := (&big.Float{}).SetPrec(bits).SetInt(b)
		bigP := big.NewFloat(ρ).SetPrec(bits)

		i1 := bigfloat.Pow(bigP, big.NewFloat(float64(k)))
		i2 := bigfloat.Pow((&big.Float{}).SetPrec(bits).Sub(big.NewFloat(1.0), bigP), big.NewFloat(float64(n-k)))
		i := i1.Mul(i1, i2)
		z = z.Mul(z, i)

		// Discarding accuracy for now...
		retval, _ := z.Float64()
		return retval
	}, nil
}
