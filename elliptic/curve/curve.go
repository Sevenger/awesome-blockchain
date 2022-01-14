package elliptic

import "math/big"

type Curve interface {
	Params() *CurveParams
	IsOnCurve(x, y *big.Int) bool
	Add(x1, y1 *big.Int, x2, y2 *big.Int) (x, y *big.Int)
	Mul(x1, y1 *big.Int, n *big.Int) (x, y *big.Int)
}

type CurveParams struct {
	A      *big.Int
	B      *big.Int
	P      *big.Int
	Gx, Gy *big.Int
}

func (curve *CurveParams) Params() *CurveParams {
	return curve
}

func (curve *CurveParams) IsOnCurve(x, y *big.Int) bool {
	y2 := new(big.Int).Mul(y, y)
	y2.Mod(y2, curve.P)
	return curve.polynomial(x).Cmp(y2) == 0
}

// return x^3+ax+b
func (curve *CurveParams) polynomial(x *big.Int) *big.Int {
	x3 := new(big.Int).Mul(x, x)
	x3.Mul(x3, x)

	ax := new(big.Int).Mul(curve.A, x)

	x3.Add(x3, ax)
	x3.Add(x3, curve.B)
	x3.Mod(x3, curve.P)
	return x3
}

func (curve *CurveParams) Add(x1, y1 *big.Int, x2, y2 *big.Int) (x, y *big.Int) {
	// def1: p1=O => p1+p2=p2
	// def2: p2=O => p1+p2=p1
	// def3: x1=x2 && (y1+y2)%order=0 => p1+p2=O
	switch {
	case curve.IsO(x1, y1):
		return x2, y2
	case curve.IsO(x2, y2):
		return x1, y1
	case x1.Cmp(x2) == 0:
		yy := new(big.Int).Add(y1, y2)
		yy.Mod(yy, curve.P)
		if yy.Cmp(big.NewInt(0)) == 0 {
			return big.NewInt(0), big.NewInt(0)
		}
	}

	slop := curve.slop(x1, y1, x2, y2)

	// x = slop^2 - p.X - order.X
	x = new(big.Int).Mul(slop, slop)
	x.Sub(x, x1)
	x.Sub(x, x2)
	x.Mod(x, curve.P)

	// y = slop * (x-p.X) + p.Y
	y = new(big.Int).Sub(x, x1)
	y.Mul(y, slop)
	y.Add(y, y1)
	y.Mul(y, new(big.Int).SetInt64(-1)) // 取y'
	y.Mod(y, curve.P)

	return x, y
}

func (curve *CurveParams) IsO(x, y *big.Int) bool {
	zero := big.NewInt(0)
	return x.Cmp(zero) == 0 && y.Cmp(zero) == 0
}

func (curve *CurveParams) slop(x1, y1 *big.Int, x2, y2 *big.Int) *big.Int {
	var n, d *big.Int

	switch {
	case x1.Cmp(x2) == 0 && y1.Cmp(y2) == 0:
		// k = (3X^2+a)/2Y
		n = new(big.Int).Add(
			new(big.Int).Mul(big.NewInt(3), new(big.Int).Mul(x1, x1)),
			curve.A,
		)
		d = new(big.Int).Mul(big.NewInt(2), y1)

	default:
		n = new(big.Int).Sub(y1, y2)
		d = new(big.Int).Sub(x1, x2)
	}

	slop := new(big.Rat).SetFrac(n, d)
	return ratMod(slop, curve.P)
}

func ratMod(rat *big.Rat, p *big.Int) *big.Int {
	n := rat.Num()
	d := rat.Denom()
	if d.Cmp(big.NewInt(1)) == 0 {
		return new(big.Int).Mod(n, p)
	}

	fastPow := func(a, n *big.Int, p *big.Int) *big.Int {
		res := big.NewInt(1)
		for n.Int64() != 0 {
			if n.Int64()&1 == 1 {
				res.Mul(res, a).Mod(res, p)
			}
			a.Mul(a, a).Mod(a, p)
			n.Rsh(n, 1)
		}
		return res
	}
	// ( n·d^(p-2) ) % p
	res := fastPow(d, new(big.Int).Sub(p, big.NewInt(2)), p)
	res.Mul(res, n)
	res.Mod(res, p)
	return res
}

func (curve *CurveParams) Mul(x1, y1 *big.Int, n *big.Int) (x, y *big.Int) {
	x, y = big.NewInt(0), big.NewInt(0)
	for i := big.NewInt(0); i.Cmp(n) < 0; i.Add(i, big.NewInt(1)) {
		x, y = curve.Add(x, y, x1, y1)
	}
	return x, y
}

func (curve *CurveParams) MulG(n *big.Int) (x, y *big.Int) {
	return curve.Mul(curve.Gx, curve.Gy, n)
}
