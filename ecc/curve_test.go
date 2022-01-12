package ecc

import (
	"fmt"
	"testing"
)

func TestEllipticCurve(t *testing.T) {
	curve := NewEllipticCurve(2, 3, 97)

	p := NewPoint(3, 6)
	assert(t, curve.Verify(p), true)
	curve.Add(p, p).Print()
	assert(t, curve.Verify(curve.Mul(p, 2)), true)
}

func assert(t *testing.T, res, except interface{}) {
	s := fmt.Sprintf("res: %v, except: %v", res, except)
	if res != except {
		t.Logf("faild, %s", s)
		return
	}
	t.Logf("success, %s", s)
}
