package main

//Generate primes using the sieve of eratosthenes method
//Estimated maximum of 7 digits for this problem
func eratosthenes(max int) (primes []bool) {
    primes = make([]bool, max)
    for i := range(primes){primes[i] = true}
    p := 2
    for p * p <= max {
        if primes[p] == true {
        }
    }

    return
}

func main() {
    primes := eratosthenes(10)
    for i := range(primes){
        if primes[i] == true {
            print("true")
        } else {
            print("false")
        }
    }
}
