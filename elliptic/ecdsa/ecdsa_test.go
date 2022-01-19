package ecdsa

import (
	elliptic "elliptic/curve"
	"fmt"
	"testing"
)

func TestECDSA(t *testing.T) {

	curve, _ := elliptic.NewCommonCurveStr("2", "3", "97", "5", "3", "6")
	curve = elliptic.SECP256k1()
	priv := GenerateKey(curve)

	msg := []byte("hello world")
	r, s := priv.Signature(msg)

	fmt.Println(priv.Verify(r, s, msg))

	msg = []byte("hello jojo")
	fmt.Println(priv.Verify(r, s, msg))
}
