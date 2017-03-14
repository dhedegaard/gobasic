package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"github.com/dhedegaard/gobasic/tokenizer"
	"github.com/dhedegaard/gobasic/tokens"
)

func main() {
	var program []byte
	if len(os.Args) >= 2 {
		// Read program from filename supplied as arg.
		filename := os.Args[len(os.Args)-1]
		var err error
		program, err = ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to read from file %s because: %s\n", filename, err.Error())
			os.Exit(1)
		}
		fmt.Println("Read program from file:", filename)
	} else {
		// Read program from stdin.
		fmt.Println("Reading from stdin")
		var err error
		if program, err = ioutil.ReadAll(os.Stdin); err != nil {
			fmt.Fprintln(os.Stderr, "Unable to read from stdin because:", err.Error())
			os.Exit(1)
		}
	}
	// Tokenize everything.
	fmt.Println("tokenizing")
	tokenizedProgramMap, err := tokenizer.TokenizeProgram(string(program))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error tokenizing program:", err.Error())
		return
	}
	// Sort the program linenumbers.
	sortedTokenizedLinenumbers := make([]int, 0, len(tokenizedProgramMap))
	for key := range tokenizedProgramMap {
		sortedTokenizedLinenumbers = append(sortedTokenizedLinenumbers, key)
	}
	sort.Ints(sortedTokenizedLinenumbers)
	// Map of program line index -> linenumber.
	programIndexToLinenumber := make(map[int]int, len(tokenizedProgramMap))
	// Map of program linenumber -> PC index.
	linenumberToProgramIndex := make(map[int]int, len(tokenizedProgramMap))
	for idx, linenumber := range sortedTokenizedLinenumbers {
		programIndexToLinenumber[idx] = linenumber
		linenumberToProgramIndex[linenumber] = idx
	}
	// Execute the program line for line
	for pc := 0; pc < len(sortedTokenizedLinenumbers); pc++ {
		lineNumber := programIndexToLinenumber[pc]
		programLine, ok := tokenizedProgramMap[lineNumber]
		if !ok {
			fmt.Fprintln(os.Stderr, "Unable to find line for linenumber:", lineNumber)
			return
		}
		for tokenIdx := 0; tokenIdx < len(programLine); tokenIdx++ {
			token := programLine[tokenIdx]
			token.Eval(tokens.EvalParam{
				LineTokens:               &programLine,
				LineTokenPosition:        tokenIdx,
				ProgramCounter:           &pc,
				LinenumberToProgramIndex: &linenumberToProgramIndex,
			})
		}
	}
}
