package sota_pairing

import (
	"crypto/rand"

	fields "sota_pairing/fields_bn254"
	fields_towered "sota_pairing/fields_bn254_towered"
	sw_towered "sota_pairing/sw_bn254_towered"

	. "sota_pairing/sw_bn254"

	"github.com/consensys/gnark-crypto/ecc/bn254"
)

type TwoPairs[T1, T2 any] struct {
	In1G1 T1
	In2G1 T1
	In1G2 T2
	In2G2 T2
}

type MulEls[T1 any] struct {
	A, B, C T1
}

func RandomG1G2Affines() (p bn254.G1Affine, q bn254.G2Affine, err error) {
	_, _, G1AffGen, G2AffGen := bn254.Generators()
	mod := bn254.ID.ScalarField()
	s1, err := rand.Int(rand.Reader, mod)
	if err != nil {
		return p, q, err
	}
	s2, err := rand.Int(rand.Reader, mod)
	if err != nil {
		return p, q, err
	}
	p.ScalarMultiplication(&G1AffGen, s1)
	q.ScalarMultiplication(&G2AffGen, s2)
	return
}

func RandomG1G2AffinesPair() (p1, p2 bn254.G1Affine, q1, q2 bn254.G2Affine) {
	p1, q1, err := RandomG1G2Affines()
	if err != nil {
		panic(err)
	}
	p2.Double(&p1).Neg(&p2)
	q2.Set(&q1)
	q1.Double(&q1)
	return
}

func RandomE12Els() (a, b, c bn254.E12) {
	a.SetRandom()
	b.SetRandom()
	c.Mul(&a, &b)
	return
}

func RandomToweredPairs() TwoPairs[sw_towered.G1Affine, sw_towered.G2Affine] {
	p1, p2, q1, q2 := RandomG1G2AffinesPair()
	return TwoPairs[sw_towered.G1Affine, sw_towered.G2Affine]{
		In1G1: sw_towered.NewG1Affine(p1),
		In2G1: sw_towered.NewG1Affine(p2),
		In1G2: sw_towered.NewG2Affine(q1),
		In2G2: sw_towered.NewG2Affine(q2),
	}
}

func RandomPairs() TwoPairs[G1Affine, G2Affine] {
	p1, p2, q1, q2 := RandomG1G2AffinesPair()
	return TwoPairs[G1Affine, G2Affine]{
		In1G1: NewG1Affine(p1),
		In2G1: NewG1Affine(p2),
		In1G2: NewG2Affine(q1),
		In2G2: NewG2Affine(q2),
	}
}

func RandomToweredE12Mul() MulEls[fields_towered.E12] {
	a, b, c := RandomE12Els()
	return MulEls[fields_towered.E12]{
		A: fields_towered.FromE12(&a),
		B: fields_towered.FromE12(&b),
		C: fields_towered.FromE12(&c),
	}
}

func RandomE12Mul() MulEls[GTEl] {
	a, b, c := RandomE12Els()
	return MulEls[GTEl]{
		A: fields.FromE12(&a),
		B: fields.FromE12(&b),
		C: fields.FromE12(&c),
	}
}
