package tokens

import (
	"fmt"
)

// PrintToken is a token that prints something, usually the next token in a list.
type PrintToken struct {
	nextToken *Token
}

// Type returns the PrintTokenType
func (PrintToken) Type() TokenType {
	return PrintTokenType
}

// Eval takes the next token after this one and returns nil.
func (PrintToken) Eval(param EvalParam) (result EvalResult) {
	nextParam := param.Copy()
	nextParam.LineTokenPosition++

	// Eval the next token.
	resp := (*param.LineTokens)[nextParam.LineTokenPosition].Eval(nextParam)
	if resp.Error != nil {
		result.Error = resp.Error
		return
	}
	fmt.Println(resp.String)
	return
}

func (p PrintToken) String() string {
	return "PrintToken"
}
