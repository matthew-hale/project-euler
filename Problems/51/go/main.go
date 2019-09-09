package main

import (
)

//Function to create an arbitrary range of integers
//We'll be giving it fixed values in main
func makeRange(min, max int) (integers []int) {
    integers := make([int], max-1)
    for i := range integers {
        integers[i] = i + min
    }
    return
}

//Generate primes using the sieve of eratosthenes method
//Estimated maximum of 7 digits for this problem
func eratosthenes([]int integers) (primes []int) {
}

func main() {
}
