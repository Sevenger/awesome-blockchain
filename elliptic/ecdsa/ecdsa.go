package ecdsa

import (
	"bytes"
	"crypto/sha1"
	"elliptic/curve"
	"encoding/binary"
	"math/big"
	"math/rand"
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

func (p *PrivateKey) Sign(data []byte) (r, s *big.Int) {
	return Sign(p, data)
}

func GenerateKey(curve elliptic.Curve, K *big.Int) *PrivateKey {
	priv := new(PrivateKey)
	priv.PublicKey.Curve = curve
	priv.PublicKey.X, priv.PublicKey.Y = curve.MulG(K)
	priv.D = K
	return priv
}

func Sign(priv *PrivateKey, msg []byte) (r, s *big.Int) {
	k := big.NewInt(4)

	r, _ = priv.Curve.MulG(k)

	z := MsgToInt(msg)
	dr := new(big.Int).Mul(priv.D, r)
	zdr := new(big.Int).Add(z, dr)

	s = ratMod(new(big.Rat).SetFrac(zdr, k), priv.Curve.Params().P)

	return r, s
}

func Verify(pub *PublicKey, r, s *big.Int, msg []byte) bool {
	P := pub.Curve.Params().P
	ss := ratMod(new(big.Rat).SetFrac(big.NewInt(1), s), P)

	z := MsgToInt(msg)
	zs := new(big.Int).Mul(z, ss)
	zs.Mod(zs, P)
	lx, ly := pub.Curve.MulG(zs)

	rs := new(big.Int).Mul(r, ss)
	rs.Mod(rs, P)
	rx, ry := pub.Curve.Mul(pub.X, pub.Y, rs)

	x, _ := pub.Curve.Add(lx, ly, rx, ry)

	return x.Cmp(r) == 0
}

func Rand() []byte {
	ku16 := uint16(rand.Uint64())
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.LittleEndian, ku16)
	return buf.Bytes()
}

func MsgToInt(msg []byte) *big.Int {
	md := sha1.New()
	md.Write(msg)
	hashed := md.Sum(nil)
	hashedInt := new(big.Int).SetBytes(hashed)
	return hashedInt
}

func ratMod(rat *big.Rat, p *big.Int) *big.Int {
	n := rat.Num()
	d := rat.Denom()
	if d.Cmp(big.NewInt(1)) == 0 {
		return new(big.Int).Mod(n, p)
	}

	fastPow := func(a, n *big.Int, p *big.Int) *big.Int {
		res := big.NewInt(1)
		for n.Int64() != 0 {
			if n.Int64()&1 == 1 {
				res.Mul(res, a).Mod(res, p)
			}
			a.Mul(a, a).Mod(a, p)
			n.Rsh(n, 1)
		}
		return res
	}
	// ( n·d^(p-2) ) % p
	res := fastPow(d, new(big.Int).Sub(p, big.NewInt(2)), p)
	res.Mul(res, n)
	res.Mod(res, p)
	return res
}
