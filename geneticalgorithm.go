// 
// A genetic algorithm in the go programming language
// By Jason Brownlee, 2009

// 
// A binary representation, tournament selection, one point crossover and point mutations
// The algorithm is demonstrated on the OneMax problem.
// 

// uses many go features including types, objects, interfaces (polymorphism), channels, goroutines, and gc (lots of object creation)

package main

import (
	"fmt"; 
	"rand";
	"math";
) 

//
// types
//

type Solution struct {
	bitstring *[]bool;
	score float;
}

type Problem interface {
	Evaluate(sol *Solution) float;
	GetProblemSize() int;
	String() string;
}

type Algorithm interface {
	Run(prob *Problem) string;
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

// solution

func (s *Solution) String() string {
	return fmt.Sprintf("(%v) [%v]", s.score, s.bitstring);
}

// problem

func (problem *OneMax) Evaluate(sol *Solution) float {
	sum := 0.0;
	for i := 0; i<problem.problemSize; i++ {
		if sol.bitstring[i]==true { 
			sum += 1.0;
		}
	}
	return sum;
}

func (problem *OneMax) GetProblemSize() int {
	return problem.problemSize;
}

func (problem *OneMax) String() string {
	return fmt.Sprintf("OneMax, len=%v", problem.problemSize);
}

func (problem *OneMax) NewRandomSolution() *Solution {
	s := new(Solution);
	b := [(*problem).problemSize]bool{};
	s.bitstring = &b;
	s.score = math.NaN;
	// populate with random values
	for i:=0; i<len(s.bitstring); i++ {
		s.bitstring = (rand.Intn(2) == 1);
	}
	return s;
}

func NewOneMax(size int) *OneMax {
	p := new(OneMax);
	p.problemSize = size;
	return p;
}

// algorithm

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
		
		fmt.Printf(" > gen %v, best=%v\n", i, best);
	}
	
	return best;
}

//
// entry point
//
func main() {
	fmt.PrintLine("Booting");	
	
	// seed the random number generator
	// TODO
	
	// var s = bool[10];
	
	// create new objects on the heap
	// algorithm := new(GeneticAlgorithm);
	problem := NewOneMax(64);
	
	s := problem.NewRandomSolution();
	fmt.Printf("%v\n", s);
	
}

