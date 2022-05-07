package main

import "fmt"

/**
 * If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
 * Find the sum of all the multiples of 3 or 5 below 1000.
*/
func main() {
    fmt.Println(bruteSolution())
    fmt.Println(sumSolution())
}

func sumSolution() int32{
    var i3, i5, t3, t5 int32 = 3, 5, 0, 0
    for {
        if i3 < 1000 && !(i3 % 5 == 0){
            t3 += i3
        }
        if i5 < 1000 {
            t5 += i5 
        }
        i3 += 3
        i5 += 5
        if i3 >= 1000 && i5 >= 1000 {
            break
        }
    }

    return t3 + t5
}

func bruteSolution() int {
    total := 0
    for i := 0; i < 1000; i++ {
        if i % 3 == 0 || i % 5 == 0 {
            total += i
        }
    }
    return total

}
