# Rebuttal — #185

Thank you for the careful reading.

**A, compile memory.** 34.5–38.8% across the four configurations, tracking the 33–35%
constraint and ~53% commitment reductions. Each eliminated remainder removes 12
emulated-field elements together with their limb decompositions, hint bindings and
sparse-matrix entries, so the saving is in per-wire bookkeeping, not constraints alone. We
will note this.

**A, irreducibility.** P_12 is irreducible because it is the modulus of the extension field
F_q^12, a property of the pairing instantiation, not of the toolkit. Identity checking needs
only a monic modulus of fixed degree, which gives unique division with deg(R) < deg(P); the
argument in §5.6.1 is purely a degree argument. We will state this in §6 and add a monic
check at ring-group registration.

**B, security level.** Agreed, we will state it explicitly.

**C, attribution.** The Schwartz-Zippel paradigm is Feltroidprime's [Fel23]. [Hou23]
contributes the affine Miller formulae and sparse line representations, and contains no
polynomial identity testing.

**C, new paradigm or local rewriting (§4.1.1 vs §4.1.2).** §4.1 is only a cost summary; the
content is §4.2 and §5. There are two distinct baselines. Table 6 (Cairo) is over [Fel23],
which checks each F_p^12 multiplication as its own identity; we check an entire Miller step,
so a 3-pair zero-bit step drops from four chained relations and 48 commitments to one and
12. Table 7 is over upstream gnark, the fastest production emulated pairing, already
carrying [Hou23] and [NE24]; there the same consolidation gives 33–35% fewer constraints and
~53% fewer commitments. Reaching this requires the operation polynomials of §4.2 for Miller
loop bits, Frobenius correction and final-exponentiation elimination via [NE24], and the
protocol of §5: prover and verifier algorithms, quotient batching, Fiat-Shamir
instantiation, and a soundness proof.

**C, probabilistic or deterministic (§2.2).** Probabilistic. For BN254 with n = 67 and
d_x ≤ 256, soundness error is at most 2^-245 (Theorem 1, §5.6.2) and 2^-244 after
Fiat-Shamir; completeness is perfect. We will say so at §2.2.

**C, side-channel countermeasures.** The pairing inputs are public, so neither party holds a
secret to leak and we claim no resistance. For a native implementation on secret inputs, the
chapter 12 countermeasures are orthogonal: we do not touch the field arithmetic, and our
constraint count and trace depend only on the loop parameter, not the input points. We will
scope this in the introduction.

**Changes.** Monic precondition and check (§5.2, §5.6.1, §6); compile-memory mechanism in
the benchmark discussion; probabilistic statement at §2.2; pointer from §4.1 to §4.2 and §5;
explicit BLS12-381 security level; scoping sentence on physical attacks.

We note this submission is a Major Revision, and that the binding criteria from the previous
round are confirmed as met by Reviewers A and B.
