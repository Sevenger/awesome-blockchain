package ecc2

import (
	"fmt"
	"math/big"
)

type EllipticCurve struct {
	a, b  *big.Int
	order *big.Int
}

func NewEllipticCurve(A, B, Order *big.Int) EllipticCurve {
	return EllipticCurve{A, B, Order}
}

func NewEllipticCurve64(A, B, Order int64) EllipticCurve {
	return NewEllipticCurve(big.NewInt(A), big.NewInt(B), big.NewInt(Order))
}

func NewEllipticCurveStr(A, B, Order string, base int) (EllipticCurve, error) {
	a, ok := new(big.Int).SetString(A, base)
	if !ok {
		return EllipticCurve{}, fmt.Errorf(convertBigIntErr, "A")
	}
	b, ok := new(big.Int).SetString(B, base)
	if !ok {
		return EllipticCurve{}, fmt.Errorf(convertBigIntErr, "B")
	}
	order, ok := new(big.Int).SetString(Order, base)
	if !ok {
		return EllipticCurve{}, fmt.Errorf(convertBigIntErr, "Order")
	}

	return NewEllipticCurve(a, b, order), nil
}

func (f EllipticCurve) Add(p, q Point) Point {
	// def1: p1=O => p1+p2=p2
	// def2: p2=O => p1+p2=p1
	// def3: x1=x2 && (y1+y2)%order=0 => p1+p2=O

	switch {
	case p.IsO():
		return q
	case q.IsO():
		return p
	case IsEqual(p.X, q.X):
		yy := new(big.Int).Add(p.Y, q.Y)
		if IsEqual(f.ModOrder(yy), Zero) {
			return PointO
		}
	}

	slop := f.GetSlop(p, q)

	// x = slop^2 - p.X - order.X
	x := new(big.Int).Mul(slop, slop)
	x.Sub(x, p.X)
	x.Sub(x, q.X)

	// y = slop * (x-p.X) + p.Y
	y := new(big.Int).Sub(x, p.X)
	y.Mul(y, slop)
	y.Add(y, p.Y)
	y.Mul(y, new(big.Int).SetInt64(-1))
	return Point{f.ModOrder(x), f.ModOrder(y)}
}

func (f EllipticCurve) Mul(p Point, n int) Point {
	if n == 0 {
		return PointO
	}

	res := p
	for i := 1; i < n; i++ {
		res = f.Add(res, p)
	}
	return res
}

// GetSlop 计算两点斜率
func (f EllipticCurve) GetSlop(p, q Point) *big.Int {
	var n, d *big.Int

	switch {
	case p.Equal(q):
		// k = (3X^2+a)/2Y
		n = new(big.Int).Add(
			new(big.Int).Mul(big.NewInt(3), new(big.Int).Mul(p.X, p.X)),
			f.a,
		)
		d = new(big.Int).Mul(big.NewInt(2), p.Y)

	default:
		n = new(big.Int).Sub(p.Y, q.Y)
		d = new(big.Int).Sub(p.X, q.X)
	}

	slop := new(big.Rat).SetFrac(n, d)

	return RatMod(slop, f.order)
}

// OnCurve 校验点P是否在曲线上
func (f EllipticCurve) OnCurve(p Point) bool {
	//return p.X^3+f.a*p.X+f.b == p.Y*p.Y
	x3 := new(big.Int).Mul(p.X, new(big.Int).Mul(p.X, p.X))

	ax := new(big.Int).Mul(f.a, p.X)
	b := new(big.Int).Set(f.b)

	y2 := new(big.Int).Mul(p.Y, p.Y)

	res := new(big.Int).Add(x3, new(big.Int).Add(ax, b))

	res = res.Sub(res, y2)
	return IsEqual(f.ModOrder(res), Zero)
}

func (f EllipticCurve) ModOrder(x *big.Int) *big.Int {
	return new(big.Int).Mod(x, f.order)
}
