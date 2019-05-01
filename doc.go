/*
Package bigbinomial implements binomial distribution PMF and CDF functions with math/big support

Example Usage

The binomial PMF calculates, for n independant binary trials each with success rate ρ, the probability that k out of n will be successful.
For example: if you flip a fair coin 50 times (ρ=0.5, n=50), the probability of flipping heads exactly 25 times (k=25) is PMF(ρ, n, k).
This package implements a function `PMF(ρ, n)` that returns a function `pmf(k)`.

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

		// So to calculate the answer to the example above for ρ=0.5 and n=50
		pmf := binomialPMF(0.5, 50) // Return a function to calculate PMF(k)
		prob := pmf(50)             // prob == 0.07958923738717867, or about 8%

		// But what if we want the answer for 500 heads out of 1000 coin flips?
		// The float64 implementation does not work for large values of n
		pmf = binomialPMF(0.5, 1000)
		prob = pmf(500)  // prob == NaN (!)

		// bigbinomial uses the golang math/big library to remove this limitation
		pmf, _ = bigbinomial.PMF(0.5, 50)
		prob = pmf(50)   // prob == 0.07958923738717877

		pmf, _ = bigbinomial.PMF(0.5, 1000)
		prob = pmf(500)  // prob == 0.0252250181783608

The binomial CDF calculates, for n independant binary trials each with success rate ρ, CDF(ρ, n, k) is the probability that between 0 and k trials will be successful.
So for 1000 flips of a fair coin, cdf(500) calculates the probabilty that the number of heads will be less than or equal to 500.
This package implements a function `CDF(ρ, n)` that returns a function `cdf(k)`. This implementation of cdf ["memoizes"](https://en.wikipedia.org/wiki/Memoization) its results for increasing consecutive values of k, so that calling `cdf(k)` for k = 0...n will be much faster than for k = n...0 (O(n) vs O(n^2) time, using constant memory).

		import "github.com/vsivsi/bigbinomial"
		cdf, _ := bigbinomial.CDF(0.5, 1000)
		prob := cdf(500)   // prob == 0.5126125090891803
*/
package bigbinomial
