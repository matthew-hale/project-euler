package main

import (
    "strconv"
)

//Generate primes using the sieve of eratosthenes method
//Estimated maximum of 7 digits for this problem
func eratosthenes(max int) []int {
    //Create a slice of booleans with indexes from 0-max, and make them all true
    primes := make([]bool, max+1)
    for i := range(primes){primes[i] = true}

    //Set initial value of p
    p := 2

    for p * p <= max {
        //If the integer is not already marked, loop through its multiples and mark them
        //Then, increment p
        if primes[p] == true {
            for i := p * 2; i <= max; i += p {
                primes[i] = false
            }
        }
        p++
    }

    //As this is usually done with 2 as a minimum value, set 0 and 1 to false
    primes[0] = false
    primes[1] = false

    //Iterate over booleans, adding the indices which are true to the list of primes
    var result []int
    for num := 1; num <= max; num++ {
        if primes[num] {
            result = append(result, num)
        }
    }

    return result
}

//Generate families for each prime, by replacing digits
func primeFamilies(x int) []int {
    //Convert the int to a string, then a slice of ints
    str := strconv.Itoa(x)
    var ints []int
    for _, char := range str {
        digit, _ := strconv.Atoi(string(char))
        ints = append(ints, digit)
    }
    return ints
}

func main() {
    primes := eratosthenes(9999999)
    println(primes[len(primes)-1])
    result := primeFamilies(20)
    println(result[0])
    println(result[1])
}
