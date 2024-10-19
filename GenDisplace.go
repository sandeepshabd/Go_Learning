package main

import (
    "fmt"
)

// GenDisplaceFn returns a function that calculates displacement based on time t
func GenDisplaceFn(a, vo, so float64) func(t float64) float64 {
    return func(t float64) float64 {
        return so + vo*t + 0.5*a*t*t
    }
}

func main() {
    var a, vo, so, t float64

    fmt.Print("Enter acceleration (a): ")
    fmt.Scan(&a)

    fmt.Print("Enter initial velocity (vo): ")
    fmt.Scan(&vo)

    fmt.Print("Enter initial displacement (so): ")
    fmt.Scan(&so)

    // Generate the displacement function using the provided acceleration, initial velocity, and displacement
    displacementFn := GenDisplaceFn(a, vo, so)

    fmt.Print("Enter time (t): ")
    fmt.Scan(&t)

    // Calculate the displacement after time t
    displacement := displacementFn(t)

    fmt.Printf("The displacement after %.2f seconds is: %.2f units\n", t, displacement)
}
