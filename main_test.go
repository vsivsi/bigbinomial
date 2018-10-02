package BigBinomial

import (
   "testing"
	"math"
	"math/big"
)

// TestExp tests the BigBinomial.Exp function
func TestPow(t *testing.T) {

   Zero := big.NewFloat(0.0)
   NegZero := big.NewFloat(0.0).Neg(Zero)
   One := big.NewFloat(1.0)
   NegOne := big.NewFloat(-1.0)
   Two := big.NewFloat(2.0)
   Ten := big.NewFloat(10.0)
   Inf := big.NewFloat(0.0).SetInf(false)
   NegInf := big.NewFloat(0.0).SetInf(true)

   t.Run("X=-Inf,n=0", func(t *testing.T) {
      if Pow(Inf, 0).Cmp(One) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=+Inf,n=1", func(t *testing.T) {
      if Pow(Inf, 1).Cmp(Inf) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=+Inf,n=2", func(t *testing.T) {
      if Pow(Inf, 2).Cmp(Inf) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=+Inf,n=-1", func(t *testing.T) {
      if Pow(Inf, -1).Cmp(Zero) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=+Inf,n=-2", func(t *testing.T) {
      if Pow(Inf, -2).Cmp(Zero) != 0 {
         t.FailNow()
      }
   })

   t.Run("X=-Inf,n=0", func(t *testing.T) {
      if Pow(NegInf, 0).Cmp(Pow(NegZero, 0)) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=-Inf,n=1", func(t *testing.T) {
      if Pow(NegInf, 1).Cmp(Pow(NegZero, -1)) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=-Inf,n=2", func(t *testing.T) {
      if Pow(NegInf, 2).Cmp(Pow(NegZero, -2)) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=-Inf,n=-1", func(t *testing.T) {
      if Pow(NegInf, -1).Cmp(Pow(NegZero, 1)) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=-Inf,n=-2", func(t *testing.T) {
      if Pow(NegInf, -2).Cmp(Pow(NegZero, 2)) != 0 {
         t.FailNow()
      }
   })

   t.Run("X=0.0,n=0", func(t *testing.T) {
      if Pow(Zero, 0).Cmp(One) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=0.0,n=1", func(t *testing.T) {
      if Pow(Zero, 1).Cmp(Zero) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=0.0,n=2", func(t *testing.T) {
      if Pow(Zero, 2).Cmp(Zero) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=0.0,n=-1", func(t *testing.T) {
      if Pow(Zero, -1).Cmp(Inf) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=0.0,n=-2", func(t *testing.T) {
      if Pow(Zero, -2).Cmp(Inf) != 0 {
         t.FailNow()
      }
   })

   t.Run("X=-0.0,n=0", func(t *testing.T) {
      if Pow(NegZero, 0).Cmp(One) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=-0.0,n=1", func(t *testing.T) {
      if Pow(NegZero, 1).Cmp(NegZero) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=-0.0,n=2", func(t *testing.T) {
      if Pow(NegZero, 2).Cmp(Zero) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=-0.0,n=-1", func(t *testing.T) {
      if Pow(NegZero, -1).Cmp(NegInf) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=-0.0,n=-2", func(t *testing.T) {
      if Pow(NegZero, -2).Cmp(Inf) != 0 {
         t.FailNow()
      }
   })

   t.Run("X=1.0,n=0", func(t *testing.T) {
      if Pow(One, 0).Cmp(One) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=1.0,n=1", func(t *testing.T) {
         if Pow(One, 1).Cmp(One) != 0 {
            t.FailNow()
         }
   })
   t.Run("X=1.0,n=2", func(t *testing.T) {
      if Pow(One, 2).Cmp(One) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=1.0,n=-1", func(t *testing.T) {
         if Pow(One, -1).Cmp(One) != 0 {
            t.FailNow()
         }
   })
   t.Run("X=1.0,n=-2", func(t *testing.T) {
      if Pow(One, -2).Cmp(One) != 0 {
         t.FailNow()
      }
   })

   t.Run("X=-1.0,n=0", func(t *testing.T) {
      if Pow(NegOne, 0).Cmp(One) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=-1.0,n=1", func(t *testing.T) {
         if Pow(NegOne, 1).Cmp(NegOne) != 0 {
            t.FailNow()
         }
   })
   t.Run("X=-1.0,n=2", func(t *testing.T) {
         if Pow(NegOne, 2).Cmp(One) != 0 {
            t.FailNow()
         }
   })
   t.Run("X=-1.0,n=-1", func(t *testing.T) {
         if Pow(NegOne, 1).Cmp(NegOne) != 0 {
            t.FailNow()
         }
   })
   t.Run("X=-1.0,n=-2", func(t *testing.T) {
      if Pow(NegOne, 2).Cmp(One) != 0 {
         t.FailNow()
      }
   })

   t.Run("X=2.0,n=0", func(t *testing.T) {
      if Pow(Two, 0).Cmp(One) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=2.0,n=1", func(t *testing.T) {
      if Pow(Two, 1).Cmp(Two) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=2.0,n=2", func(t *testing.T) {
      if Pow(Two, 2).Cmp(big.NewFloat(0).Mul(Two, Two)) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=2.0,n=-1", func(t *testing.T) {
      if Pow(Two, -1).Cmp(big.NewFloat(0).Quo(One, Two)) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=2.0,n=-2", func(t *testing.T) {
      if Pow(Two, -2).Cmp(big.NewFloat(0).Quo(One, big.NewFloat(0).Mul(Two, Two))) != 0 {
         t.FailNow()
      }
   })

   t.Run("X=10.0,n=2", func(t *testing.T) {
      if Pow(Ten, 2).Cmp(big.NewFloat(math.Pow(10.0, 2))) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=10.0,n=5", func(t *testing.T) {
      if Pow(Ten, 5).Cmp(big.NewFloat(math.Pow(10.0, 5))) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=10.0,n=25", func(t *testing.T) {
      if Pow(Ten, 25).Cmp(big.NewFloat(math.Pow(10.0, 25))) != 0 {
         t.FailNow()
      }
   })
   t.Run("X=10.0,n=250", func(t *testing.T) {
      if Pow(Ten, 250).Cmp(big.NewFloat(math.Pow(10.0, 250))) != 0 {
         t.FailNow()
      }
   })

   // These don't work as equalities because the two libraries calculate negative exponents differently

   t.Run("X=10.0,n=-2", func(t *testing.T) {
      a, _ := Pow(Ten, -2).Float64()
      b := math.Pow(10.0, -2)
      if math.Abs(a - b) / (a + b) > 1.0e-13 {
         t.FailNow()
      }
   })
   t.Run("X=10.0,n=-5", func(t *testing.T) {
      a, _ := Pow(Ten, -5).Float64()
      b := math.Pow(10.0, -5)
      if math.Abs(a - b) / (a + b) > 1.0e-13 {
         t.FailNow()
      }
   })
   t.Run("X=10.0,n=-25", func(t *testing.T) {
      a, _ := Pow(Ten, -25).Float64()
      b := math.Pow(10.0, -25)
      if math.Abs(a - b) / (a + b) > 1.0e-13 {
         t.FailNow()
      }
   })
   t.Run("X=10.0,n=-250", func(t *testing.T) {
      a, _ := Pow(Ten, -250).Float64()
      b := math.Pow(10.0, -250)
      if math.Abs(a - b) / (a + b) > 1.0e-13 {
         t.FailNow()
      }
   })
}

func TestPMF(t *testing.T) {

   BinomialPMF := func (ρ float64, n int64) func(k int64) float64 {
       return func(k int64) float64 {
           p := math.Pow(ρ, float64(k)) * math.Pow(1-ρ, float64(n-k))
           p *= math.Gamma(float64(n+1)) / (math.Gamma(float64(k+1)) * math.Gamma(float64(n-k+1)))
           return p
       }
   }

   BinError := func (p float64, n int64) float64 {
      FloatPMF := BinomialPMF(p, n)
      BigPMF := PMF(p, n)
      err := 0.0
      for x := int64(0); x <= n; x++ {
         err += math.Abs(BigPMF(x) - FloatPMF(x))
      }
      return err
   }

   t.Run("p=0.5,n=2", func(t *testing.T) {
      if BinError(0.5, 2) > 1.0e-14 {
         t.FailNow()
      }
   })
   t.Run("p=0.5,n=20", func(t *testing.T) {
      if BinError(0.5, 20) > 1.0e-14 {
         t.FailNow()
      }
   })
   t.Run("p=0.5,n=200", func(t *testing.T) {
      if BinError(0.5, 200) > 1.0e-14 {
         t.FailNow()
      }
   })

   t.Run("p=0.05,n=2", func(t *testing.T) {
      if BinError(0.5, 2) > 1.0e-14 {
         t.FailNow()
      }
   })
   t.Run("p=0.05,n=20", func(t *testing.T) {
      if BinError(0.5, 20) > 1.0e-14 {
         t.FailNow()
      }
   })
   t.Run("p=0.05,n=200", func(t *testing.T) {
      if BinError(0.5, 200) > 1.0e-14 {
         t.FailNow()
      }
   })

   t.Run("p=0.0005,n=2", func(t *testing.T) {
      if BinError(0.5, 2) > 1.0e-14 {
         t.FailNow()
      }
   })
   t.Run("p=0.0005,n=20", func(t *testing.T) {
      if BinError(0.5, 20) > 1.0e-14 {
         t.FailNow()
      }
   })
   t.Run("p=0.005,n=200", func(t *testing.T) {
      if BinError(0.5, 200) > 1.0e-14 {
         t.FailNow()
      }
   })
}
