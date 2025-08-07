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
