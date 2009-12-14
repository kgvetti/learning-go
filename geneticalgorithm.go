// 
// A genetic algorithm in the go programming language
// By Jason Brownlee, 2009

// 
// A binary representation, tournament selection, one point crossover and point mutations
// The algorithm is demonstrated on the OneMax problem.
// 

// This application is an attempt at getting familar with the syntax and with goroutnes/channels.
// The application is written using a procedural methodology.

package main

import (
	"fmt"; 
	"rand";
	// "math";
	"time";
) 


// algorithm configuration
const (
	ProblemNumBits = 64;
	ProbCrossover = 0.98;
	ProbMutation = 1.0/ProblemNumBits;
	SelectionNumBouts = 3;
	PopulationSize = 50;
	AlgorithmNumGenerations = 100;
)


func NewRandomBitstring() *[]bool {
	var bitstring []bool = make([]bool, 100);
	// populate with random values
	for i:=0; i<len(bitstring); i++ {
		// rand in [0,1]
		if rand.Intn(2) == 1 {
			bitstring[i] = true;
		}
	}
	return &bitstring;
}

func BitstringToString(bitstring *[]bool) (s string) {
	for i:=0; i<len(*bitstring); i++ {
		// no ternary 
		if (*bitstring)[i] {
			s+="1";
		}else{
			s+="0";
		}
	}
	return s;
}


// entry point
func main() {
	fmt.Printf("Starting...\n");	
	// set the random seed
	seed := time.LocalTime().Seconds();
	rand.Seed(seed);
	// run the optimization and get the best solution
	// TODO
	best := NewRandomBitstring();
	// display the result
	fmt.Printf("Finished!\nThe Best solution is: %v\n", BitstringToString(best));
}
