/**
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.

I implement 3 type of solutions and add the benchmark for all of then, the arithmetic solution prove that the mathematics are
amazing!
*/
package main

import "fmt"

func main() {
    fmt.Println(arithmeticSolution())
    fmt.Println(bruteSolution())
    fmt.Println(sumSolution())
}

func arithmeticSolution() int32 {
    t3 := arithmeticProgression(3, 999) 
    t5 := arithmeticProgression(5, 995)
    t3sub := arithmeticProgression(15, 990)
    return t3 - t3sub + t5 
}

// Use the arithmetic progression sum equation to get the sum of the secuences
// see https://en.wikipedia.org/wiki/Arithmetic_progression
func arithmeticProgression(d int32, last int32) int32{
    n := (last - d)/d + 1
    return int32(float32(n) / 2 * float32(d + last))
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
