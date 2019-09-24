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
