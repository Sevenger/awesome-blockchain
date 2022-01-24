package schnorr

import (
	elliptic "elliptic/curve"
	"testing"
)

func TestSchnorr(t *testing.T) {
	curve := elliptic.SECP256k1()

	privkey := GenerateKey(curve)
	pubkey := privkey.PublicKey

	msg := []byte("I love you")

	C, Z := Signature(privkey, msg)

	if res := Verify(&pubkey, msg, C, Z); res != true {
		t.Errorf("Signature failed, excpeted res is true")
	}

	fakeMsg := []byte("I hate you")
	if res := Verify(&pubkey, fakeMsg, C, Z); res != false {
		t.Errorf("Verify failed, excepet res is false")
	}
}
