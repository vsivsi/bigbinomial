# BigBinomial
Golang package implementing binomial distribution PMF and CDF functions using golang's math/big library to allow larger values of n. Also exports a big.Float Pow method for int64 exponents.

## Binomial Distribution Probability Mass Function (PMF)  

```golang
import (
	"math"
	"github.com/vsivsi/bigbinomial"
)

// It is straightforward to implement the Binomial PMF using math library functions:

func binomialPMF(ρ float64, n int64) func(k int64) float64 {
	return func(k int64) float64 {
		p := math.Pow(ρ, float64(k)) * math.Pow(1-ρ, float64(n-k))
		p *= math.Gamma(float64(n+1)) / 
			(math.Gamma(float64(k+1)) * math.Gamma(float64(n-k+1)))
		return p
	}
}

// This works
pmf := binomialPMF(0.5, 100)
prob := pmf(50) // prob == 0.07958923738717867

// But a float64 implementation does not work for large values of n
pmf = binomialPMF(0.5, 1000)
prob = pmf(500) // prob == NaN (!)

// BigBinomial uses the golang math/big library to remove this limitation
pmf, _ = BigBinomial.PMF(0.5, 100)
prob = pmf(50)  // prob == 0.07958923738717877

pmf, _ = BigBinomial.PMF(0.5, 1000)
prob = pmf(500)  // prob == 0.0252250181783608
```

## Binomial Distribution Cumulative Distribution Function (CDF)

```golang
import (
	"github.com/vsivsi/bigbinomial"
)

cdf, _ := BigBinomial.CDF(0.5, 1000)
prob := cdf(500)  // prob == 0.5126125090891803
```

## Pow method for big.Float type, with int64 exponents

A big.Float implementation of Pow was required for the above, and since the math/big library doesn't currently provide one, it is also exported. It is currently only defined for integer exponents.

```golang
import (
	"github.com/vsivsi/bigbinomial"
	"math/big"
	"math"
)

// BigBinomial also implements, uses and exports a bigFloat implementation
// of math.Pow(), but restricted to integer exponents

// Defaults to double precision equivalent
val := math.Pow(10.0, 250)           //     val == 1.0000000000000004e+250
bigVal := BigBinomial.Pow(10.0, 250) //  bigVal == 1.0000000000000004e+250
cmp := BigBinomial.Pow(10.0, 250).Cmp(big.NewFloat(math.Pow(10.0, 250))) // cmp == 0 --> equal

// More accurate for higher precision inputs
bigten := (&big.Float{}).SetPrec(300).SetInt64(10)
bigVal = BigBinomial.Pow(bigten, 250) //  bigVal == 1e+250

// Handles much larger exponents
val = math.Pow(10.0, 2500)            // val == +Inf
bigVal = BigBinomial.Pow(10.0, 2500)  // bigVal == 1.0000000000000052e+2500

// And much smaller ones
val = math.Pow(10.0, -2500)            // val == 0
bigVal = BigBinomial.Pow(10.0, -2500)  // bigVal == 1.0000000000002682e-2500
```
