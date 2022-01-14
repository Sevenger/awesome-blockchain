package ecc2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEllipticCurve(t *testing.T) {
	curve := NewEllipticCurve64(2, 3, 97)
	G := NewPoint64(3, 6)

	assert.Equal(t, false, curve.OnCurve(curve.Mul(G, 0)))
	assert.Equal(t, true, curve.OnCurve(curve.Mul(G, 1)))
	assert.Equal(t, true, curve.OnCurve(curve.Mul(G, 2)))
	assert.Equal(t, true, curve.OnCurve(curve.Mul(G, 3)))
	assert.Equal(t, true, curve.OnCurve(curve.Mul(G, 4)))
	assert.Equal(t, false, curve.OnCurve(curve.Mul(G, 5)))
	assert.Equal(t, true, curve.OnCurve(curve.Mul(G, 6)))
	assert.Equal(t, true, curve.OnCurve(curve.Mul(G, 7)))
	assert.Equal(t, true, curve.OnCurve(curve.Mul(G, 8)))
	assert.Equal(t, true, curve.OnCurve(curve.Mul(G, 9)))
}

func TestBitCoinEllipticCurve(t *testing.T) {
	curve, err := NewEllipticCurveStr("0", "7", "115792089237316195423570985008687907853269984665640564039457584007908834671663", 10)
	if err != nil {
		panic(err)
	}

	G, err := NewPointStr("55066263022277343669578718895168534326250603453777594175500187360389116729240",
		"32670510020758816978083085130507043184471273380659243275938904335757337482424",
		10)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, false, curve.OnCurve(curve.Mul(G, 0)))
	assert.Equal(t, true, curve.OnCurve(curve.Mul(G, 1)))
	assert.Equal(t, true, curve.OnCurve(curve.Mul(G, 2)))
	assert.Equal(t, true, curve.OnCurve(curve.Mul(G, 3)))
	assert.Equal(t, true, curve.OnCurve(curve.Mul(G, 4)))
	assert.Equal(t, true, curve.OnCurve(curve.Mul(G, 5)))
	assert.Equal(t, true, curve.OnCurve(curve.Mul(G, 6)))
	assert.Equal(t, true, curve.OnCurve(curve.Mul(G, 7)))
	assert.Equal(t, true, curve.OnCurve(curve.Mul(G, 8)))
	assert.Equal(t, true, curve.OnCurve(curve.Mul(G, 9)))
}
