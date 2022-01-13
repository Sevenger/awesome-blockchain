package ecc

import (
	"fmt"
	"testing"
)

func TestEllipticCurve(t *testing.T) {
	curve := NewEllipticCurve(2, 3, 97)
	G := NewPoint(3, 6)

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

func assert(t *testing.T, res, except interface{}) {
	s := fmt.Sprintf("res: %v, except: %v", res, except)
	if res != except {
		t.Logf("faild, %s", s)
		return
	}
	t.Logf("success, %s", s)
}
