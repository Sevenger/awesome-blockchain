package ecdsa

import (
	"crypto/rand"
	"crypto/sha1"
	"elliptic/curve"
	"math/big"
)

type PublicKey struct {
	Curve elliptic.Curve
	X, Y  *big.Int
}

func (p *PublicKey) Verify(r, s *big.Int, msg []byte) bool {
	return Verify(p, r, s, msg)
}

type PrivateKey struct {
	PublicKey
	D *big.Int
}

func (p *PrivateKey) Signature(data []byte) (r, s *big.Int) {
	return Signature(p, data)
}

func GenerateKey(curve elliptic.Curve) *PrivateKey {
	k, _ := rand.Int(rand.Reader, curve.Params().N)

	prikey := new(PrivateKey)
	prikey.PublicKey.Curve = curve
	prikey.PublicKey.X, prikey.PublicKey.Y = curve.MulG(k)
	prikey.D = k
	return prikey
}

func Signature(priv *PrivateKey, msg []byte) (r, s *big.Int) {
	N := priv.Curve.Params().N

	k, _ := rand.Int(rand.Reader, N)
	kInv := new(big.Int).ModInverse(k, N)

	r, _ = priv.Curve.MulG(k)
	r.Mod(r, N)

	z := MsgToInt(msg)
	dr := new(big.Int).Mul(priv.D, r)
	zdr := new(big.Int).Add(z, dr)

	s = new(big.Int).Mul(zdr, kInv)
	s.Mod(s, N)

	return r, s
}

func Verify(pub *PublicKey, r, s *big.Int, msg []byte) bool {
	N := pub.Curve.Params().N

	var w *big.Int
	w = new(big.Int).ModInverse(s, N)

	z := MsgToInt(msg)

	u1 := z.Mul(z, w)
	u1.Mod(u1, N)
	u2 := w.Mul(r, w)
	u2.Mod(u2, N)

	x1, y1 := pub.Curve.MulG(u1)
	x2, y2 := pub.Curve.Mul(pub.X, pub.Y, u2)
	x, _ := pub.Curve.Add(x1, y1, x2, y2)
	x.Mod(x, N)
	return x.Cmp(r) == 0
}

func MsgToInt(msg []byte) *big.Int {
	md := sha1.New()
	md.Write(msg)
	hashed := md.Sum(nil)
	hashedInt := new(big.Int).SetBytes(hashed)
	return hashedInt
}
