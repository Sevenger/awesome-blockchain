package ecc

import (
	"fmt"
	"math/big"
)

type Point struct {
	X, Y *big.Rat
}

func (p Point) Print() {
	fmt.Printf("(%v, %v)\n", p.X, p.Y)
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
