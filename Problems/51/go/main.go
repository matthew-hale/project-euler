package main

import (
    "strconv"
    "math/big"
    "os"
)

/*
00000000
00000001
00000010
00000011
00000100
00000101
00000110
00000111
00001000
00001001
00001010
00001011
00001100
00001101
00001110
00001111
00010000
00010001
00010010
00010011
00010100
00010101
00010110
00010111
00011000
00011001
00011010
00011011
00011100
00011101
00011110
00011111
00100000
00100001
00100010
00100011
00100100
00100101
00100110
00100111
00101000
00101001
00101010
00101011
00101100
00101101
00101110
00101111
00110000
00110001
00110010
00110011
00110100
00110101
00110110
00110111
00111000
00111001
00111010
00111011
00111100
00111101
00111110
00111111
01000000
01000001
01000010
01000011
01000100
01000101
01000110
01000111
01001000
01001001
01001010
01001011
01001100
01001101
01001110
01001111
01010000
01010001
01010010
01010011
01010100
01010101
01010110
01010111
01011000
01011001
01011010
01011011
01011100
01011101
01011110
01011111
01100000
01100001
01100010
01100011
01100100
01100101
01100110
01100111
01101000
01101001
01101010
01101011
01101100
01101101
01101110
01101111
01110000
01110001
01110010
01110011
01110100
01110101
01110110
01110111
01111000
01111001
01111010
01111011
01111100
01111101
01111110
01111111
*/
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

    /*
    The masks are hard-coded here, but we'll implement a function
    later to generate them.

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
    }

    //Do digit replacement for all possible combinations.
    var families [][]int

    //Pick the mask based on the length of the integer
    switch length := len(str); length {
        case 1:
            for i := 0; i < 10; i++ {
                families[0][i] = i+1
            }
        case 2:
        case 3:
        case 4:
        case 5:
        case 6:
        case 7:
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
