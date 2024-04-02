package main

import (
	"fmt"
	"math/big"
)

func main() {
	var q = big.NewInt(10009)
	// ? make cyclic group array
	// ? check for element which gcd with 1
	var g = big.NewInt(4)
	var a = big.NewInt(10)
	var b = big.NewInt(12)

	A := new(big.Int)
	B := new(big.Int)
	C := new(big.Int)
	temp := new(big.Int)

	A.Exp(g, a, q)
	B.Exp(g, b, q)
	C.Exp(g, temp.Mul(a, b), q)

	// fmt.Println(A, B, C)

	// ? proof generation part
	Y1 := new(big.Int)
	Y2 := new(big.Int)

	// ? secret
	x := big.NewInt(34)

	Y1.Exp(g, x, q)
	Y2.Exp(B, x, q)

	// ? choose random number s = 300

	s := big.NewInt(300)
	Z := new(big.Int)
	Z = Z.Mul(a, s).Add(Z, x).Mod(Z, q)

	// ? SEND Z
	result := verify(g, Z, q, A, B, C, Y1, Y2, s)
	fmt.Print(result)
}

func verify(g *big.Int, Z *big.Int, q *big.Int, A *big.Int, B *big.Int, C *big.Int, Y1 *big.Int, Y2 *big.Int, s *big.Int) bool {

	LHS1 := new(big.Int)
	RHS1 := new(big.Int)

	LHS1 = LHS1.Exp(g, Z, q)
	RHS1 = RHS1.Exp(A, s, nil).Mul(RHS1, Y1)
	RHS1 = RHS1.Mod(RHS1, q)

	LHS2 := new(big.Int)
	RHS2 := new(big.Int)

	LHS2 = LHS2.Exp(B, Z, q)
	RHS2 = RHS2.Exp(C, s, nil).Mul(RHS2, Y2)
	RHS2 = RHS2.Mod(RHS2, q)

	return LHS1.Cmp(RHS1) == 0 && LHS2.Cmp(RHS2) == 0
}
