package sota_pairing

import (
	"crypto/rand"

	"github.com/consensys/gnark-crypto/ecc/bn254"
)

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
