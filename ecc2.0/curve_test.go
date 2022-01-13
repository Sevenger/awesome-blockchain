package ecc2

import (
	"fmt"
	"testing"
)

func TestEllipticCurve(t *testing.T) {
	curve := NewEllipticCurve64(2, 3, 97)
	G := NewPoint64(3, 6)

	p0 := curve.Mul(G, 0)
	p1 := curve.Mul(G, 1)
	p2 := curve.Mul(G, 2)
	p3 := curve.Mul(G, 3)
	p4 := curve.Mul(G, 4)
	p5 := curve.Mul(G, 5)
	p6 := curve.Mul(G, 6)
	p7 := curve.Mul(G, 7)
	p8 := curve.Mul(G, 8)
	p9 := curve.Mul(G, 9)

	assert(t, curve.OnCurve(p0), false)
	assert(t, curve.OnCurve(p1), true)
	assert(t, curve.OnCurve(p2), true)
	assert(t, curve.OnCurve(p3), true)
	assert(t, curve.OnCurve(p4), true)
	assert(t, curve.OnCurve(p5), false)
	assert(t, curve.OnCurve(p6), true)
	assert(t, curve.OnCurve(p7), true)
	assert(t, curve.OnCurve(p8), true)
	assert(t, curve.OnCurve(p9), true)

	p0.Print()
	p1.Print()
	p2.Print()
	p3.Print()
	p4.Print()
	p5.Print()
	p6.Print()
	p7.Print()
	p8.Print()
	p9.Print()
}

func TestEllipticCurve2(t *testing.T) {
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

	set := make(map[Point]struct{})
	for i := 0; ; i++ {
		gi := curve.Mul(G, i)
		if _, ok := set[gi]; ok {
			break
		}
		set[gi] = struct{}{}
		fmt.Print(i)
		gi.Print()
	}
}

func TestEllipticCurve3(t *testing.T) {
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

	println(curve.OnCurve(curve.Mul(G, 1)))
	println(curve.OnCurve(curve.Mul(G, 2)))
}

func assert(t *testing.T, res, except interface{}) {
	s := fmt.Sprintf("res: %v, except: %v", res, except)
	if res != except {
		t.Logf("faild, %s", s)
		return
	}
	t.Logf("success, %s", s)
}
