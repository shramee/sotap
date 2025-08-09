package sota_pairing

import (
	"fmt"
	"testing"

	"github.com/consensys/gnark/frontend"

	. "sota_pairing/sw_bn254"
	sw_towered "sota_pairing/sw_bn254_towered"
)

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

func (c *Bench3_Pairing_PiR1) Init() Benchmarkable {
	return &Bench3_Pairing_PiR1{
		Pairs: RandomToweredPairs(),
	}
}

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

func (c *Bench4_Pairing_FXFM) Init() Benchmarkable {
	return &Bench4_Pairing_FXFM{
		Pairs: RandomPairs(),
	}
}

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

func (c *Bench5_Pairing_ElFX) Init() Benchmarkable {
	return &Bench5_Pairing_ElFX{
		Pairs: RandomPairs(),
	}
}

// bench
func BenchmarkPairing(b *testing.B) {
	fmt.Printf("%s\n", BenchmarkCircuitStr(&Bench3_Pairing_PiR1{}, b))
	fmt.Printf("%s\n", BenchmarkCircuitStr(&Bench4_Pairing_FXFM{}, b))
	fmt.Printf("%s\n", BenchmarkCircuitStr(&Bench5_Pairing_ElFX{}, b))
}
