package ECELGamal

import (
	"crypto/rand"
	"crypto/sha1"
	"elliptic/curve"
	"math/big"
)

// tutor see:https://asecuritysite.com/encryption/go_elgamal_ecc

type Point struct {
	X, Y *big.Int
}

type PublicKey struct {
	Curve elliptic.Curve
	A     Point
}

type PrivateKey struct {
	PublicKey
	a *big.Int
}

func (p *PrivateKey) Encrypt(data []byte) (K, C Point) {
	return Encrypt(p, data)
}

func (p *PrivateKey) Decrypt(K, C Point) (M Point) {
	return Decrypt(p, K, C)
}

// GenerateKey 生成私钥
func GenerateKey(curve elliptic.Curve, a *big.Int) *PrivateKey {
	params := curve.Params()

	priv := new(PrivateKey)
	priv.A.X, priv.A.Y = curve.Mul(params.Gx, params.Gy, a)
	priv.PublicKey.Curve = curve
	priv.a = a
	return priv
}

func Encrypt(priv *PrivateKey, data []byte) (K, C Point) {
	k, _ := rand.Int(rand.Reader, priv.Curve.Params().P)

	// K = k*G
	K.X, K.Y = priv.Curve.Mul(priv.Curve.Params().Gx, priv.Curve.Params().Gy, k)

	// KA = k*priv.A
	var kA Point
	kA.X, kA.Y = priv.Curve.Mul(priv.A.X, priv.A.Y, k)

	M := embedData(priv.Curve, data)

	// C = kA + M
	C.X, C.Y = priv.Curve.Add(M.X, M.Y, kA.X, kA.Y)

	return
}

func Decrypt(priv *PrivateKey, K, C Point) (M Point) {
	// S = k*priv.a
	var S Point
	S.X, S.Y = priv.Curve.Mul(K.X, K.Y, priv.a)
	S.Y.Mul(S.Y, big.NewInt(-1))

	M.X, M.Y = priv.Curve.Add(C.X, C.Y, S.X, S.Y)
	return
}

// hash data and embed to curve
func embedData(curve elliptic.Curve, data []byte) (M Point) {
	md := sha1.New()
	md.Write(data)
	hashed := md.Sum(nil)
	hashedInt := new(big.Int).SetBytes(hashed)

	M.X, M.Y = curve.Mul(curve.Params().Gx, curve.Params().Gy, hashedInt)
	return
}
