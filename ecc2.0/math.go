package ecc2

import (
	"math/big"
)

var convertBigIntErr = "convert %s to big.Int failed"

var Zero = new(big.Int).SetInt64(0)
var One = new(big.Int).SetInt64(1)
var Two = new(big.Int).SetInt64(2)

func IsEqual(x, y *big.Int) bool {
	return x.Cmp(y) == 0
}

// RatMod 利用费马小定理求分数mod
func RatMod(rat *big.Rat, p *big.Int) *big.Int {
	n := rat.Num()
	d := rat.Denom()
	if IsEqual(d, One) {
		return new(big.Int).Mod(n, p)
	}

	// ( n·d^(p-2) ) % p
	res := FastPow(d, new(big.Int).Sub(p, Two), p)
	res.Mul(res, n)
	res.Mod(res, p)
	return res
}

func FastPow(a, n, p *big.Int) *big.Int {
	if IsEqual(n, Zero) {
		return big.NewInt(1)
	} else if IsEqual(new(big.Int).Mod(n, Two), One) {
		//	奇数 a^n == a^(n-1) * a
		res := FastPow(a, new(big.Int).Sub(n, One), p)
		res.Mod(res, p)
		return new(big.Int).Mul(res, a)
	} else {
		//	偶数 a^n == a^(n/2) * a^(n/2)
		res := FastPow(a, new(big.Int).Div(n, Two), p)
		res.Mod(res, p)
		return new(big.Int).Mul(res, res)
	}
}
