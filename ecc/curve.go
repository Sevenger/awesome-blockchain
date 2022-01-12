package ecc

import (
	"fmt"
	"math/big"
)

type EllipticCurve struct {
	a, b *big.Rat
	q    *big.Int
}

func NewEllipticCurve(A, B, Q int64) EllipticCurve {
	return EllipticCurve{
		a: new(big.Rat).SetInt64(A),
		b: new(big.Rat).SetInt64(B),
		q: big.NewInt(Q),
	}
}

func (f EllipticCurve) Add(p, q Point) Point {
	k := f.getK(p, q)

	// x = k^2 - p.X - q.X
	x := new(big.Rat).Mul(k, k)
	x.Sub(x, p.X)
	x.Sub(x, q.X)

	// y = k * (x-p.X) + p.Y
	y := new(big.Rat).Sub(x, p.X)
	y.Mul(y, k)
	y.Add(y, p.Y)
	y.Mul(y, new(big.Rat).SetInt64(-1))
	return Point{f.modQ(x), f.modQ(y)}
}

func (f EllipticCurve) getK(p, q Point) *big.Rat {
	var n, d *big.Rat

	switch {
	case p.Equal(q):
		// k = (3X^2+a)/2Y
		n = new(big.Rat).Add(
			new(big.Rat).Mul(new(big.Rat).SetInt64(3), new(big.Rat).Mul(p.X, p.X)),
			f.a,
		)
		d = new(big.Rat).Mul(new(big.Rat).SetInt64(2), p.Y)
	default:
		n = new(big.Rat).Sub(p.Y, q.Y)
		d = new(big.Rat).Sub(p.X, q.X)
	}

	k := new(big.Rat).SetFrac(n.Num(), d.Num())
	//TODO error happened, except res = 59
	fmt.Println("k mod", f.modQ(k))
	return f.modQ(k)
}

func (f EllipticCurve) modQ(x *big.Rat) *big.Rat {
	num := x.Num()
	mod := new(big.Int).Mod(num, f.q)
	return new(big.Rat).SetInt(mod)
}

func (f EllipticCurve) Mul(p Point, n int) Point {
	if n == 0 {
		return NewPoint(0, 0)
	}

	res := p
	for i := 1; i < n; i++ {
		res = f.Add(res, p)
	}
	return res
}

func (f EllipticCurve) Verify(p Point) bool {
	//return p.X^3+f.a*p.X+f.b == p.Y*p.Y
	x := new(big.Rat).Mul(p.X, new(big.Rat).Mul(p.X, p.X))
	x.Add(x, new(big.Rat).Mul(f.a, p.X))
	x.Add(x, f.b)

	y := new(big.Rat).Mul(p.Y, p.Y)

	res := new(big.Rat).Sub(x, y)

	return f.modQ(res).Cmp(&big.Rat{}) == 0
}
