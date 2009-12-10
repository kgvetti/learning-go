// 
//  echo command, from: http://github.com/jbrownlee/learning-go
// 

package main

import (
	"os"; 
	"flag";
)

var omitNewLine = flag.Bool("n", false, "don't print final newline")

// can also define const's the c-way: const name = value
const (
	Space = " ";
	NewLine = "\n";
)

func main() {
	// scans the arg list and sets up flags
	flag.Parse();
	var s string = "";
	
	for i := 0; i<flag.NArg(); i++ {
		if i > 0 {
			s += Space
		}
		s += flag.Arg(i);	
	}
	
	if !*omitNewLine {
		s += NewLine
	}
	
	os.Stdout.WriteString(s);
}