package main

import (
	"fmt"
	"math/big"
)

type Point struct {
	X, Y *big.Rat
}

func (p Point) Print() {
	fmt.Printf("(%v, %v)\n", p.X, p.Y)
}

func NewPoint(x, y int64) Point {
	return Point{
		X: new(big.Rat).SetInt64(x),
		Y: new(big.Rat).SetInt64(y),
	}
}

func (p Point) Equal(q Point) bool {
	return p.X.Cmp(q.X) == 0 && p.Y.Cmp(q.Y) == 0
}

type ECCFunc struct {
	A, B *big.Rat
	Q    *big.Int
}

func NewECCFunc(A, B, Q int64) ECCFunc {
	return ECCFunc{
		A: new(big.Rat).SetInt64(A),
		B: new(big.Rat).SetInt64(B),
		Q: big.NewInt(Q),
	}
}

func (f ECCFunc) Add(p, q Point) Point {
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

func (f ECCFunc) getK(p, q Point) *big.Rat {
	var n, d *big.Rat

	switch {
	case p.Equal(q):
		// k = (3X^2+A)/2Y
		n = new(big.Rat).Add(
			new(big.Rat).Mul(new(big.Rat).SetInt64(3), new(big.Rat).Mul(p.X, p.X)),
			f.A,
		)
		d = new(big.Rat).Mul(new(big.Rat).SetInt64(2), p.Y)
	default:
		n = new(big.Rat).Sub(p.Y, q.Y)
		d = new(big.Rat).Sub(p.X, q.X)
	}

	k := new(big.Rat).SetFrac(n.Num(), d.Num())
	return f.modQ(k)
}

func (f ECCFunc) modQ(x *big.Rat) *big.Rat {
	num := x.Num()
	mod := new(big.Int).Mod(num, f.Q)
	return new(big.Rat).SetInt(mod)
}

func (f ECCFunc) Mul(p Point, n int) Point {
	if n == 0 {
		return NewPoint(0, 0)
	}

	res := p
	for i := 1; i < n; i++ {
		res = f.Add(res, p)
	}
	return res
}

func (f ECCFunc) Verify(p Point) bool {
	//return p.X^3+f.A*p.X+f.B == p.Y*p.Y
	x := new(big.Rat).Mul(p.X, new(big.Rat).Mul(p.X, p.X))
	x.Add(x, new(big.Rat).Mul(f.A, p.X))
	x.Add(x, f.B)

	y := new(big.Rat).Mul(p.Y, p.Y)

	res := new(big.Rat).Sub(x, y)

	return f.modQ(res).Cmp(&big.Rat{}) == 0
}

func main() {
	f := NewECCFunc(-15, 18, 17)
	p := NewPoint(7, 16)

	p = f.Mul(p, 1)
	fmt.Println(f.Verify(p), p)
	p = f.Mul(p, 2)
	fmt.Println(f.Verify(p), p)
	p = f.Mul(p, 3)
	fmt.Println(f.Verify(p), p)
}
