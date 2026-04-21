# IACR TCHES Peer Review

**Manuscript:** *Pairing Proofs and Polynomial Ring Toolkit*
**Venue:** IACR Transactions on Cryptographic Hardware and Embedded Systems
**Reviewer Confidence:** High (familiar with ZKP constraint systems, pairing-based cryptography, and gnark)

---

## 1. Summary

This paper addresses the high cost of verifying elliptic curve pairings inside arithmetized constraint systems (R1CS, SCS/Plonkish) and resource-constrained virtual machines (Ethereum VM, BitVM, ZKVMs). The authors propose a public-coin interactive proof system that consolidates the Miller loop iterations and final exponentiation of an optimal ate pairing over BN254 into a single multi-operand polynomial identity, verified probabilistically via the Schwartz-Zippel lemma. Rather than checking each field multiplication individually (as in [Fel23]), the scheme batches entire Miller loop steps—including line function evaluations, residue witness multiplications, and Frobenius corrections—into one polynomial equation per step, then further batches all equations via a random linear combination into a single accumulated quotient checked at a random evaluation point.

The paper makes two concrete contributions: (1) the proof system itself, with a formal two-round interactive protocol (and its Fiat-Shamir non-interactive variant), yielding a 39% reduction in SCS constraints and a 60% reduction in commitments for a Groth16 verification simulation on BN254; and (2) a reusable **polynomial ring toolkit** implemented in gnark's `emulated` package, providing a clean API (`NewPolyRingCheck`, `MulPolyRings`, `PolyRingAccumulator`, `performDeferredRingChecks`) for deferred polynomial-ring multiplication and verification in arbitrary emulated fields.

---

## 2. Novelty & Impact

**Originality: Moderate-to-High.** The paper synthesizes three prior ideas—El Housni's R1CS-optimized Miller loop formulas [Hou23], Feltroidprime's Schwartz-Zippel-based polynomial identity testing for $\mathbb{F}_{q^{12}}$ [Fel23], and Novakovic-Eagen's residue witness technique for final exponentiation [NE24]—into a unified framework. The genuine novelty lies in the *granularity shift*: verifying entire Miller loop steps as single polynomial identities rather than individual binary multiplications, thereby eliminating intermediate remainder commitments. This is a clean and effective insight.

**Significance: High for the target domain.** Pairing verification is a critical bottleneck in recursive proof composition and on-chain verification. A 60% reduction in commitments directly translates to smaller proof transcripts and lower gas costs in blockchain deployments. The reusable toolkit also lowers the barrier for other developers to adopt the technique.

**Limitations on novelty:**
- The core mathematical machinery (Schwartz-Zippel PIT, Euclidean division in $\mathbb{F}_q[x]/P_{12}(x)$) is standard. The contribution is primarily an *engineering optimization* within an existing algebraic framework, not a new cryptographic primitive or fundamentally new proof technique.
- The scope is narrow: only BN254 with optimal ate pairings. Generalization to BLS12-381, BW6-761, or other pairing-friendly curves is mentioned nowhere, and degree/sparsity arguments would need revisiting for different tower constructions.

---

## 3. Technical Soundness

### 3.1 Correctness of the Polynomial Formulation

The polynomial equations in Section 4.2 (Equations 1–4) are correctly derived from Algorithm 2. The degree analysis in Tables 2–5 is consistent: for a 3-pair zero-bit step, the LHS product $F^2 \cdot L_1 \cdot L_2 \cdot L_3$ has degree $2 \times 11 + 3 \times 9 = 49$, and the RHS $R + Q \cdot P_{12}$ with $\deg(Q) = 49 - 12 + 1 = 38$ checks out. The coefficient counts are correct given $\mathbb{F}_{q^{12}}$ representation (12 coefficients per element) and Sparse$_{01379}$ representation (4 non-zero coefficients per line function).

### 3.2 Security Analysis

The soundness proof (Theorem 1, Section 5.7.2) is essentially correct but has notable issues:

- **The soundness bound application is slightly imprecise.** The bivariate Schwartz-Zippel bound used is $\Pr[P_{\text{err}}(z,x) = 0] \leq \frac{\deg_z + \deg_x}{|\mathbb{F}_q|}$. However, the standard multivariate Schwartz-Zippel lemma bounds by $\frac{d}{|\mathbb{F}|}$ where $d$ is the *total degree*, not the sum of individual degrees. For a bivariate polynomial, these coincide only when the polynomial is examined appropriately. The authors should clarify whether they are applying the lemma to $P_{\text{final}}$ viewed as a bivariate polynomial over $\mathbb{F}_q^2$ (where $(z, x)$ are drawn independently) or sequentially. As stated, the bound $\frac{n - 1 + d_x}{|\mathbb{F}_q|}$ is correct for independent uniform sampling over $\mathbb{F}_q^2$, but the text could be more precise about which form of the lemma is being invoked.

- **Remark 1 is critical but underemphasized.** The entire soundness argument assumes the verifier evaluates polynomials without arithmetic error. The deferred-reduction inner product optimization (Section 6.6) introduces a risk: if intermediate limb values overflow the native characteristic $p$, the algebraic identity breaks silently. Lemma 1 provides a concrete overflow bound ($2^{146} \ll 2^{254}$), which is reassuring for BN254 with 4-limb/64-bit emulation. However, this bound is *parameter-specific* and would need re-derivation for any other curve or limb width. The paper should state this caveat explicitly and ideally provide a general formula.

- **The transition from IOP to non-interactive proof via Fiat-Shamir is handled too briefly.** Section 5.5 cites [BSCS16] for the knowledge soundness bound but does not discuss potential pitfalls of applying Fiat-Shamir in the algebraic group model or when the hash function is instantiated concretely (e.g., MiMC). gnark's specific commitment mechanism via MiMC hashing is mentioned in Section 6.5 but its security implications are not analyzed.

### 3.3 Threat Model

The paper implicitly assumes a standard honest-verifier model for the interactive protocol, which is appropriate. However, no explicit threat model is stated. For a TCHES submission where hardware/embedded deployment is a concern, the authors should discuss whether the verifier's random sampling is vulnerable to side-channel leakage in constrained environments, and whether the Fiat-Shamir instantiation is constant-time.

---

## 4. Evaluation & Reproducibility

### 4.1 Benchmarks

The benchmark in Table 6 is well-structured: it compares both SCS (Plonkish) and R1CS backends on a realistic Groth16 verification simulation circuit over BN254. The reported numbers (39% SCS constraint reduction, 60% commitment reduction, ~42% compile-time memory reduction, ~52% solver memory reduction) are substantial and credible given the theoretical analysis.

**Strengths:**
- The benchmark circuit (`Groth16Simulation`) is well-chosen: it directly models the pairing operations in real recursive proof verification.
- Both constraint system backends (SCS, R1CS) are tested.
- Multiple metrics are reported (constraints, commitments, compile time, compile memory, solve time, solve memory).
- Code is publicly available (GitHub links and Gist provided).

**Weaknesses:**
- **Single curve, single circuit.** All benchmarks are on BN254 with one specific circuit. No other pairing-friendly curves or circuit configurations are tested. This significantly limits the generalizability claims.
- **No prover/verifier wall-clock time breakdown.** The "solve time" metric conflates prover computation, but the paper does not separate the cost of hint computation (polynomial multiplication and division outside the circuit) from constraint solving. For practical deployment, this distinction matters.
- **No comparison with Garaga [FE].** The paper acknowledges Garaga as "an earlier proof of concept" in Cairo/Starknet but does not benchmark against it, even qualitatively. Given that Garaga targets a similar problem, this omission weakens the comparative analysis.
- **Hardware specification is minimal.** "Apple M3 Pro (goarch:arm64)" is stated, but no information on RAM, OS version, Go version, or gnark version is provided. Reproducibility would benefit from a more complete environment specification.
- **No statistical rigor.** Benchmark numbers appear to be single-run measurements with no variance, confidence intervals, or repeated trials reported.

### 4.2 Reproducibility

The provision of both the circuit definition (GitHub Gist) and the implementation (GitHub repository) is commendable and significantly aids reproducibility. However:
- The Gist link and repo link should be anonymized for double-blind review (they appear to contain the author's username "shramee" and organization "mistcash").
- No build/run instructions or specific dependency versions are documented in the paper.

---

## 5. Clarity

**Overall: Good, with room for improvement.**

The paper is generally well-organized with a logical flow: motivation → background → related work → polynomial formulation → full proof protocol → toolkit → benchmarks → conclusion. The algorithms (1–6) are clearly presented with line-by-line comments. The protocol diagrams (Figures in Sections 5.4 and 5.6) are helpful.

**Issues:**
- The notation is introduced somewhat abruptly. $\text{Sparse}_{01379}$ is used before being formally defined (first appears in Section 4.1, defined parenthetically).
- The paper switches between "constraint" (C) and "commitment" metrics without always making clear which is being discussed, particularly in Section 4.1 where the comparative analysis jumps between the two.
- Section 6 (Polynomial Ring Toolkit) reads more like API documentation than a research paper section. While the detail is valuable for practitioners, it disrupts the theoretical flow. Consider moving the detailed API to an appendix.
- The paper is 22 pages, which is on the longer side for TCHES. Some compression is possible in Sections 6.3–6.5 without loss of substance.

---

## 6. Constructive Feedback

### Major Issues

- **M1: Anonymization failure.** The GitHub links in Section 7.2 (footnotes 2 and 3) contain identifiable usernames ("shramee", "mistcash"), violating double-blind review requirements. These must be anonymized or replaced with anonymous repositories.

- **M2: Lack of curve generality.** The entire paper is specialized to BN254. The authors should either (a) provide a general framework showing how the polynomial equations adapt to BLS12-381 or other curves, or (b) explicitly scope the contribution to BN254 and discuss what changes for other curves (different tower constructions, different sparsity patterns, different loop lengths).

- **M3: Incomplete security discussion.** The Fiat-Shamir instantiation via MiMC and gnark's commitment mechanism needs more rigorous treatment. What are the concrete security assumptions? Is MiMC's algebraic structure a concern when hashing algebraic data? The [BSB, BSB23] references describe the mechanism, but the security implications for this specific protocol should be analyzed.

- **M4: Missing comparison with Garaga.** Given [FE] targets the same problem space, a direct comparison (even at the constraint-count level if runtime comparison is infeasible) would strengthen the paper.

- **M5: Overflow safety is parameter-specific.** Lemma 1's bound is proved only for BN254 with 4-limb/64-bit arithmetic. A general formula or at least a discussion of when this bound might be violated (larger extension degrees, different limb widths) is needed.

### Minor Issues

- **m1:** The abstract claims "~42% reduction in compile-time memory" but Table 6 shows 41.8% (SCS) and 43.1% (R1CS). State the range or be precise.
- **m2:** Section 2.2 title says "Schwartz-Zippel lemma" but the text correctly notes it should be "Demillo-Lipton-Schwartz-Zippel." Be consistent.
- **m3:** In Algorithm 2, the comment "Skip $T \leftarrow T - Q_2(Q)$" on line 22 is confusing—clarify whether this is an intentional optimization or a notational shorthand.
- **m4:** Table 1 footnotes use "C*" to denote costs excluding Fiat-Shamir overhead, but this notation is not defined before the table.
- **m5:** The `emulated` package reference [Conb] points to a specific tree path in the gnark repo. This URL may break with future refactors; consider citing a tagged release.
- **m6:** Section 4.3 states $\deg(P) < 2^8$ with a footnote giving the approximate formula $\deg(P) \approx 88 + 18(k-3)$ for $k$-pair pairings. For $k = 3$, this gives $\deg(P) \approx 88$, which is less than $2^8 = 256$. However, the footnote then claims "this bound can be extended to $2^{10}$"—clarify what motivates this extension and under what conditions.
- **m7:** No acknowledgments section is present (expected for camera-ready but acceptable for submission).
- **m8:** References [BSB] is a HackMD link—this is not a stable archival reference. If there is a preprint or published version, cite that instead.
- **m9:** The paper would benefit from a table summarizing all polynomial equations (Eqs. 1–4) with their per-step and total costs for the full pairing, rather than having this information spread across Tables 2–5.

---

## 7. Recommendation

**Decision: Minor Revision**

**Justification:** The paper presents a clean and effective optimization for pairing verification in constraint systems, achieving meaningful improvements (39% constraint reduction, 60% commitment reduction) confirmed by a working gnark implementation. The core technical contribution—consolidating Miller loop steps into single polynomial identities with batched quotient accumulation—is sound and well-motivated. The reusable polynomial ring toolkit is a valuable engineering contribution.

However, the paper requires revision on several fronts before acceptance: (a) anonymization must be fixed for double-blind compliance, (b) the security analysis needs strengthening around the Fiat-Shamir instantiation and overflow safety generalization, (c) the evaluation should include at least a discussion of other curves and a comparison with Garaga, and (d) minor presentation issues should be addressed. None of these issues are fundamental—they can be resolved in a single revision cycle.

The work is relevant to TCHES's scope (efficient cryptographic implementations in constrained environments) and would be a solid contribution once these issues are addressed.

---

*Review prepared in the capacity of an expert peer reviewer for IACR TCHES. All assessments reflect the reviewer's independent technical judgment.*
