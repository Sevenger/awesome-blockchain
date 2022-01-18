package kate

import (
	"math/big"
	"math/rand"
	"testing"
)

func TestKate(t *testing.T) {
	var king Polynomial = Example{}
	test := []struct {
		Challenger Polynomial
		Except     bool
	}{
		{Example{}, true},
		{Faker{}, false},
	}

	r, c := Commit(king)
	for _, tt := range test {
		z := big.NewFloat(rand.Float64())
		s, w := Challenge(tt.Challenger, r, z)
		if got := Verify(r, c, z, s, w); got != tt.Except {
			t.Errorf("Except: %v, Actual: %v", tt.Except, got)
		}
	}
}

type Faker struct{}

func (Faker) Solve(x *big.Float) *big.Float {
	return new(big.Float).Set(x)
}
