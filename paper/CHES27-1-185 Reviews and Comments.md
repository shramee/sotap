tches2027_1 Paper #185 Reviews and Comments
===========================================================================
Paper #185 Pairing Proofs and Polynomial Ring Toolkit


Review #185A
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
2. I have passing knowledge of the subfield (I know a couple of related
   works)

Paper summary
-------------
This paper presents a public‑coin interactive proof system for verifying elliptic curve pairings in arithmetized circuits. Its core idea is to consolidate the Miller loop iterations—traditionally verified as a sequence of individual extension‑field multiplications (as in Fel23)—into a single polynomial identity per Miller step, leveraging the Schwartz‑Zippel lemma for batch verification. The reduction in intermediate remainder commitments (from 48 to 12 for a 3‑pair zero‑bit step) lowers both constraint counts and on‑chain proof sizes. The authors provide a complete, reusable implementation in gnark as a polynomial‑ring toolkit, and benchmark their scheme on BN254 and BLS12‑381, for both SCS (Plonkish) and R1CS backends. For a 3‑pairing product check (the workload of Groth16 verification), they report ≈33–35% fewer constraints, ≈53% fewer committed elements, and substantial reductions in compile‑time and solver memory. The revision thoroughly addresses prior reviewer concerns: it quantifies the gap over Fel23, adds BLS12‑381 results, replaces the benchmark circuit with a uniform 3‑pairing check, clarifies the Fiat‑Shamir transformation, and provides a clear mapping from algorithms to code. The paper also includes a full soundness proof (Theorem 1) and a formal overflow‑safety lemma for deferred reduction. Overall, the work offers a practical, well‑systematised improvement with immediate impact on recursive proof systems and resource‑constrained verification environments.

Comments for authors
--------------------
I have only a few minor suggestions:

Compile‑memory reduction: Briefly explain why a 35% constraint reduction yields ~38% lower compile memory (e.g., fewer intermediate variables).

Modulus irreducibility: The toolkit claims to support “any modulus”, but soundness relies on irreducibility. Please clarify this precondition in the documentation or add a check.



Review #185B
===========================================================================

Overall merit
-------------
4. Minor Revision (MinR)

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
The paper essentially addresses all my reviewer comments:
- It clarifies the relationship with previous work in a more precise manner
- It adds timings for BLS12-381
- It motivates and clarifies where the 3-pairing benchmark comes from

I also see improvements in Section 5 to better relate the protocols and algorithms, and clarifications to the commitment metric. One last presentation issue is rephrasing the sentence "We also benchmark BLS12–381 (Table 7) for higher security requirements" which could read misleading as if BLS12-381 is a high-security choice of parameters when it barely meets the 128-bit requirement. Maybe just stating the security level explicitly is a better option already.

I still think the contribution is somewhat incremental, but the submitted Major Revision appears to satisfy the binding list of revision criteria.



Review #185C
===========================================================================

Overall merit
-------------
1. Reject

Relevance to TCHES
------------------
3. Right in the middle of CHES

Novelty/Contribution
--------------------
2. Minor Contribution

Reviewer expertise
------------------
3. I am knowledgeable in this subfield but not expert (I know some related
   works and may have a publication in the subfield)

Paper summary
-------------
This articles addresses an optimization of elliptic curve pairings, with applications to embedded systems.
The core of the idea is to shift the subject of the verification to a higher level, thereby enabling computational saving.
As far as the reviewer understands, this rewriting happens without changing the fonctionality.

The results are quantitively significant.
They are fully based on the paradigm change introduced by El Housni [Hou23], leveraging Schwartz-Zippel lemma.

However, the technical improvement is rather shallow when contrasting sections 4.1.1 and 4.1.2.

Comments for authors
--------------------
The authors should make it more clear whether there is in this paper a new paradigm or just a local rewriting.

Questions for authors’ response
-------------------------------
Questions:

It is not clear in §2.2 if the Schwartz-Zippel lemma yields a probabilistic or a deterministic verification. Please be precise about the applicability of this lemma in the context of this algorithm.

Obviously, embedded systems are subject to physical attacks, such as side-channel analyses.
One natural question is whether the proposed optimizations are compatible with known countermeasures (see for instance chapter 12 of "Guide to Pairing-Based Cryptography").