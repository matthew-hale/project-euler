package main

import (
    "strconv"
    "math/big"
    "os"
)

/*
The masks are hard-coded here, but we'll implement a function
later to generate them (via simple binary counting/conversion).

Masks are indexed as such:

    The number of digits minus 1 (3 digit masks, x == 2)
          |
          |
          v
    masks[x][y]
             ^
             |
             |
    The index of each digit in the mask, continuously, from the rightmost digit (0 == 1's place).

For performance reasons, we use a 2d slice, and just stack up all the masks back-to-back.
We know the width of each mask, so this works out nicely.

A true value means the digit is to be replaced, while a false value means it stays as is.
*/

var masks [][]bool
//Single digit doesn't need a mask
masks[0][0] = true

//2 digits, 3 total masks (2 types of single digit replacement, 1 type of double)
masks[1] = {
    true, false,
    false, true,
    true, true,
}

//3 digits
masks[2] = {
    true, false, false,
    false, true, false,
    false, false, true,
    true, true, false,
    true, false, true,
    false, true, true,
    true, true, true,
}

//4 digits
masks[3] = {
    false, false, false, true,
    false, false, true, false,
    false, true, false, false,
    true, false, false, false,
    false, false, true, true,
    false, true, false, true,
    false, true, true, false,
    true, false, false, true,
    true, false, true, false,
    true, true, false, false,
    false, true, true, true,
    true, false, true, true,
    true, true, false, true,
    true, true, true, false,
    true, true, true, true,
}

//5 digits
masks[4] = {
    false, false, false, false, true,
    false, false, false, true, false,
    false, false, false, true, true,
    false, false, true, false, false,
    false, false, true, false, true,
    false, false, true, true, false,
    false, false, true, true, true,
    false, true, false, false, false,
    false, true, false, false, true,
    false, true, false, true, false,
    false, true, false, true, true,
    false, true, true, false, false,
    false, true, true, false, true,
    false, true, true, true, false,
    false, true, true, true, true,
    true, false, false, false, false,
    true, false, false, false, true,
    true, false, false, true, false,
    true, false, false, true, true,
    true, false, true, false, false,
    true, false, true, false, true,
    true, false, true, true, false,
    true, false, true, true, true,
    true, true, false, false, false,
    true, true, false, false, true,
    true, true, false, true, false,
    true, true, false, true, true,
    true, true, true, false, false,
    true, true, true, false, true,
    true, true, true, true, false,
    true, true, true, true, true,
}

//6 digits
masks[5] = {
    false, false, false, false, false, false,
    false, false, false, false, false, true,
    false, false, false, false, true, false,
    false, false, false, false, true, true,
    false, false, false, true, false, false,
    false, false, false, true, false, true,
    false, false, false, true, true, false,
    false, false, false, true, true, true,
    false, false, true, false, false, false,
    false, false, true, false, false, true,
    false, false, true, false, true, false,
    false, false, true, false, true, true,
    false, false, true, true, false, false,
    false, false, true, true, false, true,
    false, false, true, true, true, false,
    false, false, true, true, true, true,
    false, true, false, false, false, false,
    false, true, false, false, false, true,
    false, true, false, false, true, false,
    false, true, false, false, true, true,
    false, true, false, true, false, false,
    false, true, false, true, false, true,
    false, true, false, true, true, false,
    false, true, false, true, true, true,
    false, true, true, false, false, false,
    false, true, true, false, false, true,
    false, true, true, false, true, false,
    false, true, true, false, true, true,
    false, true, true, true, false, false,
    false, true, true, true, false, true,
    false, true, true, true, true, false,
    false, true, true, true, true, true,
    true, false, false, false, false, false,
    true, false, false, false, false, true,
    true, false, false, false, true, false,
    true, false, false, false, true, true,
    true, false, false, true, false, false,
    true, false, false, true, false, true,
    true, false, false, true, true, false,
    true, false, false, true, true, true,
    true, false, true, false, false, false,
    true, false, true, false, false, true,
    true, false, true, false, true, false,
    true, false, true, false, true, true,
    true, false, true, true, false, false,
    true, false, true, true, false, true,
    true, false, true, true, true, false,
    true, false, true, true, true, true,
    true, true, false, false, false, false,
    true, true, false, false, false, true,
    true, true, false, false, true, false,
    true, true, false, false, true, true,
    true, true, false, true, false, false,
    true, true, false, true, false, true,
    true, true, false, true, true, false,
    true, true, false, true, true, true,
    true, true, true, false, false, false,
    true, true, true, false, false, true,
    true, true, true, false, true, false,
    true, true, true, false, true, true,
    true, true, true, true, false, false,
    true, true, true, true, false, true,
    true, true, true, true, true, false,
    true, true, true, true, true, true,
}

//7 digits
masks[6] = {
    false, false, false, false, false, false, true,
    false, false, false, false, false, true, false,
    false, false, false, false, false, true, true,
    false, false, false, false, true, false, false,
    false, false, false, false, true, false, true,
    false, false, false, false, true, true, false,
    false, false, false, false, true, true, true,
    false, false, false, true, false, false, false,
    false, false, false, true, false, false, true,
    false, false, false, true, false, true, false,
    false, false, false, true, false, true, true,
    false, false, false, true, true, false, false,
    false, false, false, true, true, false, true,
    false, false, false, true, true, true, false,
    false, false, false, true, true, true, true,
    false, false, true, false, false, false, false,
    false, false, true, false, false, false, true,
    false, false, true, false, false, true, false,
    false, false, true, false, false, true, true,
    false, false, true, false, true, false, false,
    false, false, true, false, true, false, true,
    false, false, true, false, true, true, false,
    false, false, true, false, true, true, true,
    false, false, true, true, false, false, false,
    false, false, true, true, false, false, true,
    false, false, true, true, false, true, false,
    false, false, true, true, false, true, true,
    false, false, true, true, true, false, false,
    false, false, true, true, true, false, true,
    false, false, true, true, true, true, false,
    false, false, true, true, true, true, true,
    false, true, false, false, false, false, false,
    false, true, false, false, false, false, true,
    false, true, false, false, false, true, false,
    false, true, false, false, false, true, true,
    false, true, false, false, true, false, false,
    false, true, false, false, true, false, true,
    false, true, false, false, true, true, false,
    false, true, false, false, true, true, true,
    false, true, false, true, false, false, false,
    false, true, false, true, false, false, true,
    false, true, false, true, false, true, false,
    false, true, false, true, false, true, true,
    false, true, false, true, true, false, false,
    false, true, false, true, true, false, true,
    false, true, false, true, true, true, false,
    false, true, false, true, true, true, true,
    false, true, true, false, false, false, false,
    false, true, true, false, false, false, true,
    false, true, true, false, false, true, false,
    false, true, true, false, false, true, true,
    false, true, true, false, true, false, false,
    false, true, true, false, true, false, true,
    false, true, true, false, true, true, false,
    false, true, true, false, true, true, true,
    false, true, true, true, false, false, false,
    false, true, true, true, false, false, true,
    false, true, true, true, false, true, false,
    false, true, true, true, false, true, true,
    false, true, true, true, true, false, false,
    false, true, true, true, true, false, true,
    false, true, true, true, true, true, false,
    false, true, true, true, true, true, true,
    true, false, false, false, false, false, false,
    true, false, false, false, false, false, true,
    true, false, false, false, false, true, false,
    true, false, false, false, false, true, true,
    true, false, false, false, true, false, false,
    true, false, false, false, true, false, true,
    true, false, false, false, true, true, false,
    true, false, false, false, true, true, true,
    true, false, false, true, false, false, false,
    true, false, false, true, false, false, true,
    true, false, false, true, false, true, false,
    true, false, false, true, false, true, true,
    true, false, false, true, true, false, false,
    true, false, false, true, true, false, true,
    true, false, false, true, true, true, false,
    true, false, false, true, true, true, true,
    true, false, true, false, false, false, false,
    true, false, true, false, false, false, true,
    true, false, true, false, false, true, false,
    true, false, true, false, false, true, true,
    true, false, true, false, true, false, false,
    true, false, true, false, true, false, true,
    true, false, true, false, true, true, false,
    true, false, true, false, true, true, true,
    true, false, true, true, false, false, false,
    true, false, true, true, false, false, true,
    true, false, true, true, false, true, false,
    true, false, true, true, false, true, true,
    true, false, true, true, true, false, false,
    true, false, true, true, true, false, true,
    true, false, true, true, true, true, false,
    true, false, true, true, true, true, true,
    true, true, false, false, false, false, false,
    true, true, false, false, false, false, true,
    true, true, false, false, false, true, false,
    true, true, false, false, false, true, true,
    true, true, false, false, true, false, false,
    true, true, false, false, true, false, true,
    true, true, false, false, true, true, false,
    true, true, false, false, true, true, true,
    true, true, false, true, false, false, false,
    true, true, false, true, false, false, true,
    true, true, false, true, false, true, false,
    true, true, false, true, false, true, true,
    true, true, false, true, true, false, false,
    true, true, false, true, true, false, true,
    true, true, false, true, true, true, false,
    true, true, false, true, true, true, true,
    true, true, true, false, false, false, false,
    true, true, true, false, false, false, true,
    true, true, true, false, false, true, false,
    true, true, true, false, false, true, true,
    true, true, true, false, true, false, false,
    true, true, true, false, true, false, true,
    true, true, true, false, true, true, false,
    true, true, true, false, true, true, true,
    true, true, true, true, false, false, false,
    true, true, true, true, false, false, true,
    true, true, true, true, false, true, false,
    true, true, true, true, false, true, true,
    true, true, true, true, true, false, false,
    true, true, true, true, true, false, true,
    true, true, true, true, true, true, false,
    true, true, true, true, true, true, true,
}

//Generate primes using the sieve of eratosthenes method
//Estimated maximum of 7 digits for this problem
func eratosthenes(max int) (primeBooleans []bool) {
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


    return
}

//Generate families for each prime, by replacing digits
func primeFamilies(x int) (families [][]int) {
    //Convert the int to a string, then a slice of ints
    str := strconv.Itoa(x)
    var ints []int
    for _, char := range str {
        digit, _ := strconv.Atoi(string(char))
        ints = append(ints, digit)
    }

    //Pick the mask set index (the first number) based on the length of the integer
    //Additionally determine the total size of the set of masks
    length := len(str)
    maskIndex := length-1
    setLength := len(masks[maskIndex])

    //Generate a family for each mask in the set
    for i := range masks[maskIndex] {
        //We need to grab the current working mask from the set
        var currentMask []bool
        for j := 0; j < length; j++ {
            currentMask[j] = masks[maskIndex][i+j]
        }

        for j := range currentMask {
            for k := 0; k <= 9; k++ {
            }
            ints[j]
        }

        i += length
    }

    return
}

func main() {
    //We want to keep this slice as it offers an easy prime check for our prime families
    primeBooleans := eratosthenes(9999999)

    //Iterate over booleans, adding the indices which are true to the list of primes
    var primes []int
    for num := 1; num <= max; num++ {
        if primeBooleans[num] {
            primes = append(primes, num)
        }
    }

    //For each prime, generate its families
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
