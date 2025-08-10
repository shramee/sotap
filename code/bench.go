package sota_pairing

import (
	"bytes"
	"fmt"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/frontend/cs/scs"
)

type Benchmarkable interface {
	frontend.Circuit
	Init() Benchmarkable
}

func checkErr(err error, what string) {
	if err != nil {
		panic(fmt.Sprintf("Failed to %s: %v", what, err))
	}
}

func BenchmarkCircuit(circuit Benchmarkable) (int, int) {
	var buf bytes.Buffer
	witness := circuit.Init()

	w, err := frontend.NewWitness(witness, ecc.BN254.ScalarField())
	checkErr(err, "frontend.NewWitness")

	// Test SCS compilation and solving
	scs, err := frontend.Compile(ecc.BN254.ScalarField(), scs.NewBuilder, circuit.Init())
	checkErr(err, "scs compile")

	// Verify SCS solves correctly
	_, err = scs.Solve(w)
	checkErr(err, "solve scs")

	buf.Reset()
	_, err = scs.WriteTo(&buf)
	checkErr(err, "scs.WriteTo")

	// Test R1CS compilation and solving
	r1cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, circuit.Init())
	checkErr(err, "r1cs compile")

	// Verify R1CS solves correctly
	_, err = r1cs.Solve(w)
	checkErr(err, "solve r1cs")

	buf.Reset()
	_, err = r1cs.WriteTo(&buf)
	checkErr(err, "r1cs.WriteTo")

	return scs.GetNbConstraints(), r1cs.GetNbConstraints()
}

func BenchmarkCircuitStr(circuit Benchmarkable) string {
	scs, r1cs := BenchmarkCircuit(circuit)
	return fmt.Sprintf("SCS: %d, R1CS: %d", scs, r1cs)
}
