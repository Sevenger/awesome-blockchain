package ECEIGamal

import (
	elliptic "elliptic/curve"
	"math/big"
	"testing"
)

func TestSign(t *testing.T) {
	priv := GenerateKey(elliptic.SECP256k1(), big.NewInt(114514))
	tests := []struct {
		msg string
	}{
		{"Hello, world"},
		{"今天我请客"},
		{"alice send 100UTX to bob"},
	}

	for _, tt := range tests {
		M := priv.Decrypt(priv.Encrypt([]byte(tt.msg)))
		eM := embedData(priv.Curve, []byte(tt.msg))
		if M.X.Cmp(eM.X) != 0 || M.Y.Cmp(eM.Y) != 0 {
			t.Errorf("Execpt M:%+v, Actul M:%+v", M, eM)
		}
	}
}
