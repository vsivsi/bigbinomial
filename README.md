# BigBinomial
Golang package implementing math/big support for binomial distribution MDF and CDF and Pow for integer exponents

```golang
import (
   "github.com/vsivsi/bigbinomial"
   "math/big"
   "math"
)

// BigBinomial implements, uses and exports a bigFloat implementation
// of math.Pow(), but restricted to integer exponents

// Defaults to double precision equivalent
math.Pow(10.0, 250)        // 1.0000000000000004e+250
BigBinomial.Pow(10.0, 250) // 1.0000000000000004e+250
BigBinomial.Pow(10.0, 250).Cmp(big.NewFloat(math.Pow(10.0, 250))) // 0 --> equal

// More accurate for higher precision inputs
bigten := (&big.Float{}).SetPrec(300).SetInt64(10)
BigBinomial.Pow(bigten, 250)  // 1e+250

// Handles much larger exponents
math.Pow(10.0, 2500)         // +Inf
BigBinomial.Pow(10.0, 2500)  // 1.0000000000000052e+2500

// And much smaller ones
math.Pow(10.0, -2500)         // 0
BigBinomial.Pow(10.0, -2500)  // 1.0000000000002682e-2500

```
