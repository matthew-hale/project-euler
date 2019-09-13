package main

import (
    "strconv"
    "math/big"
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
    //digits := len(ints)

    return true
}

func main() {
    primes := eratosthenes(9999999)
    test := big.NewInt(0).Binomial(4,2)
    println(test.Int64())
}
