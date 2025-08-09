package sota_pairing

import (
	"bytes"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/frontend/cs/scs"
)

type Benchmarkable interface {
	frontend.Circuit
	Init() Benchmarkable
}

func checkErr(err error, what string, b *testing.B) {
	if err != nil {
		b.Fatalf("Failed to %s: %v", what, err)
	}
}

func BenchmarkCircuit(circuit Benchmarkable, b *testing.B) {
	var buf bytes.Buffer
	witness := circuit.Init()

	w, err := frontend.NewWitness(witness, ecc.BN254.ScalarField())
	checkErr(err, "frontend.NewWitness", b)

	// Test SCS compilation and solving
	scs, err := frontend.Compile(ecc.BN254.ScalarField(), scs.NewBuilder, circuit.Init())
	checkErr(err, "scs compile", b)

	// Verify SCS solves correctly
	_, err = scs.Solve(w)
	checkErr(err, "solve scs", b)

	buf.Reset()
	_, err = scs.WriteTo(&buf)
	checkErr(err, "scs.WriteTo", b)
	b.Logf("SCS - Size: %d bytes, Constraints: %d, Instructions: %d",
		buf.Len(), scs.GetNbConstraints(), scs.GetNbInstructions())

	// Test R1CS compilation and solving
	r1cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, circuit.Init())
	checkErr(err, "r1cs compile", b)

	// Verify R1CS solves correctly
	_, err = r1cs.Solve(w)
	checkErr(err, "solve r1cs", b)

	buf.Reset()
	_, err = r1cs.WriteTo(&buf)
	checkErr(err, "r1cs.WriteTo", b)

	b.Logf("R1CS - Size: %d bytes, Constraints: %d, Instructions: %d",
		buf.Len(), r1cs.GetNbConstraints(), r1cs.GetNbInstructions())
}
