package sota_pairing

import (
	"fmt"
	"testing"

	"github.com/consensys/gnark/frontend"

	. "sota_pairing/sw_bn254"
	sw_towered "sota_pairing/sw_bn254_towered"
)

type Bench3_Pair_PairingsInR1CS struct {
	In1G1 sw_towered.G1Affine
	In2G1 sw_towered.G1Affine
	In1G2 sw_towered.G2Affine
	In2G2 sw_towered.G2Affine
}

func (c *Bench3_Pair_PairingsInR1CS) Define(api frontend.API) error {
	pairing, err := sw_towered.NewPairing(api)
	if err != nil {
		return fmt.Errorf("new pairing: %w", err)
	}
	el, err := pairing.Pair([]*sw_towered.G1Affine{&c.In1G1, &c.In2G1}, []*sw_towered.G2Affine{&c.In1G2, &c.In2G2})
	if err != nil {
		return fmt.Errorf("pair: %w", err)
	}
	pairing.Ext12.IsEqual(el, pairing.Ext12.One())
	return nil
}

func (c *Bench3_Pair_PairingsInR1CS) Init() Benchmarkable {
	p1, p2, q1, q2 := RandomG1G2AffinesPair()
	return &Bench3_Pair_PairingsInR1CS{
		In1G1: sw_towered.NewG1Affine(p1),
		In1G2: sw_towered.NewG2Affine(q1),
		In2G1: sw_towered.NewG1Affine(p2),
		In2G2: sw_towered.NewG2Affine(q2),
	}
}

type Bench4_Pair_FasterFieldExtensionMuls struct {
	In1G1 G1Affine
	In2G1 G1Affine
	In1G2 G2Affine
	In2G2 G2Affine
}

func (c *Bench4_Pair_FasterFieldExtensionMuls) Define(api frontend.API) error {
	pairing, err := NewPairing(api)
	if err != nil {
		return fmt.Errorf("new pairing: %w", err)
	}
	el, err := pairing.Pair([]*G1Affine{&c.In1G1, &c.In2G1}, []*G2Affine{&c.In1G2, &c.In2G2})
	if err != nil {
		return fmt.Errorf("pair: %w", err)
	}
	pairing.Ext12.IsEqual(el, pairing.Ext12.One())
	return nil
}

func (c *Bench4_Pair_FasterFieldExtensionMuls) Init() Benchmarkable {
	p1, p2, q1, q2 := RandomG1G2AffinesPair()
	return &Bench4_Pair_FasterFieldExtensionMuls{
		In1G1: NewG1Affine(p1),
		In1G2: NewG2Affine(q1),
		In2G1: NewG1Affine(p2),
		In2G2: NewG2Affine(q2),
	}
}

type Bench5_Pair_EliminatingFinalExponentiation struct {
	In1G1 G1Affine
	In2G1 G1Affine
	In1G2 G2Affine
	In2G2 G2Affine
}

func (c *Bench5_Pair_EliminatingFinalExponentiation) Define(api frontend.API) error {
	pairing, err := NewPairing(api)
	if err != nil {
		return fmt.Errorf("new pairing: %w", err)
	}
	err = pairing.PairingCheck([]*G1Affine{&c.In1G1, &c.In2G1}, []*G2Affine{&c.In1G2, &c.In2G2})
	if err != nil {
		return fmt.Errorf("pair: %w", err)
	}
	return nil
}

func (c *Bench5_Pair_EliminatingFinalExponentiation) Init() Benchmarkable {
	p1, p2, q1, q2 := RandomG1G2AffinesPair()
	return &Bench5_Pair_EliminatingFinalExponentiation{
		In1G1: NewG1Affine(p1),
		In1G2: NewG2Affine(q1),
		In2G1: NewG1Affine(p2),
		In2G2: NewG2Affine(q2),
	}
}

// bench
func BenchmarkPairing(b *testing.B) {
	BenchmarkCircuit(&Bench3_Pair_PairingsInR1CS{}, b)
	BenchmarkCircuit(&Bench4_Pair_FasterFieldExtensionMuls{}, b)
	BenchmarkCircuit(&Bench5_Pair_EliminatingFinalExponentiation{}, b)
}
