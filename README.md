<h1>Efficient Verification of Pairing Computation<br />
<small>Schwartz-Zippel Lemma for Miller Loop</small>
</h1>

<p align="center">
Shramee Srivastav
<br/>
MIST.cash — FOCBB
</p>
<p align="center">
<b>Abstract</b>
</p>

In this work, we consider the setting where elliptic curve pairings need to be performed within another proving system. This appears in proof recursion where a proof verifies other proofs, or when pairing-based protocols like BLS (Boneh-Lynn-Shacham) signatures or KZG commitments are used within a proof. We show a construction that simultaneously reduces arithmetic circuit depth and communication complexity.

We present \emph{a two-round public-coin interactive proof system} that consolidates the Miller loop iterations and final exponentiation into unified polynomial relations, verified as polynomial identities over $\mathbb{F}_q[x]$. This construction significantly reduces computational overhead across R1CS, Plonkish and AIR-based systems. We prove knowledge soundness $\delta_s \leq 2^{-245}$ for a three-pair optimal ate multi-pairing on BN254 curve.

As our \emph{second contribution} we present a generalised construction of our core technique, parameterised by an arbitrary modulus polynomial, designed for easy application across various protocols and schemes beyond pairings that utilize polynomial ring multiplications.

We demonstrate the practical efficacy of our construction across R1CS, Plonkish, and AIR arithmetisations on BN254 and BLS12-381, with implementations in gnark---the open-source zk-SNARK ecosystem---and in Garaga, the standard library for pairing-based ZK proof verification on Cairo, a leading AIR-based CPU architecture.

## Structure

- `paper/` - LaTeX source files for the paper
- `bench/` - Benchmarking code

## Building the Paper

Multiple passes required 

```bash
cd paper
pdflatex main.tex
bibtex main
pdflatex main.tex
pdflatex main.tex
```

## Running Experiments

```bash
cd experiments
go run main.go
```

## Results

Results are available in [results/](./results) directory.

## Dependencies

- LaTeX distribution (TeXLive, MiKTeX, etc.)
- Go 1.19+ 
- gnark library

## Benchmarks Garaga
```
Before implementation:

| circuit                    | MULMOD | ADDMOD | POSEIDON | ~cycles |
| -------------------------- | ------ | ------ | -------- | ------- |
| Miller n=3 BLS12_381       | 11356  | 11608  | 3088     | 198070  |
| MultiPairing n=3 BLS12_381 | 16484  | 20669  | 5421     | 325757  |
| Miller n=3 BN254           | 14456  | 14463  | 3758     | 236382  |
| MultiPairing n=3 BN254     | 19142  | 21686  | 5689     | 338678  |

New results:

| circuit                    | MULMOD | ADDMOD | POSEIDON | ~cycles |
| -------------------------- | ------ | ------ | -------- | ------- |
| Miller n=3 BLS12_381       | 6164   | 6364   | 834      | 91528   |
| MultiPairing n=3 BLS12_381 | 11287  | 15420  | 3167     | 219155  |
| Miller n=3 BN254           | 7975   | 7924   | 876      | 110666  |
| MultiPairing n=3 BN254     | 12656  | 15142  | 2807     | 212902  |
```

## Benchmarks Gnark

```
----------------------------------------
        GNARK BASE BENCHMARK
----------------------------------------

goos: darwin
goarch: arm64
pkg: github.com/consensys/gnark/std/algebra/emulated/sw_bn254
cpu: Apple M3 Pro
BenchmarkGroth16Simulation
BenchmarkGroth16Simulation/compile_scs
BenchmarkGroth16Simulation/compile_scs-12                      2         775639125 ns/op        1767674616 B/op  9739433 allocs/op
    g16_simulation_test.go:178: nb commitments: 739198, scs size: 50245511 (bytes), nb constraints 1890771, nbInstructions: 1961583
BenchmarkGroth16Simulation/solve_scs
BenchmarkGroth16Simulation/solve_scs-12                        3         479610014 ns/op        510358354 B/op   4762353 allocs/op
BenchmarkGroth16Simulation/compile_r1cs
BenchmarkGroth16Simulation/compile_r1cs-12                     2         780398875 ns/op        2091155060 B/op 18979147 allocs/op
    g16_simulation_test.go:204: nb commitments: 739198, r1cs size: 40080665 (bytes), nb constraints 588013, nbInstructions: 658825
BenchmarkGroth16Simulation/solve_r1cs
BenchmarkGroth16Simulation/solve_r1cs-12                       4         320837448 ns/op        330537186 B/op   3772976 allocs/op
PASS
ok      github.com/consensys/gnark/std/algebra/emulated/sw_bn254        10.938s


----------------------------------------
         THIS WORK BENCHMARK
----------------------------------------

goos: darwin
goarch: arm64
pkg: github.com/consensys/gnark/std/algebra/emulated/sw_bn254
cpu: Apple M3 Pro
BenchmarkGroth16Simulation
BenchmarkGroth16Simulation/compile_scs
BenchmarkGroth16Simulation/compile_scs-12                      3         406713917 ns/op        1029059832 B/op  6041293 allocs/op
    g16_simulation_test.go:178: nb commitments: 295946, scs size: 27893950 (bytes), nb constraints 1158194, nbInstructions: 1196936
BenchmarkGroth16Simulation/solve_scs
BenchmarkGroth16Simulation/solve_scs-12                        3         356417264 ns/op        376566120 B/op   2466802 allocs/op
BenchmarkGroth16Simulation/compile_r1cs
BenchmarkGroth16Simulation/compile_r1cs-12                     3         423090320 ns/op        1190704520 B/op 10771092 allocs/op
    g16_simulation_test.go:204: nb commitments: 295946, r1cs size: 21254062 (bytes), nb constraints 353881, nbInstructions: 392625
BenchmarkGroth16Simulation/solve_r1cs
BenchmarkGroth16Simulation/solve_r1cs-12                       5         214731150 ns/op        158506558 B/op   1501978 allocs/op
PASS
ok      github.com/consensys/gnark/std/algebra/emulated/sw_bn254        10.117s

----------------------------------------
       Conducted on 2026-04-15
```