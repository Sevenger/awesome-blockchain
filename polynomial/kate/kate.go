package kate

import (
	"crypto/rand"
	"math"
	"math/big"
)

type Polynomial interface {
	Solve(x *big.Float) *big.Float
}

type Example struct{}

func (Example) Solve(x *big.Float) *big.Float {
	return new(big.Float).Mul(x, x)
}

func Commit(p Polynomial) (r, c *big.Float) {
	k, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt))
	r = new(big.Float).SetInt(k)
	c = p.Solve(r)
	return r, c
}

func Challenge(p Polynomial, r, z *big.Float) (s, w *big.Float) {
	s = p.Solve(z)
	t := func(x *big.Float) *big.Float {
		tx := new(big.Float).Sub(p.Solve(x), s)
		ty := new(big.Float).Sub(x, z)
		return tx.Quo(tx, ty)
	}
	return s, t(r)
}

func Verify(r, c, z, s, w *big.Float) bool {
	u1 := new(big.Float).Sub(r, z)
	u1.Mul(u1, w)

	u2 := new(big.Float).Sub(c, s)

	return u1.Cmp(u2) == 0
}
