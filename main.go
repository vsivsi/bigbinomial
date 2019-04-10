package BigBinomial

import (
	"fmt"
	"math/big"
)

// Pow Calculates X^n for a bigFloat X for any int64 n
func Pow(X *big.Float, n int64) *big.Float {

	x := (&big.Float{}).Copy(X)
	y := (&big.Float{}).SetPrec(x.Prec()).SetUint64(1)

	if n == 0 {
		// X^0 == 1.0
		// including when X == 0.0, even though that case may be considered indeterminate.
		// See: https://github.com/golang/go/issues/7583#issuecomment-66092687
		return y
	}

	if n < 0 {
		// X^-n == (1/X)^n
		x = x.Quo(y, x)
		n = -n
	}

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
		fmt.Println("Parameter out of range in BigBinomial.PMF")
		panic(0)
	}

	return func(k int64) float64 {

		if k < 0 || k > n {
			fmt.Println("k out of range in PMF function call")
			panic(0)
		}

		b := (&big.Int{}).Binomial(n, k)
		bits := uint(b.BitLen()) + 64
		z := (&big.Float{}).SetPrec(bits).SetInt(b)
		bigP := big.NewFloat(ρ).SetPrec(bits)

		i1 := Pow(bigP, k)
		i2 := Pow((&big.Float{}).Sub(big.NewFloat(1.0), bigP), n-k)
		i := (&big.Float{}).Mul(i1, i2)
		z = z.Mul(z, i)

		// Discarding accuracy for now...
		retval, _ := z.Float64()
		return retval
	}
}

// CDF returns a function that calculates the probability ρ Binomial
// Cumulative Distribution Function for n trials, for any value of
// k: 0 <= k <= n
func CDF(ρ float64, n int64) func(k int64) float64 {

	if ρ < 0.0 || ρ > 1.0 || n < 1 {
		fmt.Println("Parameter out of range in BigBinomial.CDF")
		panic(0)
	}

	pmfFunc := PMF(ρ, n)
	lastK := int64(-1)
	lastVal := float64(0.0)

	return func(k int64) float64 {

		if k < 0 || k > n {
			fmt.Println("k out of range in CDF function call")
			panic(0)
		}

		if k == lastK {
			return lastVal
		}

		if k == lastK+1 {
			lastK++
			lastVal += pmfFunc(k)
			return lastVal
		}

		var x int64
		lastK = k
		lastVal = 0.0
		for x = 0; x <= k; x++ {
			lastVal += pmfFunc(x)
		}
		return lastVal
	}
}
