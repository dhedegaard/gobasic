package tokenizer

import (
	"fmt"
	"strconv"
	"strings"

	"errors"

	"github.com/dhedegaard/gobasic/tokens"
)

// NilLinenumber repressents a nil line number.
const NilLinenumber = -1

// Splits the line into a slice of strings without any empty strings.
func splitNoEmpty(line string, sep string) (result []string) {
	for _, elem := range strings.Split(line, sep) {
		if elem != "" {
			result = append(result, elem)
		}
	}
	return
}

func tokenizeLine(line string) (linenumber int, result []tokens.Token, err error) {
	linenumber = NilLinenumber
	// Split on spaces.
	splittedLine := splitNoEmpty(line, " ")
	if len(splittedLine) <= 1 {
		err = errors.New("The program line does not contain any instructions")
		return
	}
	// Parse the linenumber.
	linenumber, err = strconv.Atoi(splittedLine[0])
	if err != nil {
		err = fmt.Errorf("Unable to parse lineNumber from string: %v", splittedLine[0])
		return
	}
	// Iterate on the rest.
	for i := 1; i < len(splittedLine); i++ {
		elem := strings.ToUpper(splittedLine[i])
		switch elem {
		case "PRINT":
			result = append(result, tokens.PrintToken{})
		case "GOTO":
			result = append(result, tokens.GotoToken{})
		default:
			result = append(result, tokens.StringToken{Value: elem})
		}
	}
	// If we didn't parse anything, fail.
	if len(result) == 0 {
		err = fmt.Errorf("Nothing was tokenized from line: %s", line)
		return
	}
	// All done, return the top of the token tree for the current instruction line.
	return
}

// TokenizeProgram takes a given program string and parses it all into a program that maps linenumber -> tokenized line.
func TokenizeProgram(program string) (tokenizedProgram map[int][]tokens.Token, err error) {
	tokenizedProgram = make(map[int][]tokens.Token, 0)

	// Iterate, tokenizing each line and adding it to the program map.
	for _, line := range splitNoEmpty(program, "\n") {
		linenum, tokens, lineErr := tokenizeLine(line)
		if lineErr != nil {
			err = lineErr
			return
		}
		tokenizedProgram[linenum] = tokens
	}
	return
}
