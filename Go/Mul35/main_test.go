package main

import "testing"

func TestBruteSolution(t *testing.T){
    if bruteSolution() != 233168 {
        t.Error("the brute solution don't return the expected result")
    }
}

func BenchmarkBruteSolution(b *testing.B){
    for i:=0; i < b.N; i++ {
        bruteSolution()
    }
}


func TestSumSolution(t *testing.T){
    if sumSolution() != 233168 {
        t.Error("the brute solution don't return the expected result")
    }
}

func BenchmarkSumSolution(b *testing.B) {
    
    for i:=0; i < b.N; i++ {
        sumSolution()
    }
}
