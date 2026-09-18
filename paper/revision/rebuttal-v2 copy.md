# Rebuttal — #185

We thank the reviewers for the careful reading.

## A — compile memory

Compile memory drops 34.5–38.8% across the four configurations, tracking the 33–35%
constraint and ~53% commitment reductions. Each eliminated intermediate remainder removes
12 emulated-field elements along with their limb decompositions, hint bindings, and
sparse-matrix entries, so the saving is in per-wire bookkeeping, not just constraints. We
will say this in the benchmark discussion.

## A — modulus irreducibility

P_12 is irreducible because it is the modulus of the field extension F_q^12; that is a
property of the pairing instantiation. The toolkit is not bound to it. Identity checking
requires only a monic modulus of fixed degree, which gives unique Euclidean division with
deg(R) < deg(P), and the argument in §5.6.1 is purely a degree argument. We will state this
precondition in §6 and add a monic check at ring-group registration.

## B — BLS12-381 security level

Agreed, we will state the security level explicitly.

## C — new paradigm or local rewriting? (§4.1.1 vs §4.1.2)
A clarification on attribution first: the Schwartz-Zippel paradigm is Feltroidprime's
[Fel23], not El Housni's [Hou23]. [Hou23] contributes the affine Miller formulae and sparse
line representations, and contains no polynomial identity testing. Our delta is over
[Fel23].

§4.1 is a cost-model summary; the technical content is §4.2 and §5.

[Fel23] checks each F_p^12 multiplication as its own identity, so every operation commits an
intermediate remainder. We check an entire Miller loop step as one identity, which removes
those commitments: a 3-pair zero-bit step needs four chained relations and 48 commitments
under [Fel23], versus one relation and 12 in our scheme (§4.1). Making that work requires
the operation polynomials of §4.2 for Miller loop bits, Frobenius correction, and
final-exponentiation elimination via [NE24], and the protocol of §5: prover and verifier
algorithms, quotient batching, Fiat-Shamir instantiation, and a soundness proof. The
baseline we benchmark against is gnark's emulated pairing stack, which already incorporates
[Hou23] and [NE24].

## C — is the check probabilistic? (§2.2)

Probabilistic. For BN254 with n = 67 equations and d_x <= 256, Theorem 1 (§5.6.2) bounds
soundness error by 2^-245, and 2^-244 after Fiat-Shamir. Completeness is perfect. We will
make this explicit at §2.2.

## C — side-channel countermeasures

The pairing inputs are public here, so neither party holds a secret to leak, and we make no
resistance claim. Where a native implementation does handle secret inputs, the
countermeasures of chapter 12 of the Guide are orthogonal: we do not touch the field
arithmetic, and our constraint count and trace depend only on the loop parameter, not the
input points. We will scope this in the introduction.

## Changes

- §5.2, §5.6.1, §6: separate the monic requirement of the toolkit from irreducibility of
  P_12; add a monic check.
- Benchmark discussion: compile-memory mechanism, 34.5–38.8%.
- §2.2: state the check is probabilistic, forward-reference the bounds.
- §4.1: pointer to §4.2 and §5 as the technical contribution.
- BLS12-381 security level stated explicitly; scoping sentence on physical attacks.

---

We note that this submission is a Major Revision, and that the binding revision criteria
from the previous round are confirmed as met by Reviewers A and B.
