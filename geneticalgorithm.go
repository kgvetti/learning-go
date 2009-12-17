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
	AlgorithmNumGenerations = 50;
)

type Solution struct {
	bitstring []bool;
	score float;
}

func NewSolution() *Solution {
	s := new(Solution);
	s.bitstring = make([]bool, ProblemNumBits);
	s.score = 0;
	return s;
}

func NewRandomSolution() *Solution {
	s := NewSolution();
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

func InitializePopulation() (*list.List) {
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

func EvaluatePopulation(population *list.List) (best *Solution) {	
	c := make(chan *Solution, population.Len()); // with buffer
	// evaluate the population in parallel
	for e := population.Front(); e != nil; e = e.Next() {
		var s = e.Value.(*Solution);
		go func() {Fitness(s);c<-s}();
	}
	// barrier waiting for all evaluations to finish, locate best result
	for i:=0; i<population.Len(); i++ {
		var s = <-c;
		if best==nil || s.score>best.score {
			best = s;
		} 
	}
	
	return best;
}

func Mutate(candidate *Solution) {
	for i:=0; i<len(candidate.bitstring); i++ {
		if rand.Float32() < ProbMutation {
			candidate.bitstring[i] = !candidate.bitstring[i]; // bit flip
		}
	}
}

func MakeChildren(parent1 *Solution, parent2 *Solution) (*Solution, *Solution) {
	child1 := NewSolution();
	child2 := NewSolution();
	// crossover point
	var pt = rand.Intn(len(parent1.bitstring));
	useCrossover := true;
	if rand.Float32()>ProbCrossover {useCrossover = false;}
	// can we use slices?
	for i:=0; i<len(parent1.bitstring); i++ {
		if useCrossover && i<pt {
			child1.bitstring[i] = parent1.bitstring[i];
			child2.bitstring[i] = parent2.bitstring[i];
		} else {
			child1.bitstring[i] = parent2.bitstring[i];
			child2.bitstring[i] = parent1.bitstring[i];
		}
	}
	// perform mutation
	Mutate(child1);
	Mutate(child2);

	return child1, child2;
}

func Reproduce(population *list.List) (*list.List) {
	children := list.New();
	channel := make(chan *Solution, population.Len()); // with buffer
	
	// evaluate the population in parallel
	for e1 := population.Front(); e1 != nil; e1 = e1.Next() {
		var e2 = e1.Next();
		if(e2 == nil) {e2 = population.Front();} // odd pop num case
		// reproduce in parallel
		// go func() {
			child1, child2 := MakeChildren(e1.Value.(*Solution), e2.Value.(*Solution)); 
			channel<-child1; channel<-child2;
		// }();
		// ensure we end safely
		e1 = e2;
	}
	// barrier to collect results
	for i:=0; i<population.Len(); i++ {
		children.PushBack(<-channel);
	}
	
	return children;
}

func Select(population *list.List) (*list.List) {
	selected := list.New();
	channel := make(chan *Solution, population.Len()); // with buffer
	// select
	for i:=0; i<population.Len(); i++ {
		go func() {
			var best *Solution;
			for j:=0; j<SelectionNumBouts; j++ {
				// pick
			 	index := rand.Intn(population.Len());
				var candidate *Solution; var k = 0;
				// this is crap - use a structure that allows random access 
				for e := population.Front(); e != nil; e = e.Next() {
					if(k==index) {
						candidate = e.Value.(*Solution);
						break;
					}
					k++;
				}
				// test
				if(best==nil || candidate.score>best.score) { best = candidate; }
			}
			channel<-best;
		}();
	}
	// barrier to collect results
	for i:=0; i<population.Len(); i++ {
		selected.PushBack(<-channel);
	}
	
	return selected;
}

func Evolve() (best *Solution){
	// initialize the population
	population := InitializePopulation();	
	// evolve the population in a sequence
	for i:=0; i<AlgorithmNumGenerations; i++ {
		// evaluate
		var s = EvaluatePopulation(population);
		if best==nil || s.score>best.score { best = s;}
		fmt.Printf(" > gen %v, best=%v\n", i, best);
		// select
		selected := Select(population);
		// reproduce
		children := Reproduce(selected);
		// replace
		population = children;		
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
