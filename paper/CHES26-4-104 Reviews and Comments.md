tches2026_4 Paper #104 Reviews and Comments
===========================================================================
Paper #104 Pairing Proofs and Polynomial Ring Toolkit


Review #104A
===========================================================================

Overall merit
-------------
2. Reject and Resubmit (RR)

Relevance to TCHES
------------------
3. Right in the middle of CHES

Novelty/Contribution
--------------------
2. Minor Contribution

Reviewer expertise
------------------
4. I work right in this subfield (I am well-versed with the related
   literature and publish regularly in the subfield)

Paper summary
-------------
The authors revisit the problem of proving pairing computations in SNARK frameworks. They improve on the state of the art by Housni from CT-RSA 2023 by adapting optimizations from the literature that allows to express the whole pairing computation as a single polynomial relation, by building on top of [Fel23] and [NE24]. The technique is benchmarked with a gnark implementation with significant reduction in constraints, commitments and prover resources.

Comments for authors
--------------------
Pros:
+ Proving pairing computations is an interesting research problem, and thinking about how the cost metric allows different trade-offs is always interesting
+ The savings in constraints and prover resources are substantial

Cons:
- Curve BN254 is obsolete, since it barely meets the 100-bit security level
- The relationship to Garaga is unclear, and raises novelty/anonymity concerns
- The contribution is incremental, and essentially combines techniques already present in the literature
- The benchmarking scenario involving 3 pairings is unclear

The paper is reasonably well-written and organized, but it feels quite redundant at times. There are several floats that do not seem to contribute to understanding the technique and end up consuming quite a bit of space. For example, Algorithms 4 and 5 could probably be combined into one algorithm with the Prover and Verifier parts interleaved and color-coded. The interactive and non-interactive versions of the protocol are quite similar (as expected), and providing both does not help the reader much. In any case, the paper fits the page limit comfortably, so these are more nitpicks than roadblockers.

The main shortcomings are with the some technical aspects and novelty of the contribution:
* The entire evaluation is based on computing a 3-pairing, which I understand to be a product of 3 pairings (thus 3 interleaved Miller loops and one final exponentiation), but the paper does not explain exactly how this works. The last page gives a circuit for Groth16 verification, but it is unclear if that matches the intended 3-pairing benchmark introduced earlier. For generality, the paper should give metrics for a single pairing and a product of $n$ pairings.
* The choice of benchmarking curve BN254 is never justified, which is puzzling given the well-known fact that BN curves need a larger field size to meet the 128-bit security level after the Kim-Barbulescu attack. This gets a bit confusing when numbers are given for a BLS12-381 curve in Table 6, but at presumably a different cost metric tailored to the Cairo VM? My reading of Section 7 added to the confusion, since the authors write that the proposed technique is already running in production within another library, so how can there be improvements to claim in Table 6? The statement also creates an obvious issue with the anonymity requirement for submissions.

Even without considering the confusion about what is implemented where, the contribution feels already thin: it heavily builds on previous work and arguably just combines two ideas already present in the literature, obtaining modest improvements by rearranging the polynomials in a better way.

Required Changes
----------------
- Add a modern set of parameters, such as BLS12-381
- Clarify the relationship with the Garaga library, and elaborate on the novelty over the approach implemented there

Questions for authors’ response
-------------------------------
- Why BN254 and not a widely-deployed modern alternative such as BLS12-381?
- What is the exact relationship between the techniques proposed in this work and those implemented in Garaga?
- Do you exploit the sparseness of the Miller line functions for savings in some way?
- What is the application scenario motivating the 3-pairing benchmark? Is that the same circuit described in Section 7.2 for Groth16 verification?



Review #104B
===========================================================================

Overall merit
-------------
4. Minor Revision (MinR)

Relevance to TCHES
------------------
2. Somewhat relevant

Novelty/Contribution
--------------------
3. Major Contribution

Reviewer expertise
------------------
3. I am knowledgeable in this subfield but not expert (I know some related
   works and may have a publication in the subfield)

Paper summary
-------------
This paper proposes a polynomial-identity-based approach for reducing the cost of pairing verification in arithmetized settings. The key idea is to verify larger pairing-related relations, rather than checking extension-field multiplications one by one, and to package the resulting method as a reusable polynomial ring toolkit in gnark. The paper also reports implementation results on BN254/Groth16-style verification workloads, showing noticeable reductions in constraints, commitments, and memory usage.

Comments for authors
--------------------
This paper studies a relevant practical problem and presents an implementation-oriented optimization that seems useful for recursive proving and other constrained verification settings. The paper is generally well motivated, and the reported benchmark improvements are promising.

I think the main strength of the work is the choice of verification granularity: consolidating larger pairing sub-computations into polynomial relations appears to reduce intermediate checking overhead in a natural way. The implementation in gnark and the provided performance data also make the paper more compelling.

My main suggestions are mostly about clarity and presentation. In particular, the paper would benefit from a clearer explanation of the precise contribution relative to prior work, especially in terms of what is conceptually new versus what is an integration of existing ideas. I also think the discussion of the proof protocol / Fiat-Shamir transformation and the implementation-specific commitment mechanism could be presented more clearly. Finally, some notation and terminology could be cleaned up to improve readability.

Overall, I find the paper interesting and practically relevant. With some clarification and polishing, I think it would make a useful contribution.

Required Changes
----------------
1. Clarify the main contribution relative to the most closely related prior works.
2. Improve the presentation of the interactive proof / Fiat-Shamir discussion and its relation to the gnark implementation.
3. Clean up notation and terminology to make the paper easier to follow.
4. Clarify the experimental metrics and benchmarking setup.

Questions for authors’ response
-------------------------------
1. Could the authors more clearly summarize the main conceptual novelty over the closest prior approaches?
2. Could the authors clarify the meaning of the commitment-related metrics used in the benchmarks?
3. Could the authors briefly explain how the protocol description maps to the gnark implementation, especially in the Fiat-Shamir setting?



Review #104C
===========================================================================

Overall merit
-------------
3. Major Revision (MajR)

Relevance to TCHES
------------------
3. Right in the middle of CHES

Novelty/Contribution
--------------------
2. Minor Contribution

Reviewer expertise
------------------
2. I have passing knowledge of the subfield (I know a couple of related
   works)

Paper summary
-------------
The paper improves the algebraic representation of pairing relations which subsequently translates into small and more efficient proofs of bilinear pairings with application to e.g. recursive proofs. More precisely, the authors provide a compact polynomial representation constraining an entire miller loop: In a series of multiplications, they find, it is cheaper to defer reduction by the irreducible polynomial to the very end instead of reducing each intermediate product. Applying this change and incorporating most recent improvements, the authors present significant improvements in constraints and commitments.

Comments for authors
--------------------

### Strengths
The paper clearly describes the technique and targets a very relevant issue. The performance numbers show clearly the significant impact and the implementation details show careful engineering work beyond the given "reduction optimization". 

### Weaknesses
The paper is a bit ambiguous about the relation to the previous results. Given the practical evaluation is such a strong point I believe this could be improved: For Table 6, the description (411-414) makes it not very clear what is compared. What is the exact state of "Before" and "After", when it comes to contributions by this work, by Fel23, Hou23, NE24.
A similar point applies to the "gnark baseline" in Table 7, the paper should be more precise about the concrete state of this baseline with respect to recent academic results. 
Again on Table 6: the meaning of MULMOD, ADDMOD, COMMITS, VM STEPS is never properly defined, making this table even harder to understand. 
The polynomial ring toolkit is an interesting feature but it is not really evaluated beyond the bilinear pairing use case making assessment as a separate contribution hard.

### Editorial Comments
- The points about 29%, 60%, 42% appear almost identically a couple of times.
- 178 -179, the sentence is broken.
- 258-259: It may be worth considering to not allow x to be a scalar and a formal intermediate. 
- The API description of 6.3-6.4 feels very much like a library documentation and could be deferred to the appendix.



Rebuttal Response by Author [SHRAMEE SRIVASTAV <shramee.srivastav@gmail.com>] (490 words)
---------------------------------------------------------------------------
We thank the reviewers for their thoughtful feedback. We are encouraged that they found the problem highly relevant and practical (104C), the improvements significant (104A), and our consolidation of verification into larger polynomial relations to be a Major Contribution (104B).

-------------
@104A – BN254 security and modern parameters (BLS12-381):

BN254 remains widely used in production (e.g., Ethereum, Aztec, SP1, Google's QC paper proof [Goo26]), motivating its use for our `gnark` benchmarks. As shown in Table 6, we have also successfully benchmarked our technique on BLS12-381 within the CairoVM, achieving a 32.7% reduction in VM steps and a 41.6% reduction in commitments. We commit to adding corresponding BLS12-381 `gnark` benchmarks in the next iteration.

-------------
@104A – Anonymity and Garaga:

Our double-blind anonymity remains fully intact. Garaga implemented an earlier proof-of-concept of these techniques in 2024, which was adapted from a technical write-up and implementation we circulated. The purpose of this submission is to formally introduce, mathematically ground, and systematically evaluate these techniques, expanding significantly beyond that initial Cairo-specific implementation.

-------------
@104A, @104C – Baselines and the 3-Pairing Benchmark:

* Table 6 Baselines & metrics: "Before" represents Garaga using [Fel23]; "After" uses our technique. As noted in the footer, MULMOD/ADDMOD denote modular field operations, COMMITS are Poseidon input commitments, and VM STEPS are CairoVM CPU cycles. These provide a low-level estimate of computational complexity inside the virtual machine.
* Table 7 gnark Baseline: This represents the current state-of-the-art for the fastest possible emulated pairings, incorporating techniques from [NE24] and [Hou23].
* 3-Pairing Benchmark: This does correspond to Groth16 verification. gnark API only exposes single-pairing verification with previous Fq12 accumulation via `AssertMillerLoopAndFinalExpIsOne(P G1, Q G2, prev GTEl)`, so we implemented the circuit as one multi-pairing and one single pairing with accumulation to enable a direct comparison. Our production implementation will use more efficient `AssertMultiMillerLoopAndFinalExpIsOne`, but we prioritized parity with the `gnark` baseline for this benchmark.
* Sparseness: Yes, all compared techniques (including ours) exploit the sparseness of the Miller line functions. `Sparse01379` savings are included in Section 4 constraint counts.

-------------
@104A, @104B – Novelty and Protocol mapping:

* Conceptual Novelty: Unlike [Fel23], we consolidate entire Miller loop iterations into single polynomial relationships rather than verifying individual field multiplications. We also packaged this into a generalized `gnark` toolkit.
* Commitment Mechanisms: In CairoVM (Table 6), the commitments use Poseidon hashing. In `gnark`, commitments follow [BSB23] and [BSB], using a combination of scalar multiplications and MiMC to avoid in-circuit hashing. Regardless of the environment, our work greatly reduces the absolute number of commitments necessary for sound PIT.
* Mapping to Implementation: Algorithm 4 is implemented in the Miller Loop in [Work:bn254]. Output arrays Q and R are stored in `f.deferredPolyChecks`. At the end of the circuit, `performDeferredRingChecks` in [Work:fpoly] performs the checks described in Section 5.6 using `gnark`'s `api.Commit` for Fiat-Shamir heuristic commitments.

-------------
@104C – Editorial comments:

Thank you for taking the time to share the editorial comments, we will fix all issues in the next iteration.

-------------
References

[Goo26]: https://quantumai.google/static/site-assets/downloads/cryptocurrency-whitepaper.pdf
[Work:fpoly]: https://anonymous.4open.science/r/tiny-gnark-DDCD/std/math/emulated/field_polyring.go
[Work:bn254]: https://anonymous.4open.science/r/tiny-gnark-DDCD/std/algebra/emulated/sw_bn254/pairing.go#L593-L640



Comment @A1 by Reviewer A
---------------------------------------------------------------------------
Dear authors,

After the post-rebuttal discussion, we concluded that the paper would be receive a major revision outcome.

We've collected the following binding revision criteria from the reviews and rebuttal:
- Clarify the main contribution relative to the most closely related prior works (such as Fel23 and Garaga)
- Add additional benchmarks for BLS12-381 for a more comprehensive evaluation, as per the rebuttal
- Add clarification about the benchmarking setup, as per the rebuttal, in particular for the 3-pairing setting
- Improve the mapping between algorithms and implementation, in particular the application of Fiat-Shamir heuristic
- Qualify the commitment-based performance metric by specifying the various schemes, their overhead and impact on overall performance
- Minor presentation issues pointed out by the reviewers