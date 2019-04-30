package bigbinomial

import (
	"github.com/stretchr/testify/assert"
	"github.com/vsivsi/bigfloat"
	"math"
	"math/big"
	"testing"
)

// TestPow tests the bigfloat.Pow function
func TestPow(t *testing.T) {

	Zero := big.NewFloat(0.0)
	NegZero := big.NewFloat(0.0).Neg(Zero)
	One := big.NewFloat(1.0)
	NegOne := big.NewFloat(-1.0)
	Two := big.NewFloat(2.0)
	Ten := big.NewFloat(10.0)
	Inf := big.NewFloat(0.0).SetInf(false)
	NegInf := big.NewFloat(0.0).SetInf(true)

	// Pow Calculates X^n for a bigFloat X for any int64 n
	Pow := func(X *big.Float, n int64) *big.Float {
		return bigfloat.Pow(X, big.NewFloat(float64(n)))
	}

	Epsilon := func(tol float64) assert.ComparisonAssertionFunc {
		return func(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool {
			return assert.InEpsilon(t, expected, actual, tol)
		}
	}(1.0e-15)

	toFloat64 := func(x *big.Float) float64 {
		val, _ := x.Float64()
		return val
	}

	tests := []struct {
		name   string
		actual interface{}
		expect interface{}
		assert assert.ComparisonAssertionFunc
	}{
		{"Pow(Inf, 0)", Pow(Inf, 0), One, assert.Equal},
		{"Pow(Inf, 1)", Pow(Inf, 1), Inf, assert.Equal},
		{"Pow(Inf, 2)", Pow(Inf, 2), Inf, assert.Equal},
		{"Pow(Inf, -1)", Pow(Inf, -1), Zero, assert.Equal},
		{"Pow(Inf, -2)", Pow(Inf, -2), Zero, assert.Equal},
		{"Pow(NegInf, 0)", Pow(NegInf, 0), One, assert.Equal},
		{"Pow(NegInf, 1)", Pow(NegInf, 1), NegInf, assert.Equal},
		{"Pow(NegInf, 2)", Pow(NegInf, 2), Inf, assert.Equal},
		{"Pow(NegInf, -1)", Pow(NegInf, -1), NegZero, assert.Equal},
		{"Pow(NegInf, -2)", Pow(NegInf, -2), Zero, assert.Equal},
		{"Pow(Zero, 0)", Pow(Zero, 0), One, assert.Equal},
		{"Pow(Zero, 1)", Pow(Zero, 1), Zero, assert.Equal},
		{"Pow(Zero, 2)", Pow(Zero, 2), Zero, assert.Equal},
		{"Pow(Zero, -1)", Pow(Zero, -1), Inf, assert.Equal},
		{"Pow(Zero, -2)", Pow(Zero, -2), Inf, assert.Equal},
		{"Pow(NegZero, 0)", Pow(NegZero, 0), One, assert.Equal},
		{"Pow(NegZero, 1)", Pow(NegZero, 1), NegZero, assert.Equal},
		{"Pow(NegZero, 2)", Pow(NegZero, 2), Zero, assert.Equal},
		{"Pow(NegZero, -1)", Pow(NegZero, -1), NegInf, assert.Equal},
		{"Pow(NegZero, -2)", Pow(NegZero, -2), Inf, assert.Equal},
		{"Pow(One, 0)", Pow(One, 0), One, assert.Equal},
		{"Pow(One, 1)", Pow(One, 1), One, assert.Equal},
		{"Pow(One, 2)", Pow(One, 2), One, assert.Equal},
		{"Pow(One, -1)", Pow(One, -1), One, assert.Equal},
		{"Pow(One, -2)", Pow(One, -2), One, assert.Equal},
		{"Pow(NegOne, 0)", Pow(NegOne, 0), One, assert.Equal},
		{"Pow(NegOne, 1)", Pow(NegOne, 1), NegOne, assert.Equal},
		{"Pow(NegOne, 2)", Pow(NegOne, 2), One, assert.Equal},
		{"Pow(NegOne, 1)", Pow(NegOne, 1), NegOne, assert.Equal},
		{"Pow(NegOne, 2)", Pow(NegOne, 2), One, assert.Equal},
		{"Pow(Two, 0)", Pow(Two, 0), One, assert.Equal},
		{"Pow(Two, 1)", Pow(Two, 1), Two, assert.Equal},
		{"Pow(Two, 2)", Pow(Two, 2), big.NewFloat(0).Mul(Two, Two), assert.Equal},
		{"Pow(Two, -1)", Pow(Two, -1), big.NewFloat(0).Quo(One, Two), assert.Equal},
		{"Pow(Two, -2)", Pow(Two, -2), big.NewFloat(0).Quo(One, big.NewFloat(0).Mul(Two, Two)), assert.Equal},
		// Epsilon tests
		{"Pow(Ten, 2)", toFloat64(Pow(Ten, 2)), math.Pow(10.0, 2), Epsilon},
		{"Pow(Ten, 5)", toFloat64(Pow(Ten, 5)), math.Pow(10.0, 5), Epsilon},
		{"Pow(Ten, 25)", toFloat64(Pow(Ten, 25)), math.Pow(10.0, 25), Epsilon},
		{"Pow(Ten, 250)", toFloat64(Pow(Ten, 250)), math.Pow(10.0, 250), Epsilon},
		{"Pow(Ten, -2)", toFloat64(Pow(Ten, -2)), math.Pow(10.0, -2), Epsilon},
		{"Pow(Ten, -5)", toFloat64(Pow(Ten, -5)), math.Pow(10.0, -5), Epsilon},
		{"Pow(Ten, -25)", toFloat64(Pow(Ten, -25)), math.Pow(10.0, -25), Epsilon},
		{"Pow(Ten, -250)", toFloat64(Pow(Ten, -250)), math.Pow(10.0, -250), Epsilon},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assert(t, tt.actual, tt.expect)
		})
	}
}

// TestPMF implements unit tests for the bigbinomial.PMF function
func TestPMF(t *testing.T) {

	binomialPMF := func(ρ float64, n int64) func(k int64) float64 {
		return func(k int64) float64 {
			p := math.Pow(ρ, float64(k)) * math.Pow(1-ρ, float64(n-k))
			p *= math.Gamma(float64(n+1)) / (math.Gamma(float64(k+1)) * math.Gamma(float64(n-k+1)))
			return p
		}
	}

	assertPMFTol := func(p float64, n int64, tol float64, t *testing.T) {
		t.Helper()
		floatPMF := binomialPMF(p, n)
		bigPMF, _ := PMF(p, n)
		err := 0.0
		for x := int64(0); x <= n; x++ {
			err += math.Abs(bigPMF(x) - floatPMF(x))
		}
		if err > tol {
			t.Fatal("Error:", err, "Tolerance:", tol)
		}
	}

	Tol := 1.0e-15

	t.Run("PMF(0.5, 2)", func(t *testing.T) {
		assertPMFTol(0.5, 2, Tol, t)
	})
	t.Run("PMF(0.5, 20)", func(t *testing.T) {
		assertPMFTol(0.5, 20, Tol, t)
	})
	t.Run("PMF(0.5, 200)", func(t *testing.T) {
		assertPMFTol(0.5, 200, Tol, t)
	})

	t.Run("PMF(0.15, 2)", func(t *testing.T) {
		assertPMFTol(0.15, 2, Tol, t)
	})
	t.Run("PMF(0.15, 20)", func(t *testing.T) {
		assertPMFTol(0.15, 20, Tol, t)
	})
	t.Run("PMF(0.15, 200)", func(t *testing.T) {
		assertPMFTol(0.15, 200, Tol, t)
	})

	t.Run("PMF(0.5, 2)", func(t *testing.T) {
		assertPMFTol(0.5, 2, Tol, t)
	})
	t.Run("PMF(0.5, 20)", func(t *testing.T) {
		assertPMFTol(0.5, 20, Tol, t)
	})
	t.Run("PMF(0.5, 200)", func(t *testing.T) {
		assertPMFTol(0.5, 200, Tol, t)
	})

	t.Run("PMF(0.5, 2)", func(t *testing.T) {
		assertPMFTol(0.5, 2, Tol, t)
	})
	t.Run("PMF(0.5, 20)", func(t *testing.T) {
		assertPMFTol(0.5, 20, Tol, t)
	})
	t.Run("PMF(0.5, 200)", func(t *testing.T) {
		assertPMFTol(0.5, 200, Tol, t)
	})
}

// TestCDF implements unit tests for the bigbinomial.CDF function
func TestCDF(t *testing.T) {

	ρ := 0.5
	n := int64(200)

	floatCDF, _ := CDF(ρ, n)
	floatPMF, _ := PMF(ρ, n)

	if floatCDF(n/2) != floatCDF(n/2-1)+floatPMF(n/2) {
		t.Error("CDF(k) != CDF(k-1) + PMF(k)")
	}

	if floatCDF(n/2) != floatCDF(n/2+1)-floatPMF(n/2+1) {
		t.Error("CDF(k) != CDF(k+1) - PMF(k+1)")
	}

	if floatCDF(n/2) != floatCDF(n/2) {
		t.Error("CDF(k) != CDF(k)")
	}

	if floatPMF(-1) != 0.0 {
		t.Error("PMF(-1) != 0.0")
	}

	if floatPMF(n+1) != 0.0 {
		t.Error("PMF(n+1) != 0.0")
	}

	if floatCDF(-1) != 0.0 {
		t.Error("CDF(-1) != 0.0")
	}

	if floatCDF(n+1) != 1.0 {
		t.Error("CDF(n+1) != 1.0")
	}

	if floatCDF(n) != 1.0 {
		t.Error("CDF(n) != 1.0")
	}

	if _, err := CDF(-1.0, n); err == nil {
		t.Error("ρ < 0.0 should be an error")
	}

	if _, err := CDF(2.0, n); err == nil {
		t.Error("ρ > 1.0 should be an error")
	}

	if _, err := CDF(ρ, 0); err == nil {
		t.Error("n <= 0 should be an error")
	}
}
