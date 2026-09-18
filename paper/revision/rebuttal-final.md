# TCHES #185 rebuttal, final (sent)

We thank the reviewers for their careful reading and feedback. We are glad that you found the work practical and well-systemised (A), the reductions impactful and significant (A, B, C) and prior reviewers' concerns well addressed (A, B).

Reviewer A
-----------
Compile-memory reductions, To clarify the 34-38% memory reductions versus the 33-35% constraint reductions, we will add to §7.2 that the ~53% reduction in commitments reduces per-wire bookkeeping as described in §5.5.2.

Modulus irreducibility: Identity checking requires only a monic modulus to have unique quotient and remainder from Euclidean division. We will correct line 311 to state the monic requirement on P_12 and document the precondition in the toolkit.

Reviewer B
-----------
Security level, We will state it explicitly as 128-bit.

Reviewer C
-----------
Clarification on attribution. The polynomial identity testing paradigm is [Fel23]'s. [Hou23] contributes the affine Miller formulae and sparse line representations and efficient techniques for in-circuit operations, it contains no polynomial identity testing.

New paradigm or local rewriting.

[Fel23] introduced the polynomial identity testing for Fq^12 paradigm, we clarify this in §3. We contribute: (1) optimal algebraic statements covering the Miller loop and the elimination of final exponentiation from [NE24] (§4), (2) an end-to-end specification of the proof system with Fiat-Shamir oracle instantiations with full soundness proofs (§5) and, (3) a production-ready polynomial ring toolkit accepting any monic modulus over emulated fields (§6).

The contribution is what makes the paradigm deployable. Under [Fel23] a 3-pairing check commits to 5,689 (BN254) and 5,421 (BLS12-381) elements, above Starknet's 4,000-element call-data limit; with our work it commits to 2,807 and 3,167 (Table 6), which fits Groth16 verification on Starknet (§7).

Clarify if SZ (§2.2) and the algorithm are probabilistic or deterministic: Probabilistic with soundness error δ < 2^{−244} after Fiat-Shamir. Applicability of the SZ lemma and soundness bounds for the algorithm are given in §4.3 and in the proof of Theorem 1 (§5.6.2).

Physical attack countermeasures such as side-channel analyses

In our setting the paired points are public inputs, the verification key and proof elements, so this is not within the protocol's current scope. If it helps reviewer's concern, the reformulation is confined to Fq^12 field elements and yields a fixed number of constraints regardless of inputs.

---

## Camera-ready TODOs from this pass

- 6-polyring-toolkit.tex:33 uses `\mathbb{F}_{p^{12}}`, the only p in the paper. Change to `\Fe{12}`.
- §7.2: add the per-wire / hint-call bookkeeping explanation for compile memory (promised to A).
- Line 311: monic requirement on P_12 instead of irreducibility, plus toolkit precondition (promised to A).
- State 128-bit security level explicitly for the BLS12-381 sentence in §2 background (promised to B).
- §2.2: state that SZ gives probabilistic verification, with the bound (promised to C).
