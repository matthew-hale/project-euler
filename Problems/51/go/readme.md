1. Generate primes (sieve of eratosthenes).

Generating primes is an easier problem than testing each number for primeness. Plus, there is a hard maximum for this problem in any given base, even factoring in adjustable prime family size, so eventually that maximum can be baked in.

2. For each prime, generate the families.

    + Replace all digits (except the rightmost, as only 5 possible primes exists with rightmost digit replacement, which is less than the prime family size of 8), and store the numbers.

    For each family, count the number of primes.

    + Compare the numbers to the list of primes we've already generated.

        + If the number of primes = size of prime family, return current prime
