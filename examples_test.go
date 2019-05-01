package bigbinomial_test

import "math"
import "fmt"
import "github.com/vsivsi/bigbinomial"

func Example() {

	binomialPMF := func(ρ float64, n int64) func(k int64) float64 {
		return func(k int64) float64 {
			p := math.Pow(ρ, float64(k)) * math.Pow(1-ρ, float64(n-k))
			p *= math.Gamma(float64(n+1)) /
				(math.Gamma(float64(k+1)) * math.Gamma(float64(n-k+1)))
			return p
		}
	}

	// If you flip a fair coin 100 times (ρ=0.5, n=100), the probability of flipping
	// heads exactly 50 times (k=25) is PMF(ρ, n, k):

	pmf := binomialPMF(0.5, 100) // Return a function to calculate PMF(0.5, 100, k)
	prob := pmf(50)              // prob = 0.07958923738717867, about 8%
	fmt.Println(prob)

	// But what if we want the answer for 500 heads out of 1000 flips?
	// The float64 implementation breaks for large values of n

	pmf = binomialPMF(0.5, 1000)
	prob = pmf(500) // prob == NaN  (Failure!)
	fmt.Println(prob)

	// Output:
	// 0.07958923738717867
	// NaN
}

func ExamplePMF() {
	pmf, _ := bigbinomial.PMF(0.5, 1000)
	prob := pmf(500)
	fmt.Println(prob)

	// Output:
	// 0.0252250181783608
}

func ExampleCDF() {
	cdf, _ := bigbinomial.CDF(0.5, 1000)
	prob := cdf(500)
	fmt.Println(prob)

	// Output:
	// 0.5126125090891803
}
