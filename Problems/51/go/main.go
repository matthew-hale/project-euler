package main

//Generate primes using the sieve of eratosthenes method
//Estimated maximum of 7 digits for this problem
func eratosthenes(max int) []int {
    primes := make([]bool, max+1)
    for i := range(primes){primes[i] = true}
    p := 2
    for p * p <= max {
        if primes[p] == true {
            for i := p * 2; i <= max+1; i += p {
                primes[i] = false
            }
        }
        p++
    }
    primes[0] = false
    primes[1] = false

    var result []int
    for num := 1; num <= max; num++ {
        if primes[num] {
            result = append(result, num)
        }
    }

    return result
}

func main() {
    primes := eratosthenes(100)
    for i := range primes {
        println(primes[i])
    }
}
