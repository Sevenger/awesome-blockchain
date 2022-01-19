package elliptic

import (
	"fmt"
	"math/big"
	"testing"
)

func Test_CurveMulG(t *testing.T) {
	curve, _ := NewCommonCurveStr("2", "3", "97", "5", "3", "6")
	gx, gy := curve.Params().Gx, curve.Params().Gy

	tests := []struct {
		i       int
		exceptX *big.Int
		exceptY *big.Int
	}{
		{1, big.NewInt(3), big.NewInt(6)},
		{2, big.NewInt(80), big.NewInt(10)},
		{3, big.NewInt(80), big.NewInt(87)},
		{4, big.NewInt(3), big.NewInt(91)},
		{5, big.NewInt(0), big.NewInt(0)},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%dG", tt.i)
		t.Run(name, func(t *testing.T) {
			if x, y := curve.Mul(gx, gy, big.NewInt(int64(tt.i))); x.Cmp(tt.exceptX) != 0 || y.Cmp(tt.exceptY) != 0 {
				t.Errorf("%s = (%v,%v), want (%v,%v)", name, x, y, tt.exceptX, tt.exceptY)
			}
		})
	}
}

func Test_SECP256k1(t *testing.T) {
	curve := SECP256k1()
	gx, gy := curve.Params().Gx, curve.Params().Gy

	for i := int64(1); i < 1000; i++ {
		name := fmt.Sprintf("%dG", i)
		t.Run(name, func(t *testing.T) {
			if got := curve.IsOnCurve(curve.Mul(gx, gy, big.NewInt(i))); !got {
				t.Errorf("SECP256k1() = %v, want %v", got, true)
			}
		})
	}
}
