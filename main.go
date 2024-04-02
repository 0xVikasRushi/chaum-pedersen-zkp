package main

import (
	"fmt"
	"math/big"
)

type Group struct {
	g *big.Int
	q *big.Int
}

// !  secret x and s random number
type Commitment struct {
	x *big.Int
	s *big.Int
}

type Proof struct {
	Y1 *big.Int
	Y2 *big.Int
	Z  *big.Int
}

func main() {
	var q = big.NewInt(10009) // ? prime number
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

	// ? secret
	x := big.NewInt(34)
	s := big.NewInt(300)

	group := Group{g, q}
	commitment := Commitment{x, s}
	proof := ProofGen(group, commitment, a, B)
	result := VerifyProof(group, proof, A, B, C, s)
	fmt.Println(result)
}

func ProofGen(group Group, commitment Commitment, a *big.Int, B *big.Int) Proof {
	Y1 := new(big.Int)
	Y2 := new(big.Int)

	Y1.Exp(group.g, commitment.x, group.q)
	Y2.Exp(B, commitment.x, group.q)

	Z := new(big.Int)
	Z = Z.Mul(a, commitment.s).Add(Z, commitment.x).Mod(Z, group.q)
	return Proof{Y1, Y2, Z}
}

func VerifyProof(group Group, proof Proof, A *big.Int, B *big.Int, C *big.Int, s *big.Int) bool {

	LHS1 := new(big.Int)
	RHS1 := new(big.Int)

	Y1 := proof.Y1
	Y2 := proof.Y2
	Z := proof.Z
	g := group.g
	q := group.q

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
