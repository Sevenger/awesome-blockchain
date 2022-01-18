package ecdsa

import (
	elliptic "elliptic/curve"
	"fmt"
	"math/big"
	"testing"
)

func TestECDSA(t *testing.T) {
	priv := GenerateKey(elliptic.SECP256k1(), big.NewInt(123))

	msg := []byte("hello world")
	r, s := priv.Sign(msg)

	fmt.Println(priv.Verify(r, s, msg))
}
