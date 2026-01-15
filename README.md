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

In this paper, we examine recent advances in the efficient proof of correctness for ellip-
tic curve pairing computations. Rather than focusing on the raw computation of pairings
themselves, our work centers on practical, sound, and high-performance techniques for prov-
ing that a pairing computation has been performed correctly—an essential step in modern
zero-knowledge proof (ZKP) systems. These techniques are foundational for SNARK-based
proof verification, especially in applications where succinctness and scalability matter, such
as public blockchains.
Our target use case involves zero-knowledge statements for moderately sized problems
(around one million constraints). So we focus our benchmarks on the Groth16 proof sys-
tem [2] instantiated with the BN254 curve as used in Ethereum. The theoretical advances
presented in this work are, however, broadly applicable to any pairing-based cryptographic
system. Throughout this paper, we systematically present and benchmark a series of ad-
vances in proof strategies: from recent state-of-the-art pairings, to new field extension arith-
metic, to optimized proof verification techniques including refined Miller loop and final
exponentiation steps.
We introduce each technique, explain its motivation and implementation, and directly
compare its performance against the previous generation, highlighting measurable improve-
ments at every stage. Our roadmap provides insight for future applications and paves the
way towards efficient, sound, and scalable ZK-SNARK verification for real-world use cases.

## Structure

- `code/` - Code for implementations and benchmarks
- `paper/` - LaTeX source files for the research paper
- `results/` - Experimental results and data

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

## Benchmarks

Before implementation:

| circuit                    | MULMOD | ADDMOD | POSEIDON | ~cycles |
| -------------------------- | ------ | ------ | -------- | ------- |
| Miller n=1 BLS12_381       | 4936   | 4966   | 1580     | 90154   |
| Miller n=2 BLS12_381       | 8030   | 8171   | 2276     | 141734  |
| Miller n=3 BLS12_381       | 11356  | 11608  | 3088     | 198070  |
| MultiPairing n=1 BLS12_381 | 10064  | 14027  | 3913     | 217841  |
| MultiPairing n=2 BLS12_381 | 13158  | 17232  | 4609     | 269421  |
| MultiPairing n=3 BLS12_381 | 16484  | 20669  | 5421     | 325757  |
| Miller n=1 BN254           | 5984   | 5927   | 1810     | 101558  |
| Miller n=2 BN254           | 10132  | 10107  | 2740     | 167298  |
| Miller n=3 BN254           | 14456  | 14463  | 3758     | 236382  |
| MultiPairing n=1 BN254     | 10670  | 13150  | 3741     | 203854  |
| MultiPairing n=2 BN254     | 14818  | 17330  | 4671     | 269594  |
| MultiPairing n=3 BN254     | 19142  | 21686  | 5689     | 338678  |

New results:

| circuit                    | MULMOD | ADDMOD | POSEIDON | ~cycles |
| -------------------------- | ------ | ------ | -------- | ------- |
| Miller n=1 BLS12_381       | 2672   | 2686   | 790      | 47588   |
| Miller n=2 BLS12_381       | 4418   | 4525   | 812      | 69558   |
| Miller n=3 BLS12_381       | 6164   | 6364   | 834      | 91528   |
| MultiPairing n=1 BLS12_381 | 7795   | 11742  | 3123     | 175215  |
| MultiPairing n=2 BLS12_381 | 9541   | 13581  | 3145     | 197185  |
| MultiPairing n=3 BLS12_381 | 11287  | 15420  | 3167     | 219155  |
| Miller n=1 BN254           | 3303   | 3228   | 828      | 53130   |
| Miller n=2 BN254           | 5639   | 5576   | 852      | 81898   |
| Miller n=3 BN254           | 7975   | 7924   | 876      | 110666  |
| MultiPairing n=1 BN254     | 7984   | 10446  | 2759     | 155366  |
| MultiPairing n=2 BN254     | 10320  | 12794  | 2783     | 184134  |
| MultiPairing n=3 BN254     | 12656  | 15142  | 2807     | 212902  |