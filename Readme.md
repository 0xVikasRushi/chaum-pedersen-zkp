# Chaum Pedersen Zkp protocol

### Setup

1. Choose large random prime no = q
2. choose element g such g belongs to Gq (Cyclic group)will the generator
   1. Cyclic Group definition
      1. set of number = {1,.....q-1}
      2. set of number which has satisfy gcd (1, 1,....q-1) = 1 then its cyclic group
3. choose random a,b
4. compute the value <g,A,B,C> and sent the value to prover & verifier
   1. A = g^a mod q
   2. B = g^b mod q
   3. c = g^ab mod q

### Proof Generation

prover wants to prove the secret x to verifer then

1. step 1: Compute Y1, Y2
   - Y1 =g^x mod q
   - Y2 = B^x mode q
2. choose a random number s = 300
   1. z = (x + a\*s) mod q

### Verification

1. g^z mod q = A^s \* Y1 mod q
2. B^z mod q = C^s \* Y2 mod q
