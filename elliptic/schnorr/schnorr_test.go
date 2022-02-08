package schnorr

import (
	"crypto/rand"
	elliptic "elliptic/curve"
	"math/big"
	"testing"
)

func TestSchnorr(t *testing.T) {
	curve := elliptic.SECP256k1()

	privkey := GenerateKey(curve)
	pubkey := privkey.PublicKey

	msg := []byte("alice send 5U to bob")

	C, S := privkey.Signature(msg)
	if res := pubkey.Verify(C, S, msg); res != true {
		t.Errorf("Signature failed, excepted res is true")
	}

	// 篡改后的消息无法通过验证
	msg = []byte("bob send 10U to alice")
	if res := pubkey.Verify(C, S, msg); res != false {
		t.Errorf("Signature faield, excepted res is false")
	}
}

// 聚合签名
func TestSchnorr_MuSig(t *testing.T) {
	curve := elliptic.SECP256k1()

	sk1, sk2 := GenerateKey(curve), GenerateKey(curve)
	pk1, pk2 := sk1.PublicKey, sk2.PublicKey

	// 聚合公钥
	P := new(PublicKey)
	P.Curve = curve
	P.X, P.Y = curve.Add(pk1.X, pk1.Y, pk2.X, pk2.Y)

	msg := []byte("aggregate account send 100U to Bob")

	// A生成R1, 发送给B
	var k1 *big.Int
	var R1x, R1y *big.Int
	{
		k1, _ = rand.Int(rand.Reader, curve.Params().N)
		R1x, R1y = curve.MulG(k1)
	}

	// B收到R1后生成R2,S2,发送给A
	var R2x, R2y, S2 *big.Int
	{
		k2, _ := rand.Int(rand.Reader, curve.Params().N)
		R2x, R2y = curve.MulG(k2)

		// 根据A发送的S1得出聚合后的R
		Rx, Ry := curve.Add(R1x, R1y, R2x, R2y)
		// 根据R计算出C
		C := HashToInt(msg, Rx.Bytes(), Ry.Bytes())

		// 计算S2
		S2 = new(big.Int).Mul(C, sk2.sk)
		S2.Add(S2, k2)
	}

	// A收到R2,S2后生成S1，完成签名
	var S, C *big.Int
	{
		// 根据B发送的S2得出聚合后的R
		Rx, Ry := curve.Add(R1x, R1y, R2x, R2y)
		// 根据R计算出C
		C = HashToInt(msg, Rx.Bytes(), Ry.Bytes())
		// 计算S1
		S1 := new(big.Int).Mul(C, sk1.sk)
		S1.Add(S1, k1)

		// 聚合S
		S = new(big.Int).Add(S1, S2)
	}

	println(P.Verify(C, S, msg))
}
