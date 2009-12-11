// 
// A genetic algorithm in the go programming language
// By Jason Brownlee, 2009

// 
// A binary representation, tournament selection, one point crossover and point mutations
// The algorithm is demonstrated on the OneMax problem.
// 


package main

import (
	"os"; 
	"fmt"; // printf 
) 

//
// types
//

type Solution struct {
	bitstring []bool;
	score float;
}

type Problem interface {
	Evaluate(candidate Solution) float;
	GetProblemSize() int;
	String() string;
}

type Algorithm interface {
	Run(p Problem) string;
	String() string;
}

type GeneticAlgorithm struct {
	crossover float;
	mutation float;
	bouts int;
	populationSize int;
	generations int;
}

type OneMax struct {
	problemSize int;
}

//
// functions
//

func (problem *OneMax) Evaluate(candidate []bool) float {
	sum := 0.0;
	for i := 0; i<problem.problemSize; i++ {
		if candidate[i] { sum += 1.0;}
	}
	return sum;
}

func (problem *OneMax) GetProblemSize() int {
	return problem.problemSize;
}

func (problem *OneMax) String() string {
	return fmt.Sprintf("OneMax, len=%v", problem.problemSize);
}


func (algorithm *GeneticAlgorithm) String() string {
	return fmt.Sprintf("GeneticAlgorithm, cros=%v, muta=%v, bout=%v, size=%v, gen=%v", 
		algorithm.crossover, algorithm.mutation, algorithm.bouts, algorithm.populatonSize, algorithm.generations);
}

func (algorithm *GeneticAlgorithm) Run(problem Problem) string {
	best := nil;
	
	// initialize
	
	for i:=0; i<algorithm.generations; i++ {
		// evaluate
		
		// select
		
		// reproduce
		
		// replace
		
		fmt.Printf(" > gen %v, best=%v [%v]\n", i, problem.Evaluate(best), best);
	}
	
	return best;
}

//
// entry point
//
func main() {
	
	// create new objects on the heap
	algorithm := new(GeneticAlgorithm);
	problem := new(OneMax);
	
	fmt.Printf("Hello, World\n");
}

