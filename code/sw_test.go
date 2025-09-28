package sota_pairing

import (
	"fmt"
	"testing"

	"github.com/consensys/gnark/frontend"

	. "sota_pairing/benchmark"
	. "sota_pairing/sw_bn254"
	sw_towered "sota_pairing/sw_bn254_towered"
)

// DESCRIPTION
// ----------------------------------------------------------
// For each section we find number of constraints used for,
// 1. Whole Pairing
// 2. Miller Loop
// 3. Final Exponentiation (Whole Pairing cost - Miller Loop cost)
// 4. 𝔽p¹² multiplication
// This is done for both R1CS constraints and SCS constraints
// ----------------------------------------------------------
// Tests are included for sections,
// Section 3: Pairing in R1CS
// Section 4: Faster Field Extension Multiplication
// Section 5: Eliminating Final Exponentiation
// Section 6: Miller loop step polynomial
// ----------------------------------------------------------

// --------------------------
// Section 3: Pairing in R1CS
// --------------------------

type Bench3_Pairing_PiR1 struct {
	Pairs TwoPairs[sw_towered.G1Affine, sw_towered.G2Affine]
}

func (c *Bench3_Pairing_PiR1) Define(api frontend.API) error {
	pairing, err := sw_towered.NewPairing(api)
	if err != nil {
		return fmt.Errorf("new pairing: %w", err)
	}
	el, err := pairing.Pair([]*sw_towered.G1Affine{&c.Pairs.In1G1, &c.Pairs.In2G1}, []*sw_towered.G2Affine{&c.Pairs.In1G2, &c.Pairs.In2G2})
	if err != nil {
		return fmt.Errorf("pair: %w", err)
	}
	pairing.Ext12.IsEqual(el, pairing.Ext12.One())
	return nil
}

func (c *Bench3_Pairing_PiR1) Init() *Bench3_Pairing_PiR1 {
	return &Bench3_Pairing_PiR1{
		Pairs: RandomToweredPairs(),
	}
}

// ------------------------------------------------
// Section 4: Faster Field Extension Multiplication
// ------------------------------------------------

type Bench4_Pairing_FXFM struct {
	Pairs TwoPairs[G1Affine, G2Affine]
}

func (c *Bench4_Pairing_FXFM) Define(api frontend.API) error {
	pairing, err := NewPairing(api)
	if err != nil {
		return fmt.Errorf("new pairing: %w", err)
	}
	el, err := pairing.Pair([]*G1Affine{&c.Pairs.In1G1, &c.Pairs.In2G1}, []*G2Affine{&c.Pairs.In1G2, &c.Pairs.In2G2})
	if err != nil {
		return fmt.Errorf("pair: %w", err)
	}
	pairing.Ext12.IsEqual(el, pairing.Ext12.One())
	return nil
}

func (c *Bench4_Pairing_FXFM) Init() *Bench4_Pairing_FXFM {
	return &Bench4_Pairing_FXFM{
		Pairs: RandomPairs(),
	}
}

// -------------------------------------------
// Section 5: Eliminating Final Exponentiation
// -------------------------------------------

type Bench5_Pairing_ElFX struct {
	Pairs TwoPairs[G1Affine, G2Affine]
}

func (c *Bench5_Pairing_ElFX) Define(api frontend.API) error {
	pairing, err := NewPairing(api)
	if err != nil {
		return fmt.Errorf("new pairing: %w", err)
	}
	err = pairing.PairingCheck([]*G1Affine{&c.Pairs.In1G1, &c.Pairs.In2G1}, []*G2Affine{&c.Pairs.In1G2, &c.Pairs.In2G2})
	if err != nil {
		return fmt.Errorf("pair: %w", err)
	}
	return nil
}

func (c *Bench5_Pairing_ElFX) Init() *Bench5_Pairing_ElFX {
	return &Bench5_Pairing_ElFX{
		Pairs: RandomPairs(),
	}
}

// --------------------------------------
// Section 6: Miller loop step polynomial
// --------------------------------------

type Bench6_Pairing_MLSBatch struct {
	Pairs TwoPairs[G1Affine, G2Affine]
}

func (c *Bench6_Pairing_MLSBatch) Define(api frontend.API) error {
	pairing, err := NewPairing(api)
	if err != nil {
		return fmt.Errorf("new pairing: %w", err)
	}
	err = pairing.PairingCheckBatched([]*G1Affine{&c.Pairs.In1G1, &c.Pairs.In2G1}, []*G2Affine{&c.Pairs.In1G2, &c.Pairs.In2G2})
	if err != nil {
		return fmt.Errorf("pair: %w", err)
	}
	return nil
}

func (c *Bench6_Pairing_MLSBatch) Init() *Bench6_Pairing_MLSBatch {
	return &Bench6_Pairing_MLSBatch{
		Pairs: RandomPairs(),
	}
}

// bench
func BenchmarkPairing(b *testing.B) {
	fmt.Printf("3: Pairing in R1CS\n%s\n", BenchmarkCircuitStr(&Bench3_Pairing_PiR1{}))
	fmt.Printf("4: Faster Field Extension Multiplication\n%s\n", BenchmarkCircuitStr(&Bench4_Pairing_FXFM{}))
	fmt.Printf("5: Eliminating Final Exponentiation\n%s\n", BenchmarkCircuitStr(&Bench5_Pairing_ElFX{}))
	fmt.Printf("6: Miller loop step polynomial\n%s\n", BenchmarkCircuitStr(&Bench6_Pairing_MLSBatch{}))
}
