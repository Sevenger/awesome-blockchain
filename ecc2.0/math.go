package ecc2

import (
	"math/big"
)

var convertBigIntErr = "convert %s to big.Int failed"

var Zero = new(big.Int).SetInt64(0)

func IsEqual(x, y *big.Int) bool {
	return x.Cmp(y) == 0
}

// RatMod 利用费马小定理求分数mod
func RatMod(rat *big.Rat, p *big.Int) *big.Int {
	n := rat.Num()
	d := rat.Denom()

	type powFunc func(num, n *big.Int) *big.Int
	var fastPow powFunc
	var one = big.NewInt(1)
	var two = big.NewInt(2)

	fastPow = func(num, n *big.Int) *big.Int {
		if IsEqual(n, one) {
			return big.NewInt(1)
		} else if IsEqual(Mod(n, two), one) {
			return Mod(Mul(fastPow(num, Sub(n, one)), num), p)
		} else {
			temp := Mod(fastPow(num, Div(n, two)), p)
			return Mod(Mul(temp, temp), p)
		}
	}

	fastPow = func(a, n *big.Int) *big.Int {
		res := big.NewInt(1)
		for i := big.NewInt(0); i.Cmp(n) < 0; i.Add(i, one) {
			res.Mul(res, a)
		}
		return res.Mod(res, p)
	}

	fastPow = func(a1, n *big.Int) *big.Int {
		res := big.NewInt(1)
		a := new(big.Int).Set(a1)
		a.Mod(a, p)
		for !IsEqual(n, Zero) {
			if IsEqual(new(big.Int).And(n, one), one) {
				res.Mul(res, a).Mod(res, p)
			}
			a.Mul(a, a).Mod(a, p)
			n.Rsh(n, 1)
		}
		return res
	}

	// ( n·d^(p-2) ) % p
	res := fastPow(d, Sub(p, big.NewInt(2)))

	res.Mul(res, Mod(n, p))
	res.Mod(res, p)
	return res
}

func ReliableRatMod(rat *big.Rat, p *big.Int) *big.Int {
	if rat.Denom().Int64() == 1 {
		return new(big.Int).Mod(rat.Num(), p)
	}
	res := big.NewInt(1)
	fastPow := ReliableFastPow(rat.Denom(), new(big.Int).Sub(p, big.NewInt(2)), p)
	res.Mod(rat.Num(), p).Mul(res, fastPow).Mod(res, p)
	return res
}

func ReliableFastPow(a, b, p *big.Int) *big.Int {
	var res = big.NewInt(1)
	a.Mod(a, p)
	for b.Int64() != 0 {
		if b.Int64()&1 == 1 {
			res.Mul(res, a).Mod(res, p)
		}
		b = b.Rsh(b, 1)
		a.Mul(a, a).Mod(a, p)
	}
	return res
}

func Mul(x, y *big.Int) *big.Int {
	return new(big.Int).Mul(x, y)
}

func Div(x, y *big.Int) *big.Int {
	return new(big.Int).Div(x, y)
}

func Add(x, y *big.Int) *big.Int {
	return new(big.Int).Add(x, y)
}

func Sub(x, y *big.Int) *big.Int {
	return new(big.Int).Sub(x, y)
}

func Mod(x, y *big.Int) *big.Int {
	return new(big.Int).Mod(x, y)
}
