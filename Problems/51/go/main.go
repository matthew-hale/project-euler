package main

import (
    "strconv"
    "math/big"
    "os"
)

//Generate primes using the sieve of eratosthenes method
//Estimated maximum of 7 digits for this problem
func eratosthenes(max int) (primes []int) {
    //Create a slice of booleans with indexes from 0-max, and make them all true
    primeBooleans := make([]bool, max+1)
    for i := range(primeBooleans){primeBooleans[i] = true}

    //Set initial value of p; the method starts with a p of 2, as using 1 would mark all the integers
    p := 2

    //As this is usually done with 2 as a minimum value, set 0 and 1 to false
    primeBooleans[0] = false
    primeBooleans[1] = false

    for p * p <= max {
        //If the integer is not already marked, loop through its multiples and mark them
        //Then, increment p
        if primeBooleans[p] == true {
            for i := p * 2; i <= max; i += p {
                primeBooleans[i] = false
            }
        }
        p++
    }

    //Iterate over booleans, adding the indices which are true to the list of primes
    for num := 1; num <= max; num++ {
        if primeBooleans[num] {
            primes = append(primes, num)
        }
    }

    return
}

//Generate families for each prime, by replacing digits
func primeFamilies(x int) bool {
    //Convert the int to a string, then a slice of ints
    str := strconv.Itoa(x)
    var ints []int
    for _, char := range str {
        digit, _ := strconv.Atoi(string(char))
        ints = append(ints, digit)
    }

    //Do digit replacement for all possible combinations
    /*
    First, we have to determine all the possible digit replacement combinations we can perform.
    This is dependent upon the number of digits in the prime.
    For a 5 digit number, we have to get the binomial coefficient of all nCk values:
        5 choose 1
        5 choose 2
        5 choose 3
        5 choose 4
        5 choose 5
    Then we have to actually get those combinations.
        e.g. for 5 choose 2, there are 10 combinations; we have to generate all 10, ensuring there are no duplicates
    */
    n := len(ints)
    for i := 1; i <= n; i++ {
        binomial := big.Binomial(n,i)
        for j := 1; j <= binomial; j++ {
            //binomial = number of combinations
            //i = number of digits we are choosing
            //n = total number of digits in the prime
            //we want a slice of bools to represent a pattern
            var pattern []bool
        }
    }

    return true
}

func main() {
    primes := eratosthenes(9999999)

    for prime in range primes {
        families := primeFamilies(prime)
        for family in range families {
            count := primeCount(family, primes)
            if count == 8 {
                println(prime)
                os.exit(0)
            }
        }
    }

    println("No such prime found; adjust max value")
}
