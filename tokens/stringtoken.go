package tokens

import (
	"fmt"
)

// StringToken repressents a token that contains text.
type StringToken struct {
	Value string
}

// Type returns the StringTokenType.
func (StringToken) Type() TokenType {
	return StringTokenType
}

// Eval returns the contained string.
func (s StringToken) Eval(param EvalParam) (result EvalResult) {
	result.String = s.Value
	return
}

func (s StringToken) String() string {
	return fmt.Sprintf("StringToken: \"%s\"", s.Value)
}
