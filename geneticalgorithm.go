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
	"container/list";
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

type Solution struct {
	bitstring []bool;
	score float;
}


func NewRandomSolution() *Solution {
	s := new(Solution);
	s.bitstring = make([]bool, ProblemNumBits);
	s.score = 0;
	// populate with random values
	for i:=0; i<len(s.bitstring); i++ {
		// rand in [0,1]
		if rand.Intn(2) == 1 {
			s.bitstring[i] = true;
		}
	}
	return s;
}

func (s *Solution) String() string {
	return fmt.Sprintf("(%v) [%v]", s.score, BitstringToString(&s.bitstring));
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

func OneMax(bitstring *[]bool) (score float) {
	for _, v := range *bitstring {
		if v { score += 1.0; }
	}
	return score;
}

func Fitness(s *Solution) {
	s.score = OneMax(&s.bitstring);
}

func InitializePopulation() (pop *list.List) {
	population := list.New();
	
	c := make(chan *Solution, PopulationSize); // with buffer
	// create all of the population in parallel
	for i:=0; i<PopulationSize; i++ {
		go func() {c<-NewRandomSolution();}();
	}
	// barrier to collect results
	for i:=0; i<PopulationSize; i++ {
		population.PushBack(<-c);
	}
	
	return population;
}

func EvaluatePopulation(pop *list.List) (best *Solution) {	
	c := make(chan *Solution, PopulationSize); // with buffer
	// evaluate the population in parallel
	for e := pop.Front(); e != nil; e = e.Next() {
		var s = e.Value.(*Solution);
		go func() {Fitness(s);c<-s}();
	}
	// barrier waiting for all evaluations to finish, locate best result
	for i:=0; i<PopulationSize; i++ {
		var s = <-c;
		if best==nil || s.score>best.score {
			best = s;
		} 
	}
	
	return best;
}

func Evolve() (best *Solution){
	// initialize the population
	population := InitializePopulation();	
	// evolve the population in a sequence
	for i:=0; i<AlgorithmNumGenerations; i++ {
		// evaluate
		var s = EvaluatePopulation(population);
		if best==nil || s.score>best.score {
			best = s;
		}
		// select
		
		// reproduce
		
		// replace
	}
	
	return best;
}

// entry point
func main() {
	// set the random seed
	seed := time.LocalTime().Seconds();
	rand.Seed(seed);
	// run the optimization and get the best solution
	fmt.Printf("Starting...[seed=%v]\n", seed);	
	best := Evolve();
	// display the result
	fmt.Printf("Finished!\nThe best solution is: %v\n", best);
}
