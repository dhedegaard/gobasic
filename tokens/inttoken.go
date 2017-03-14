package tokens

import "fmt"

// IntToken repressents a token with an integer value.
type IntToken struct {
	value int
}

// Type returns the IntTokenType.
func (IntToken) Type() TokenType {
	return IntTokenType
}

// Eval returns the integer value contained in the token.
func (t IntToken) Eval() interface{} {
	return t.value
}

func (t IntToken) String() string {
	return fmt.Sprintf("IntToken: %v", t.value)
}
