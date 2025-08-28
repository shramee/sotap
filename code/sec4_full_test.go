// DESCRIPTION
// ----------------------------------------------------------
// Section 4: Faster Field Extension Multiplication
// Prints benchmarks for Faster Field Extension Multiplication
// Prints the number of SCS and R1CS constraints required for,
// 1. 2-Pair Pairing
// 2. 2-Pair Miller Loop
// 3. Final Exponentiation (2-Pair Pairing cost - Miller Loop cost)
// 4. 𝔽p¹² multiplication
// This is done for both R1CS constraints and SCS constraints
// ----------------------------------------------------------
package sota_pairing

import (
	"fmt"
	"testing"

	"github.com/consensys/gnark/frontend"

	. "sota_pairing/benchmark"
	. "sota_pairing/fields_bn254"
	sw "sota_pairing/sw_bn254"
)

type Bench4_E12Mul struct {
	El MulEls[sw.GTEl]
}

func (circuit *Bench4_E12Mul) Define(api frontend.API) error {
	e := NewExt12(api)
	expected := e.Mul(&circuit.El.A, &circuit.El.B)
	e.AssertIsEqual(expected, &circuit.El.C)
	return nil
}

func (c *Bench4_E12Mul) Init() Benchmarkable {
	return &Bench4_E12Mul{
		El: RandomE12Mul(),
	}
}

type Bench4_MillerLoop struct {
	Pairs TwoPairs[sw.G1Affine, sw.G2Affine]
}

func (c *Bench4_MillerLoop) Init() Benchmarkable {
	return &Bench4_MillerLoop{
		Pairs: RandomPairs(),
	}
}

func (c *Bench4_MillerLoop) Define(api frontend.API) error {
	pairing, err := sw.NewPairing(api)
	if err != nil {
		return fmt.Errorf("new pairing: %w", err)
	}
	_, err = pairing.MillerLoop([]*sw.G1Affine{&c.Pairs.In1G1, &c.Pairs.In2G1}, []*sw.G2Affine{&c.Pairs.In1G2, &c.Pairs.In2G2})
	if err != nil {
		return fmt.Errorf("pair: %w", err)
	}
	return nil
}

type Bench4_Pairing struct {
	Pairs TwoPairs[sw.G1Affine, sw.G2Affine]
}

func (c *Bench4_Pairing) Init() Benchmarkable {
	return &Bench4_Pairing{
		Pairs: RandomPairs(),
	}
}

func (c *Bench4_Pairing) Define(api frontend.API) error {
	pairing, err := sw.NewPairing(api)
	if err != nil {
		return fmt.Errorf("new pairing: %w", err)
	}
	el, err := pairing.Pair([]*sw.G1Affine{&c.Pairs.In1G1, &c.Pairs.In2G1}, []*sw.G2Affine{&c.Pairs.In1G2, &c.Pairs.In2G2})
	if err != nil {
		return fmt.Errorf("pair: %w", err)
	}
	pairing.Ext12.AssertIsEqual(el, pairing.Ext12.One())
	return nil
}

// bench
func TestBench4Full(b *testing.T) {
	println("\n----------------------------------------")
	println("\n4 Faster Field Extension Multiplication")

	p_scs, p_r1cs := BenchmarkCircuit(&Bench4_Pairing{})
	ml_scs, ml_r1cs := BenchmarkCircuit(&Bench4_MillerLoop{})
	e12_scs, e12_r1cs := BenchmarkCircuit(&Bench4_E12Mul{})
	println("\n  1. 2-Pair Pairing")
	println("    SCS: ", p_scs, ", R1CS: ", p_r1cs)
	println("\n  2. 2-Pair Miller Loop")
	println("    SCS: ", ml_scs, ", R1CS: ", ml_r1cs)
	println("\n  3. Final Exponentiation")
	println("    SCS: ", p_scs-ml_scs, ", R1CS: ", p_r1cs-ml_r1cs)
	println("\n  4. 𝔽p¹² multiplication")
	println("    SCS: ", e12_scs, ", R1CS: ", e12_r1cs)
	println("\n----------------------------------------\n")
}
