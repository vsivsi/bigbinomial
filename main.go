package BigBinomial

import (
	"fmt"
	"math/big"
)

// Exp Calculates X^n for a bigFloat X for any int64 n
func Exp(X *big.Float, n int64) *big.Float {

	x := (&big.Float{}).Copy(X)

	if n < 0 {
		x = x.Quo(big.NewFloat(1.0), x)
		n = -n
	}

	if n == 0 {
		return big.NewFloat(1.0)
	}

	y := big.NewFloat(1.0)

	for n > 1 {
		if n%2 == 0 {
			x = x.Mul(x, x)
			n = n / 2
		} else {
			y = y.Mul(y, x)
			x = x.Mul(x, x)
			n = (n - 1) / 2
		}
	}

	x = x.Mul(x, y)

	return x
}

// PMF returns a function that calculates the probability ρ Binomial
// Probability Mass Function for n trials, for any value of
// k: 0 <= k <= n
func PMF(ρ float64, n int64) func(k int64) float64 {

	if ρ < 0.0 || ρ > 1.0 || n < 1 {
		fmt.Println("Parameter out of range in Big_Binomial_PMF")
		panic(0)
	}

	return func(k int64) float64 {

		b := (&big.Int{}).Binomial(n, k)
		z := (&big.Float{}).SetPrec(256).SetInt(b)
		bigP := big.NewFloat(ρ).SetPrec(256)
		i1 := Exp(bigP, k)
		i2 := Exp((&big.Float{}).Sub(big.NewFloat(1.0), bigP), n-k)
		i := (&big.Float{}).Mul(i1, i2)
		z = z.Mul(z, i)

		// Discarding accuracy for now...
		retval, _ := z.Float64()
		return retval
	}
}
