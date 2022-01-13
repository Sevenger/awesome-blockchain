package main

import "math/big"

func Pow(num *big.Int, n uint) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	} else if n%2 == 0 {
		return new(big.Int).Mul(Pow(num, n/2), Pow(num, n/2))
	} else {
		return new(big.Int).Mul(Pow(num, n-1), num)
	}
}

func main() {
	println(Pow(big.NewInt(7), 100).String())
}
