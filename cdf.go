package bigbinomial

// CDF returns a function that calculates the probability ρ Binomial
// Cumulative Distribution Function for n trials, for any value of
// k: 0 <= k <= n
//
// The binomial CDF calculates, for n independant binary trials each with
// success rate ρ, the probability that between 0 and k trials will be
// successful. So for 1000 flips of a fair coin, CDF(0.5, 1000, 500)
// calculates the probabilty that the number of heads will be less than or
// equal to 500.
//
// This package implements a function CDF(ρ, n) that returns a
// function cdf(k). This implementation of cdf "memoizes" (see:
// https://en.wikipedia.org/wiki/Memoization) its results for increasing
// consecutive values of k, so that calling cdf(k) for k = 0...n will be much
// faster than for k = n...0 (O(n) vs O(n^2) time, using constant memory).
//
// Returns an error when called with out of range values for ρ (which must be
// > 0.0 and < 1.0) or n (which must be > 1).
func CDF(ρ float64, n int64) (func(k int64) float64, error) {

	pmfFunc, err := PMF(ρ, n)
	if err != nil {
		return nil, err
	}

	lastK := int64(-1)
	lastVal := float64(0.0)

	return func(k int64) float64 {

		switch {

		case k == lastK+1:
			{
				lastK++
				lastVal += pmfFunc(k)
				return lastVal
			}

		case k < 0:
			return 0.0

		case k > n:
			return 1.0

		case k == lastK:
			return lastVal

		default:
			{
				var x int64
				lastK = k
				lastVal = 0.0
				for x = 0; x <= k; x++ {
					lastVal += pmfFunc(x)
				}
				return lastVal
			}
		}
	}, nil
}
