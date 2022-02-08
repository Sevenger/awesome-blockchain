package schnorr

import (
	"crypto/rand"
	"crypto/sha256"
	"elliptic/curve"
	"math/big"
)

type PublicKey struct {
	Curve elliptic.Curve
	X, Y  *big.Int
}

type PrivateKey struct {
	PublicKey
	sk *big.Int
}

func GenerateKey(curve elliptic.Curve) *PrivateKey {
	d, _ := rand.Int(rand.Reader, curve.Params().N)

	key := new(PrivateKey)
	key.PublicKey.Curve = curve
	key.PublicKey.X, key.PublicKey.Y = curve.MulG(d)
	key.sk = d

	return key
}

func (p *PrivateKey) Signature(msg []byte) (C, S *big.Int) {
	return Signature(p, msg)
}

func Signature(privkey *PrivateKey, msg []byte) (C, S *big.Int) {
	curve := privkey.Curve

	r, _ := rand.Int(rand.Reader, curve.Params().N)
	Rx, Ry := curve.MulG(r)

	// C = Hash(M, R)
	C = HashToInt(msg, Rx.Bytes(), Ry.Bytes())

	// S = R + C*sk
	S = new(big.Int).Mul(C, privkey.sk)
	S = S.Add(S, r)

	return C, S
}

func (p *PublicKey) Verify(C, Z *big.Int, msg []byte) bool {
	return Verify(p, C, Z, msg)
}

func Verify(pubkey *PublicKey, C, Z *big.Int, msg []byte) bool {
	curve := pubkey.Curve

	// R = Z*G - C*pk
	u1x, u1y := curve.MulG(Z)
	u2x, u2y := curve.Mul(pubkey.X, pubkey.Y, C)
	// make C*pk = -C*pk
	u2y = new(big.Int).Sub(curve.Params().P, u2y)

	Rx, Ry := curve.Add(u1x, u1y, u2x, u2y)

	// c = Hash(M, R), and verify c==C
	c := HashToInt(msg, Rx.Bytes(), Ry.Bytes())
	return C.Cmp(c) == 0
}

func HashToInt(data ...[]byte) *big.Int {
	md := sha256.New()
	for _, v := range data {
		md.Write(v)
	}
	return new(big.Int).SetBytes(md.Sum(nil))
}
