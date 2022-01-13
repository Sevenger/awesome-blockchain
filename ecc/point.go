package ecc

import (
	"fmt"
	"math/big"
)

type Point struct {
	X, Y *big.Rat
}

var PointO = NewPoint(0, 0)

func (p Point) Print() {
	format := func(rat *big.Rat) string {
		if rat.Denom().Cmp(big.NewInt(1)) == 0 {
			return rat.Num().String()
		}
		return rat.String()
	}
	fmt.Printf("(%v, %v)\n", format(p.X), format(p.Y))
}

func NewPoint(x, y int64) Point {
	return Point{
		X: new(big.Rat).SetInt64(x),
		Y: new(big.Rat).SetInt64(y),
	}
}

func (p Point) Equal(q Point) bool {
	return p.X.Cmp(q.X) == 0 && p.Y.Cmp(q.Y) == 0
}

func (p Point) IsO() bool {
	return p == PointO
}
