package ecc

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

func NewPointStr(x, y string, base int) Point {
	xb, ok := new(big.Int).SetString(x, base)
	if !ok {
		panic("set point.X failed")
	}
	yb, ok := new(big.Int).SetString(x, base)
	if !ok {
		panic("set point.Y failed")
	}
	return Point{xb, yb}
}

func (p Point) Equal(q Point) bool {
	return p.X.Cmp(q.X) == 0 && p.Y.Cmp(q.Y) == 0
}

func (p Point) IsO() bool {
	return p == PointO
}
