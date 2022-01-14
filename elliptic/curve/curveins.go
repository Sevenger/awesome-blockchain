package elliptic

import "math/big"

type CommonCurve struct {
	*CurveParams
}

func NewCommonCurve(A, B, P, Gx, Gy *big.Int) Curve {
	return &CommonCurve{
		CurveParams: &CurveParams{A, B, P, Gx, Gy},
	}
}

func NewCommonCurveStr(A, B, P, Gx, Gy string) (Curve, bool) {
	base := 10
	var a, b, p, gx, gy *big.Int
	var ok bool
	if a, ok = new(big.Int).SetString(A, base); !ok {
		return nil, false
	}
	if b, ok = new(big.Int).SetString(B, base); !ok {
		return nil, false
	}
	if p, ok = new(big.Int).SetString(P, base); !ok {
		return nil, false
	}
	if gx, ok = new(big.Int).SetString(Gx, base); !ok {
		return nil, false
	}
	if gy, ok = new(big.Int).SetString(Gy, base); !ok {
		return nil, false
	}
	return NewCommonCurve(a, b, p, gx, gy), true
}

func SECP256k1() Curve {
	curve, _ := NewCommonCurveStr(
		"0", "7", "115792089237316195423570985008687907853269984665640564039457584007908834671663",
		"55066263022277343669578718895168534326250603453777594175500187360389116729240",
		"32670510020758816978083085130507043184471273380659243275938904335757337482424")
	return curve
}
