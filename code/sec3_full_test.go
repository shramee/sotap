package sota_pairing

import (
	"fmt"
	"testing"

	"github.com/consensys/gnark/frontend"

	. "sota_pairing/benchmark"
	. "sota_pairing/fields_bn254_towered"
	sw "sota_pairing/sw_bn254_towered"
)

type Bench3_E12Mul struct {
	El MulEls[sw.GTEl]
}

func (circuit *Bench3_E12Mul) Define(api frontend.API) error {
	e := NewExt12(api)
	expected := e.Mul(&circuit.El.A, &circuit.El.B)
	e.AssertIsEqual(expected, &circuit.El.C)
	return nil
}

func (c *Bench3_E12Mul) Init() Benchmarkable {
	return &Bench3_E12Mul{
		El: RandomToweredE12Mul(),
	}
}

type Bench3_MillerLoop struct {
	Pairs TwoPairs[sw.G1Affine, sw.G2Affine]
}

func (c *Bench3_MillerLoop) Init() Benchmarkable {
	return &Bench3_MillerLoop{
		Pairs: RandomToweredPairs(),
	}
}

func (c *Bench3_MillerLoop) Define(api frontend.API) error {
	pairing, err := sw.NewPairing(api)
	if err != nil {
		return fmt.Errorf("new pairing: %w", err)
	}
	el, err := pairing.MillerLoop([]*sw.G1Affine{&c.Pairs.In1G1, &c.Pairs.In2G1}, []*sw.G2Affine{&c.Pairs.In1G2, &c.Pairs.In2G2})
	if err != nil {
		return fmt.Errorf("pair: %w", err)
	}
	pairing.Ext12.IsEqual(el, pairing.Ext12.One())
	return nil
}

type Bench3_Pairing struct {
	Pairs TwoPairs[sw.G1Affine, sw.G2Affine]
}

func (c *Bench3_Pairing) Init() Benchmarkable {
	return &Bench3_Pairing{
		Pairs: RandomToweredPairs(),
	}
}

func (c *Bench3_Pairing) Define(api frontend.API) error {
	pairing, err := sw.NewPairing(api)
	if err != nil {
		return fmt.Errorf("new pairing: %w", err)
	}
	el, err := pairing.Pair([]*sw.G1Affine{&c.Pairs.In1G1, &c.Pairs.In2G1}, []*sw.G2Affine{&c.Pairs.In1G2, &c.Pairs.In2G2})
	if err != nil {
		return fmt.Errorf("pair: %w", err)
	}
	pairing.Ext12.IsEqual(el, pairing.Ext12.One())
	return nil
}

// bench
func TestBench3Full(b *testing.T) {
	fmt.Printf("\n\n3: Pairing in R1CS\n")

	p_scs, p_r1cs := BenchmarkCircuit(&Bench3_Pairing{})
	ml_scs, ml_r1cs := BenchmarkCircuit(&Bench3_MillerLoop{})
	fmt.Printf("\n\nPairing\nSCS: %d, R1CS: %d", p_scs, p_r1cs)
	fmt.Printf("\n\nMiller Loop\nSCS: %d, R1CS: %d", ml_scs, ml_r1cs)
	fmt.Printf("\n\nFinal Exponentiation\nSCS: %d, R1CS: %d", p_scs-ml_scs, p_r1cs-ml_r1cs)
	scs, r1cs := BenchmarkCircuit(&Bench3_E12Mul{})
	fmt.Printf("\n\nFp12 Mul\nSCS: %d, R1CS: %d", scs, r1cs)
}
