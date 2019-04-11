package BigBinomial

import (
	"math"
	"math/big"
	"testing"
)

// TestPow tests the BigBinomial.Pow function
func TestPow(t *testing.T) {

	Zero := big.NewFloat(0.0)
	NegZero := big.NewFloat(0.0).Neg(Zero)
	One := big.NewFloat(1.0)
	NegOne := big.NewFloat(-1.0)
	Two := big.NewFloat(2.0)
	Ten := big.NewFloat(10.0)
	Inf := big.NewFloat(0.0).SetInf(false)
	NegInf := big.NewFloat(0.0).SetInf(true)

	bigAssert := func(got, want *big.Float, t *testing.T) {
		t.Helper()
		if got.Cmp(want) != 0 {
			t.Fatal("Wanted:", want, "Got:", got)
		}
	}

	bigAssertTol := func(got *big.Float, want, tol float64, t *testing.T) {
		t.Helper()
		gotFloat, _ := got.Float64()
		if math.Abs(gotFloat-want)/(gotFloat+want) > tol {
			t.Fatal("Wanted:", want, "Got:", gotFloat, "Tolerance:", tol)
		}
	}

	t.Run("Pow(Inf, 0)", func(t *testing.T) {
		bigAssert(Pow(Inf, 0), One, t)
	})
	t.Run("Pow(Inf, 1)", func(t *testing.T) {
		bigAssert(Pow(Inf, 1), Inf, t)
	})
	t.Run("Pow(Inf, 2)", func(t *testing.T) {
		bigAssert(Pow(Inf, 2), Inf, t)
	})
	t.Run("Pow(Inf, -1)", func(t *testing.T) {
		bigAssert(Pow(Inf, -1), Zero, t)
	})
	t.Run("Pow(Inf, -2)", func(t *testing.T) {
		bigAssert(Pow(Inf, -2), Zero, t)
	})

	t.Run("Pow(NegInf, 0)", func(t *testing.T) {
		bigAssert(Pow(NegInf, 0), Pow(NegZero, 0), t)
	})
	t.Run("Pow(NegInf, 1)", func(t *testing.T) {
		bigAssert(Pow(NegInf, 1), Pow(NegZero, -1), t)
	})
	t.Run("Pow(NegInf, 2)", func(t *testing.T) {
		bigAssert(Pow(NegInf, 2), Pow(NegZero, -2), t)
	})
	t.Run("Pow(NegInf, -1)", func(t *testing.T) {
		bigAssert(Pow(NegInf, -1), Pow(NegZero, 1), t)
	})
	t.Run("Pow(NegInf, -2)", func(t *testing.T) {
		bigAssert(Pow(NegInf, -2), Pow(NegZero, 2), t)
	})

	t.Run("Pow(Zero, 0)", func(t *testing.T) {
		bigAssert(Pow(Zero, 0), One, t)
	})
	t.Run("Pow(Zero, 1)", func(t *testing.T) {
		bigAssert(Pow(Zero, 1), Zero, t)
	})
	t.Run("Pow(Zero, 2)", func(t *testing.T) {
		bigAssert(Pow(Zero, 2), Zero, t)
	})
	t.Run("Pow(Zero, -1)", func(t *testing.T) {
		bigAssert(Pow(Zero, -1), Inf, t)
	})
	t.Run("Pow(Zero, -2)", func(t *testing.T) {
		bigAssert(Pow(Zero, -2), Inf, t)
	})

	t.Run("Pow(NegZero, 0)", func(t *testing.T) {
		bigAssert(Pow(NegZero, 0), One, t)
	})
	t.Run("Pow(NegZero, 1)", func(t *testing.T) {
		bigAssert(Pow(NegZero, 1), NegZero, t)
	})
	t.Run("Pow(NegZero, 2)", func(t *testing.T) {
		bigAssert(Pow(NegZero, 2), Zero, t)
	})
	t.Run("Pow(NegZero, -1)", func(t *testing.T) {
		bigAssert(Pow(NegZero, -1), NegInf, t)
	})
	t.Run("Pow(NegZero, -2)", func(t *testing.T) {
		bigAssert(Pow(NegZero, -2), Inf, t)
	})

	t.Run("Pow(One, 0)", func(t *testing.T) {
		bigAssert(Pow(One, 0), One, t)
	})
	t.Run("Pow(One, 1)", func(t *testing.T) {
		bigAssert(Pow(One, 1), One, t)
	})
	t.Run("Pow(One, 2)", func(t *testing.T) {
		bigAssert(Pow(One, 2), One, t)
	})
	t.Run("Pow(One, -1)", func(t *testing.T) {
		bigAssert(Pow(One, -1), One, t)
	})
	t.Run("Pow(One, -2)", func(t *testing.T) {
		bigAssert(Pow(One, -2), One, t)
	})

	t.Run("Pow(NegOne, 0)", func(t *testing.T) {
		bigAssert(Pow(NegOne, 0), One, t)
	})
	t.Run("Pow(NegOne, 1)", func(t *testing.T) {
		bigAssert(Pow(NegOne, 1), NegOne, t)
	})
	t.Run("Pow(NegOne, 2)", func(t *testing.T) {
		bigAssert(Pow(NegOne, 2), One, t)
	})
	t.Run("Pow(NegOne, 1)", func(t *testing.T) {
		bigAssert(Pow(NegOne, 1), NegOne, t)
	})
	t.Run("Pow(NegOne, 2)", func(t *testing.T) {
		bigAssert(Pow(NegOne, 2), One, t)
	})

	t.Run("Pow(Two, 0)", func(t *testing.T) {
		bigAssert(Pow(Two, 0), One, t)
	})
	t.Run("Pow(Two, 1)", func(t *testing.T) {
		bigAssert(Pow(Two, 1), Two, t)
	})
	t.Run("Pow(Two, 2)", func(t *testing.T) {
		bigAssert(Pow(Two, 2), big.NewFloat(0).Mul(Two, Two), t)
	})
	t.Run("Pow(Two, -1)", func(t *testing.T) {
		bigAssert(Pow(Two, -1), big.NewFloat(0).Quo(One, Two), t)
	})
	t.Run("Pow(Two, -2)", func(t *testing.T) {
		bigAssert(Pow(Two, -2), big.NewFloat(0).Quo(One, big.NewFloat(0).Mul(Two, Two)), t)
	})

	t.Run("Pow(Ten, 2)", func(t *testing.T) {
		bigAssert(Pow(Ten, 2), big.NewFloat(math.Pow(10.0, 2)), t)
	})
	t.Run("Pow(Ten, 5)", func(t *testing.T) {
		bigAssert(Pow(Ten, 5), big.NewFloat(math.Pow(10.0, 5)), t)
	})
	t.Run("Pow(Ten, 25)", func(t *testing.T) {
		bigAssert(Pow(Ten, 25), big.NewFloat(math.Pow(10.0, 25)), t)
	})
	t.Run("Pow(Ten, 250)", func(t *testing.T) {
		bigAssert(Pow(Ten, 250), big.NewFloat(math.Pow(10.0, 250)), t)
	})

	// These don't work as equalities because the two libraries calculate negative exponents differently

	t.Run("Pow(Ten, -2)", func(t *testing.T) {
		bigAssertTol(Pow(Ten, -2), math.Pow(10.0, -2), 1.0e-13, t)
	})
	t.Run("Pow(Ten, -5)", func(t *testing.T) {
		bigAssertTol(Pow(Ten, -5), math.Pow(10.0, -5), 1.0e-13, t)
	})
	t.Run("Pow(Ten, -25)", func(t *testing.T) {
		bigAssertTol(Pow(Ten, -25), math.Pow(10.0, -25), 1.0e-13, t)
	})
	t.Run("Pow(Ten, -250)", func(t *testing.T) {
		bigAssertTol(Pow(Ten, -250), math.Pow(10.0, -250), 1.0e-13, t)
	})
}

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

	t.Run("PMF(0.5, 2)", func(t *testing.T) {
		assertPMFTol(0.5, 2, 1.0e-14, t)
	})
	t.Run("PMF(0.5, 20)", func(t *testing.T) {
		assertPMFTol(0.5, 20, 1.0e-14, t)
	})
	t.Run("PMF(0.5, 200)", func(t *testing.T) {
		assertPMFTol(0.5, 200, 1.0e-14, t)
	})

	t.Run("PMF(0.15, 2)", func(t *testing.T) {
		assertPMFTol(0.15, 2, 1.0e-14, t)
	})
	t.Run("PMF(0.15, 20)", func(t *testing.T) {
		assertPMFTol(0.15, 20, 1.0e-14, t)
	})
	t.Run("PMF(0.15, 200)", func(t *testing.T) {
		assertPMFTol(0.15, 200, 1.0e-14, t)
	})

	t.Run("PMF(0.5, 2)", func(t *testing.T) {
		assertPMFTol(0.5, 2, 1.0e-14, t)
	})
	t.Run("PMF(0.5, 20)", func(t *testing.T) {
		assertPMFTol(0.5, 20, 1.0e-14, t)
	})
	t.Run("PMF(0.5, 200)", func(t *testing.T) {
		assertPMFTol(0.5, 200, 1.0e-14, t)
	})

	t.Run("PMF(0.5, 2)", func(t *testing.T) {
		assertPMFTol(0.5, 2, 1.0e-14, t)
	})
	t.Run("PMF(0.5, 20)", func(t *testing.T) {
		assertPMFTol(0.5, 20, 1.0e-14, t)
	})
	t.Run("PMF(0.5, 200)", func(t *testing.T) {
		assertPMFTol(0.5, 200, 1.0e-14, t)
	})
}

func TestCDF(t *testing.T) {

	floatCDF, _ := CDF(0.5, 200)
	floatPMF, _ := PMF(0.5, 200)

	if floatCDF(100) != floatCDF(99)+floatPMF(100) {
		t.Error("CDF(k) != CDF(k-1) + PMF(k)")
	}

	if floatCDF(100) != floatCDF(101)-floatPMF(101) {
		t.Error("CDF(k) != CDF(k+1) - PMF(k+1)")
	}

	if floatCDF(100) != floatCDF(100) {
		t.Error("CDF(k) != CDF(k)")
	}

	if floatPMF(-1) != 0.0 {
		t.Error("PMF(-1) != 0.0")
	}

	if floatPMF(201) != 0.0 {
		t.Error("PMF(n+1) != 0.0")
	}

	if floatCDF(-1) != 0.0 {
		t.Error("CDF(-1) != 0.0")
	}

	if floatCDF(201) != 1.0 {
		t.Error("CDF(n+1) != 1.0")
	}

	if _, err := CDF(-1.0, 200); err == nil {
		t.Error("ρ < 0.0 should be an error")
	}

	if _, err := CDF(2.0, 200); err == nil {
		t.Error("ρ > 1.0 should be an error")
	}

	if _, err := CDF(0.5, 0); err == nil {
		t.Error("n <= 0 should be an error")
	}
}
