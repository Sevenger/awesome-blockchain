package schnorr

import (
	elliptic "elliptic/curve"
	"testing"
)

func TestSchnorr(t *testing.T) {
	curve := elliptic.SECP256k1()

	privkey := GenerateKey(curve)
	pubkey := privkey.PublicKey

	//	聚合签名
	msg1 := []byte("alice send 5U to bob")
	msg2 := []byte("alice send 10U to bob")
	msg3 := []byte("bob send 100U to alice")

	C, Z := privkey.Signature(msg1, msg2, msg3)
	if res := pubkey.Verify(C, Z, msg1, msg2, msg3); res != true {
		t.Errorf("Signature failed, excepted res is true")
	}

	// 只要有一条msg被篡改验证就无法通过
	msg2 = []byte("bob send 10U to alice")
	if res := pubkey.Verify(C, Z, msg1, msg2, msg3); res != false {
		t.Errorf("Signature faield, excepted res is false")
	}
}
