# Rebuttal, TCHES 2027.1 #185
## Pairing Proofs and Polynomial Ring Toolkit

> Everything above the divider is written to be pasted into the response field. Notes below the divider are for us.

---

We thank the reviewers. Reviewers A and B confirm that the Major Revision addressed their round one concerns, and Reviewer B states that the submission satisfies the binding list of revision criteria. Reviewer C raises two questions and one concern about the depth of the contribution, which we answer first.

### C, Q1. Does Schwartz-Zippel give probabilistic or deterministic verification?

Probabilistic, with one sided error. Section 2.2 states the lemma but not its consequence for our protocol, which we will fix.

Completeness is perfect. An honest prover produces the zero polynomial identically, so the verifier always accepts (Section 5.5.3). Soundness is statistical: a cheating prover is accepted with probability at most $(n-1+d_x)/|\Fq|$ (Theorem 1). For BN254 with $n=67$ equations and $d_x \le 256$ this is below $2^{-245}$, and below $2^{-244}$ for the Fiat-Shamir version in the random oracle model.

On applicability, the lemma needs three preconditions, all of which we establish rather than assume:

1. The error polynomial is non zero as a formal polynomial. This is proved in Section 5.5.1 by a degree argument that does not depend on the challenges, a corrupted remainder has degree at most 11 while $P_{12}$ has degree 12, so no valid quotient exists.
2. The challenges are sampled after the values they must not depend on. $\mathcal{R}$ is committed before $z$, and $Q_\text{acc}$ before $x$ (Section 5.3.1).
3. Degrees are bounded by construction, $d_x < 2^8$ for all practical pair counts.

The resulting $2^{-245}$ sits far below the roughly 100 bit security of BN254 itself, so the identity test is not the weakest link. We will add a paragraph to Section 2.2 stating this, with a forward reference to Theorem 1.

### C, Q2. Compatibility with side channel countermeasures

We should have scoped this explicitly and will do so. Our contribution accelerates the *verification* of a pairing whose inputs are public. For Groth16 verification the points are proof elements and public input commitments, transmitted in the clear, and the prover supplied hints, the residue witness $c$ and the remainders $\mathcal{R}$, are absorbed into the Fiat-Shamir transcript and are therefore public by construction. There is no secret operand in the accelerated computation, so classical side channel analysis does not apply to it, and we make no side channel claims.

Where a secret pairing input does exist in some other deployment, our transformation is neutral with respect to Chapter 12 of the *Guide to Pairing-Based Cryptography*. We change how accumulator updates are *checked*, not how $T \leftarrow [2]T$, $T \leftarrow T \pm Q$ or the line functions are *computed*, so randomised projective coordinates, point blinding, and multiplicative blinding of the Miller variable act on operands our relations accept unchanged. Control flow is also operand independent, the loop structure follows the NAF of the curve constant $s = 6t+2$ and the number of relations, hint calls, commitments and evaluations is fixed at compile time.

One point we flag as genuinely open rather than as an advantage: following [OPP24] we replace the final exponentiation with a residue witness check, and the final exponentiation is known to frustrate certain fault attacks on the Miller loop. Whether the residue witness formulation preserves that property in a secret bearing setting is a question we have not evaluated and do not claim either way. We will state this as a limitation.

### C. Is this a new paradigm or a local rewriting?

We do not claim a new paradigm and we will make the claim boundary explicit. The polynomial identity paradigm is due to Feltroidprime [FXFM23], and the constraint system cost model it exploits is due to El Housni [PR1CS22]. What we claim is a change of granularity, its formalisation, and its systematic evaluation.

The reviewer is right that Sections 4.1.1 and 4.1.2 read as a local rewriting, because that contrast is deliberately minimal, it isolates one Miller step to make the delta legible. That was a presentation error on our part, since it invites the reading that the step is the whole contribution. Two things follow it that are not local:

**The commitment count changes order in the number of pairs.** For a $k$-pair Miller step, [FXFM23] chains $k+1$ multiplications and must commit a remainder for each, so commitments per step are $\Theta(k)$. Our consolidated relation commits one remainder regardless of $k$, so they are $\Theta(1)$, additional pairs add line factors to an existing relation rather than new relations. At $k=3$ this is 48 versus 12 per zero bit step, and it widens with $k$. This is not cosmetic in the target setting: commitments, not constraints, are the metered cost on chain, and the reduction is what first fit Groth16 verification inside Starknet's transaction limits (Table 6, 50.7% fewer commitments).

**The protocol and its analysis are new.** Batching all $n$ quotients across the entire pairing into a single $Q_\text{acc}$ (Section 5.1) reduces $n$ identity checks to one degree 76 evaluation, where [FXFM23] batches only within a multiplication. Theorem 1 gives the first soundness bound for the resulting bivariate identity, Section 5.3 instantiates Fiat-Shamir, and Lemma 2 gives an overflow safety condition for the deferred reduction the implementation relies on. The Garaga deployment [grg] shipped a Cairo specific version of the technique in 2024 with no formal treatment, which is the gap this paper fills.

We will add a short "what is new" paragraph at the end of Section 4.1 stating precisely which parts are inherited and which are ours.

### C. Functional equivalence

To confirm the reviewer's reading, the computed relation is unchanged, only the granularity of its verification differs. The benchmark compiles the same gnark circuit, `PairingCheck` over three pairs, against the baseline and against our fork, with no changes to the circuit definition or test harness, on both BN254 and BLS12-381 (Section 7.3).

### A1. Why 35% fewer constraints gives 38% lower compile memory

Compile memory tracks the internal variable table and range check tables as well as the constraint blob, and those shrink faster than the gate count. Each eliminated intermediate remainder removes a full emulated $\Fe{12}$ element, twelve limb decomposed wires and their range check entries, while removing only one multiplication gate. A 3-pair zero bit step goes from four committed remainders to one, so remainder wires fall by three quarters while constraints fall by 35%, and compile memory lands between the two at 38.3%. The same ordering holds on BLS12-381. We will add this explanation to the Table 7 discussion and add an internal variable count row so the effect is visible directly.

### A2. "Any modulus" versus irreducibility

The current wording conflates two separate requirements, and we will separate them.

Soundness of the identity protocol requires only that $P$ be monic, or have an invertible leading coefficient, with $\deg P \ge 1$. Euclidean division is then unique and the forgery argument rests on the degree gap $\deg(\Delta_i) < \deg(P)$, not on irreducibility. The toolkit is modulus agnostic in this sense, which is what makes the lattice ring direction of the Further Work section possible, since $x^n+1$ is reducible over $\Z_q$.

Irreducibility is required for the *pairing instantiation*, because $\Fq[x]/(P)$ must be a field for the residue witness inversion and $\Fe{12}$ arithmetic to be well defined.

We will add a precondition remark after Theorem 1, mirror it in the `NewPolyRingCheck` documentation, and add an optional irreducibility check at construction. The modulus is a compile time constant, so a Rabin test costs nothing in circuit.

### B1. BLS12-381 phrasing

Agreed, the sentence is misleading. We will replace it with explicit security levels rather than a comparative:

> "We also benchmark BLS12-381 (Table 7). Under exTNFS based estimates, BN254 offers roughly 100 bits of security in $\G_T$ and BLS12-381 roughly 126 bits, so the two curves bracket the 128 bit target rather than exceeding it."

References for both estimates will be added.

### Summary of changes

| # | Change | Location |
| --- | --- | --- |
| 1 | State one sided error, concrete soundness bound, and the three applicability preconditions | End of Section 2.2 |
| 2 | Add scope subsection on physical attacks, including the final exponentiation limitation | Section 5.5 |
| 3 | Add "what is new" paragraph separating inherited from contributed | End of Section 4.1 |
| 4 | Explain compile memory via wire and range check elimination, add internal variable row | Section 7.3, Table 7 |
| 5 | Precondition remark separating monic from irreducible, documentation, optional check | After Theorem 1, Section 6.3 |
| 6 | Replace BLS12-381 phrasing with explicit security levels plus references | Section 2.1 |

---
---

## Internal notes, do not submit

**What changed from the first pass, and why**

- **Do not lead with correcting C's attribution.** C wrote that the results are "fully based on the paradigm change introduced by El Housni, leveraging Schwartz-Zippel". That parses two ways, either C thinks El Housni introduced polynomial identity testing (wrong) or C means El Housni's cost model plus SZ (fair). Leading with a correction to the one hostile reviewer is a bad trade, it hardens them and reads as deflection. The lineage is now stated constructively inside the substantive answer instead.
- **Dropped the claim that removing the final exponentiation narrows the attack surface.** This is probably backwards. The final exponentiation is known to frustrate fault attacks on the Miller loop, so removing it plausibly *widens* the surface in a secret bearing setting. An expert reviewer would have shredded it. It is now an explicit limitation, which is also the more honest position since we have not evaluated it.
- **Fixed the asymptotic claim.** Commitments are not asymptotically better in the number of Miller steps, that count is fixed at about 64 for BN254, it is a constant factor. The real asymptotic statement is in the number of pairs $k$, $\Theta(k)$ to $\Theta(1)$ commitments per step. That version is both true and sharper.
- **Stopped arguing that this is a paradigm shift.** Reviewer B has expertise 4 and still calls it incremental. Fighting that loses. Conceding the framing and claiming systematisation plus formalisation plus artifact is both true and squarely what TCHES publishes.
- **Answered C's implicit functional equivalence question.** C wrote "as far as the reviewer understands, this rewriting happens without changing the functionality". That is a request for confirmation and the first pass did not answer it at all.
- **Conceded the 4.1.1 versus 4.1.2 point.** C is right that it reads shallow. That is a real presentation defect, not just a misreading, and conceding it costs nothing while buying credibility for the rest.

**Open decision for you**

Reviewer C's review does not reference round one anywhere, while A and B both do. If C is a newly added reviewer, the binding revision criteria argument matters and belongs in a note to the editor rather than in the response to reviewers. I left it out of the draft because it is a process argument that reads badly if C was in fact present in round one. Worth checking the round one review set before deciding. If C is new, add roughly this to the editor note:

> Reviewers A and B, both of whom reviewed the original submission, confirm that the revision addresses their concerns and satisfies the binding list of revision criteria. We have answered Reviewer C's questions in full and will make the corresponding edits.

**To do before submitting**

| Item | Why |
| --- | --- |
| Re-run benchmarks for the internal variable counts | Change 4 promises a new Table 7 row |
| Confirm the Rabin check is straightforward in `NewPolyRingCheck` | Change 5 promises an optional check, do not promise if awkward |
| Pin references for the 100 and 126 bit figures | Barbulescu-Duquesne and Guillevic, reviewer B said BLS12-381 "barely meets" 128 so roughly 126 is consistent with them |
| Check the response field length limit | Draft is around 1,100 words, trim A1 and B1 first if capped |
