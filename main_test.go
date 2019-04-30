package bigbinomial

import (
	"github.com/stretchr/testify/assert"
	"github.com/vsivsi/bigfloat"
	"math"
	"math/big"
	"testing"
)

// tolerance used for all epsilon tests
const ε = 1.0e-15

// assert.ComparisonAssertionFunction compatible wrapper for assert.InEpsilon
func epsilon(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	return assert.InEpsilon(t, expected, actual, ε)
}

// tolerance used for all delta tests
const δ = 1.0e-14

// assert.ComparisonAssertionFunction compatible wrapper for assert.InDelta
func delta(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	return assert.InDelta(t, expected, actual, δ)
}

type testCmpSlice []struct {
	name   string
	actual interface{}
	expect interface{}
	assert assert.ComparisonAssertionFunc
}

func runTests(t *testing.T, tests testCmpSlice) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assert(t, tt.actual, tt.expect)
		})
	}
}

// TestPow tests the bigfloat.Pow function
func TestPow(t *testing.T) {

	// pow Calculates X^n for a bigFloat X for any int64 n
	pow := func(X *big.Float, n int64) *big.Float {
		return bigfloat.Pow(X, big.NewFloat(float64(n)))
	}

	toFloat64 := func(x *big.Float) float64 {
		val, _ := x.Float64()
		return val
	}

	zero := big.NewFloat(0.0)
	negZero := big.NewFloat(0.0).Neg(zero)
	one := big.NewFloat(1.0)
	negOne := big.NewFloat(-1.0)
	two := big.NewFloat(2.0)
	ten := big.NewFloat(10.0)
	inf := big.NewFloat(0.0).SetInf(false)
	negInf := big.NewFloat(0.0).SetInf(true)

	tests := testCmpSlice{
		{"pow(inf, 0)", pow(inf, 0), one, assert.Equal},
		{"pow(inf, 1)", pow(inf, 1), inf, assert.Equal},
		{"pow(inf, 2)", pow(inf, 2), inf, assert.Equal},
		{"pow(inf, -1)", pow(inf, -1), zero, assert.Equal},
		{"pow(inf, -2)", pow(inf, -2), zero, assert.Equal},
		{"pow(negInf, 0)", pow(negInf, 0), one, assert.Equal},
		{"pow(negInf, 1)", pow(negInf, 1), negInf, assert.Equal},
		{"pow(negInf, 2)", pow(negInf, 2), inf, assert.Equal},
		{"pow(negInf, -1)", pow(negInf, -1), negZero, assert.Equal},
		{"pow(negInf, -2)", pow(negInf, -2), zero, assert.Equal},
		{"pow(zero, 0)", pow(zero, 0), one, assert.Equal},
		{"pow(zero, 1)", pow(zero, 1), zero, assert.Equal},
		{"pow(zero, 2)", pow(zero, 2), zero, assert.Equal},
		{"pow(zero, -1)", pow(zero, -1), inf, assert.Equal},
		{"pow(zero, -2)", pow(zero, -2), inf, assert.Equal},
		{"pow(negZero, 0)", pow(negZero, 0), one, assert.Equal},
		{"pow(negZero, 1)", pow(negZero, 1), negZero, assert.Equal},
		{"pow(negZero, 2)", pow(negZero, 2), zero, assert.Equal},
		{"pow(negZero, -1)", pow(negZero, -1), negInf, assert.Equal},
		{"pow(negZero, -2)", pow(negZero, -2), inf, assert.Equal},
		{"pow(one, 0)", pow(one, 0), one, assert.Equal},
		{"pow(one, 1)", pow(one, 1), one, assert.Equal},
		{"pow(one, 2)", pow(one, 2), one, assert.Equal},
		{"pow(one, -1)", pow(one, -1), one, assert.Equal},
		{"pow(one, -2)", pow(one, -2), one, assert.Equal},
		{"pow(negOne, 0)", pow(negOne, 0), one, assert.Equal},
		{"pow(negOne, 1)", pow(negOne, 1), negOne, assert.Equal},
		{"pow(negOne, 2)", pow(negOne, 2), one, assert.Equal},
		{"pow(negOne, 1)", pow(negOne, 1), negOne, assert.Equal},
		{"pow(negOne, 2)", pow(negOne, 2), one, assert.Equal},
		{"pow(two, 0)", pow(two, 0), one, assert.Equal},
		{"pow(two, 1)", pow(two, 1), two, assert.Equal},
		{"pow(two, 2)", pow(two, 2), big.NewFloat(0).Mul(two, two), assert.Equal},
		{"pow(two, -1)", pow(two, -1), big.NewFloat(0).Quo(one, two), assert.Equal},
		{"pow(two, -2)", pow(two, -2), big.NewFloat(0).Quo(one, big.NewFloat(0).Mul(two, two)), assert.Equal},
		// Epsilon tests
		{"pow(ten, 2)", toFloat64(pow(ten, 2)), math.Pow(10.0, 2), epsilon},
		{"pow(ten, 5)", toFloat64(pow(ten, 5)), math.Pow(10.0, 5), epsilon},
		{"pow(ten, 25)", toFloat64(pow(ten, 25)), math.Pow(10.0, 25), epsilon},
		{"pow(ten, 250)", toFloat64(pow(ten, 250)), math.Pow(10.0, 250), epsilon},
		{"pow(ten, -2)", toFloat64(pow(ten, -2)), math.Pow(10.0, -2), epsilon},
		{"pow(ten, -5)", toFloat64(pow(ten, -5)), math.Pow(10.0, -5), epsilon},
		{"pow(ten, -25)", toFloat64(pow(ten, -25)), math.Pow(10.0, -25), epsilon},
		{"pow(ten, -250)", toFloat64(pow(ten, -250)), math.Pow(10.0, -250), epsilon},
	}

	runTests(t, tests)

}

// TestPMF implements unit tests for the bigbinomial.PMF function
func TestPMF(t *testing.T) {

	binomialPMF := func(ρ float64, n int64) func(k int64) float64 {
		return func(k int64) float64 {
			res := math.Pow(ρ, float64(k)) * math.Pow(1-ρ, float64(n-k))
			res *= math.Gamma(float64(n+1)) / (math.Gamma(float64(k+1)) * math.Gamma(float64(n-k+1)))
			return res
		}
	}

	pmfErr := func(ρ float64, n int64) float64 {
		floatPMF := binomialPMF(ρ, n)
		bigPMF, _ := PMF(ρ, n)
		err := 0.0
		for x := int64(0); x <= n; x++ {
			err += math.Abs(bigPMF(x) - floatPMF(x))
		}
		return err
	}

	tests := testCmpSlice{
		{"PMF(0.5, 3)", pmfErr(0.5, 3), 0.0, delta},
		{"PMF(0.5, 30)", pmfErr(0.5, 30), 0.0, delta},
		{"PMF(0.5, 150)", pmfErr(0.5, 150), 0.0, delta},
		{"PMF(0.05, 3)", pmfErr(0.05, 3), 0.0, delta},
		{"PMF(0.05, 30)", pmfErr(0.05, 30), 0.0, delta},
		{"PMF(0.05, 150)", pmfErr(0.05, 150), 0.0, delta},
		{"PMF(0.005, 3)", pmfErr(0.005, 3), 0.0, delta},
		{"PMF(0.005, 30)", pmfErr(0.005, 30), 0.0, delta},
		{"PMF(0.005, 150)", pmfErr(0.005, 150), 0.0, delta},
	}
	runTests(t, tests)
}

// TestCDF implements unit tests for the bigbinomial.CDF function
func TestCDF(t *testing.T) {

	ρ := 0.5
	n := int64(200)

	floatCDF, _ := CDF(ρ, n)
	floatPMF, _ := PMF(ρ, n)

	tests := testCmpSlice{
		{"CDF(k) != CDF(k-1) + PMF(k)", floatCDF(n / 2), floatCDF(n/2-1) + floatPMF(n/2), assert.Equal},
		{"CDF(k) != CDF(k+1) - PMF(k+1)", floatCDF(n / 2), floatCDF(n/2+1) - floatPMF(n/2+1), assert.Equal},
		{"CDF(k) != CDF(k)", floatCDF(n / 2), floatCDF(n / 2), assert.Equal},
		{"PMF(-1) != 0.0", floatPMF(-1), 0.0, assert.Equal},
		{"PMF(n+1) != 0.0", floatPMF(n + 1), 0.0, assert.Equal},
		{"CDF(-1) != 0.0", floatCDF(-1), 0.0, assert.Equal},
		{"CDF(n+1) != 1.0", floatCDF(n + 1), 1.0, assert.Equal},
		{"CDF(n) != 1.0", floatCDF(n), 1.0, assert.Equal},
	}

	runTests(t, tests)

	errTests := []struct {
		name string
		ρ    float64
		n    int64
	}{
		{"CDF(ρ, n) for ρ < 0.0", -1.0, n},
		{"CDF(ρ, n) for ρ > 1.0", 2.0, n},
		{"CDF(ρ, n) for n <= 0.0", ρ, 0},
	}

	for _, tt := range errTests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CDF(tt.ρ, tt.n)
			assert.Error(t, err)
		})
	}
}
