package ecc2

import (
	"fmt"
	"math/big"
)

type Point struct {
	X, Y *big.Int
}

var PointO = NewPoint64(0, 0)

func (p Point) Print() {
	fmt.Printf("(%v, %v)\n", p.X.String(), p.Y.String())
}

func NewPoint(x, y *big.Int) Point {
	return Point{x, y}
}

func NewPoint64(x, y int64) Point {
	return Point{big.NewInt(x), big.NewInt(y)}
}

func NewPointStr(x, y string, base int) (Point, error) {
	xb, ok := new(big.Int).SetString(x, base)
	if !ok {
		return Point{}, fmt.Errorf(convertBigIntErr, "point.X")
	}
	yb, ok := new(big.Int).SetString(y, base)
	if !ok {
		return Point{}, fmt.Errorf(convertBigIntErr, "point.Y")
	}
	return Point{xb, yb}, nil
}

func (p Point) Equal(q Point) bool {
	return IsEqual(p.X, q.X) && IsEqual(p.Y, q.Y)
}

func (p Point) IsO() bool {
	return p.Equal(PointO)
}
