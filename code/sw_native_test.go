package sota_pairing

import (
	"fmt"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	bls12377 "github.com/consensys/gnark-crypto/ecc/bls12-377"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/frontend/cs/scs"
	"github.com/consensys/gnark/std/algebra/native/sw_bls12377"
	"github.com/consensys/gnark/test"

	. "sota_pairing/benchmark"
)

type pairingBLS377 struct {
	P1, P2 sw_bls12377.G1Affine
	Q1, Q2 sw_bls12377.G2Affine
}

func (circuit *pairingBLS377) Define(api frontend.API) error {

	res, err := sw_bls12377.Pair(api, []sw_bls12377.G1Affine{circuit.P1, circuit.P2}, []sw_bls12377.G2Affine{circuit.Q1, circuit.Q2})

	_ = res

	// var one sw_bls12377.GT
	// one.SetOne()
	// api.AssertIsEqual(res, one)

	if err != nil {
		return fmt.Errorf("pair: %w", err)
	}

	return nil
}

type millerloopBLS377 struct {
	P1, P2 sw_bls12377.G1Affine
	Q1, Q2 sw_bls12377.G2Affine
}

func (circuit *millerloopBLS377) Define(api frontend.API) error {
	res, err := sw_bls12377.MillerLoop(api, []sw_bls12377.G1Affine{circuit.P1, circuit.P2}, []sw_bls12377.G2Affine{circuit.Q1, circuit.Q2})

	_ = res

	if err != nil {
		return fmt.Errorf("pair: %w", err)
	}

	return nil
}

type pairingCheckBLS377 struct {
	P1, P2 sw_bls12377.G1Affine
	Q1, Q2 sw_bls12377.G2Affine
}

func (circuit *pairingCheckBLS377) Define(api frontend.API) error {

	err := sw_bls12377.PairingCheck(api, []sw_bls12377.G1Affine{circuit.P1, circuit.P2}, []sw_bls12377.G2Affine{circuit.Q1, circuit.Q2})

	if err != nil {
		return fmt.Errorf("pair: %w", err)
	}

	return nil
}

func pairingCheckData() (P [2]bls12377.G1Affine, Q [2]bls12377.G2Affine) {
	// e(a,2b) * e(-2a,b) == 1
	_, _, P[0], Q[0] = bls12377.Generators()
	P[1].Double(&P[0]).Neg(&P[1])
	Q[1].Set(&Q[0])
	Q[0].Double(&Q[0])

	return
}

func TestNativePairing1(t *testing.T) {

	// pairing test data
	P, Q := pairingCheckData()
	assert := test.NewAssert(t)

	witness1 := pairingBLS377{
		P1: sw_bls12377.NewG1Affine(P[0]),
		P2: sw_bls12377.NewG1Affine(P[1]),
		Q1: sw_bls12377.NewG2Affine(Q[0]),
		Q2: sw_bls12377.NewG2Affine(Q[1]),
	}
	assert.CheckCircuit(&pairingBLS377{}, test.WithValidAssignment(&witness1), test.WithCurves(ecc.BW6_761), test.NoProverChecks())

	// Test SCS compilation and solving
	scs1, _ := frontend.Compile(ecc.BW6_761.ScalarField(), scs.NewBuilder, &witness1)
	r1cs1, _ := frontend.Compile(ecc.BW6_761.ScalarField(), r1cs.NewBuilder, &witness1)

	println("---------------------------------")
	println("Pairing with final exponentiation")
	println("scs constraints:", scs1.GetNbConstraints())
	println("r1cs constraints:", r1cs1.GetNbConstraints())

	witness := pairingCheckBLS377{
		P1: witness1.P1,
		P2: witness1.P2,
		Q1: witness1.Q1,
		Q2: witness1.Q2,
	}
	assert.CheckCircuit(&pairingCheckBLS377{}, test.WithValidAssignment(&witness), test.WithCurves(ecc.BW6_761), test.NoProverChecks())

	// Test SCS compilation and solving
	scs2, _ := frontend.Compile(ecc.BW6_761.ScalarField(), scs.NewBuilder, &witness)
	r1cs2, _ := frontend.Compile(ecc.BW6_761.ScalarField(), r1cs.NewBuilder, &witness)

	println("---------------------------------")
	println("Pairing without final exponentiation")
	println("scs constraints:", scs2.GetNbConstraints())
	println("r1cs constraints:", r1cs2.GetNbConstraints())

}
func TestNativePairing(t *testing.T) {

	// pairing test data
	P, Q := pairingCheckData()
	assert := test.NewAssert(t)

	witness0 := millerloopBLS377{
		P1: sw_bls12377.NewG1Affine(P[0]),
		P2: sw_bls12377.NewG1Affine(P[1]),
		Q1: sw_bls12377.NewG2Affine(Q[0]),
		Q2: sw_bls12377.NewG2Affine(Q[1]),
	}
	assert.CheckCircuit(&millerloopBLS377{}, test.WithValidAssignment(&witness0), test.WithCurves(ecc.BW6_761), test.NoProverChecks())

	// Test SCS compilation and solving
	scs0, _ := frontend.Compile(ecc.BW6_761.ScalarField(), scs.NewBuilder, &witness0)
	r1cs0, _ := frontend.Compile(ecc.BW6_761.ScalarField(), r1cs.NewBuilder, &witness0)

	witness1 := pairingBLS377{
		P1: sw_bls12377.NewG1Affine(P[0]),
		P2: sw_bls12377.NewG1Affine(P[1]),
		Q1: sw_bls12377.NewG2Affine(Q[0]),
		Q2: sw_bls12377.NewG2Affine(Q[1]),
	}
	assert.CheckCircuit(&pairingBLS377{}, test.WithValidAssignment(&witness1), test.WithCurves(ecc.BW6_761), test.NoProverChecks())

	// Test SCS compilation and solving
	scs1, _ := frontend.Compile(ecc.BW6_761.ScalarField(), scs.NewBuilder, &witness1)
	r1cs1, _ := frontend.Compile(ecc.BW6_761.ScalarField(), r1cs.NewBuilder, &witness1)

	witness := pairingCheckBLS377{
		P1: sw_bls12377.NewG1Affine(P[0]),
		P2: sw_bls12377.NewG1Affine(P[1]),
		Q1: sw_bls12377.NewG2Affine(Q[0]),
		Q2: sw_bls12377.NewG2Affine(Q[1]),
	}

	assert.CheckCircuit(&pairingCheckBLS377{}, test.WithValidAssignment(&witness), test.WithCurves(ecc.BW6_761), test.NoProverChecks())

	// Test SCS compilation and solving
	scs2, _ := frontend.Compile(ecc.BW6_761.ScalarField(), scs.NewBuilder, &witness)
	r1cs2, _ := frontend.Compile(ecc.BW6_761.ScalarField(), r1cs.NewBuilder, &witness)

	miller_scs := scs0.GetNbConstraints()
	miller_r1cs := r1cs0.GetNbConstraints()
	classic_scs := scs1.GetNbConstraints()
	classic_r1cs := r1cs1.GetNbConstraints()
	efx_scs := scs2.GetNbConstraints()
	efx_r1cs := r1cs2.GetNbConstraints()

	println("---------------------------------")
	// println("Miller loop")
	// println("scs constraints:", scs_miller)
	// println("r1cs constraints:", r1cs_miller)

	println("Classic Pairing")
	println("whole pairing scs constraints:", classic_scs)
	println("whole pairing r1cs constraints:", classic_r1cs)
	println("net final exponentiation scs constraints:", classic_scs-miller_scs)
	println("net final exponentiation r1cs constraints:", classic_r1cs-miller_r1cs)

	println("Eliminiating Final exponentiation")
	println("whole pairing scs constraints:", efx_scs)
	println("whole pairing r1cs constraints:", efx_r1cs)
	println("net final exponentiation scs constraints:", efx_scs-miller_scs)
	println("net final exponentiation r1cs constraints:", efx_r1cs-miller_r1cs)

}

// bench
func BenchmarkNativePairing(b *testing.B) {
	fmt.Printf("5: Eliminating Final Exponentiation\n%s\n", BenchmarkCircuitStr(&Bench5_Pairing_ElFX{}))
	// fmt.Printf("6: Miller loop step polynomial\n%s\n", BenchmarkCircuitStr(&Bench6_Pairing_MLSBatch{}))
}
