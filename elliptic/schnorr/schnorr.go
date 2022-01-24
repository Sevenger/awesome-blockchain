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

func (p *PrivateKey) Signature(msg ...[]byte) (C, Z *big.Int) {
	return Signature(p, msg...)
}

func Signature(privkey *PrivateKey, msg ...[]byte) (C, Z *big.Int) {
	curve := privkey.Curve

	r, _ := rand.Int(rand.Reader, curve.Params().N)
	Rx, Ry := curve.MulG(r)

	// C = Hash(M, R)
	C = HashToInt(aggregate(msg...), Rx.Bytes(), Ry.Bytes())

	// Z = R + C*sk
	Z = new(big.Int).Mul(C, privkey.sk)
	Z = Z.Add(Z, r)

	return C, Z
}

func (p *PublicKey) Verify(C, Z *big.Int, msg ...[]byte) bool {
	return Verify(p, C, Z, msg...)
}

func Verify(pubkey *PublicKey, C, Z *big.Int, msg ...[]byte) bool {
	curve := pubkey.Curve

	// R = Z*G - C*pk
	u1x, u1y := curve.MulG(Z)
	u2x, u2y := curve.Mul(pubkey.X, pubkey.Y, C)
	// make C*pk = -C*pk
	u2y = new(big.Int).Sub(curve.Params().P, u2y)

	Rx, Ry := curve.Add(u1x, u1y, u2x, u2y)

	// c = Hash(M, R), and verify c==C
	c := HashToInt(aggregate(msg...), Rx.Bytes(), Ry.Bytes())
	return C.Cmp(c) == 0
}

func HashToInt(data ...[]byte) *big.Int {
	md := sha256.New()
	for _, v := range data {
		md.Write(v)
	}
	return new(big.Int).SetBytes(md.Sum(nil))
}

func aggregate(data ...[]byte) []byte {
	var agg []byte
	for _, v := range data {
		agg = append(agg, v...)
	}
	return agg
}
