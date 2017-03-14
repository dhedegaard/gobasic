package tokens

// GotoToken is a type of token that changes the program country (PC) when the token is evaulated.
type GotoToken struct{}

// Type returns the GotoTokenType.
func (GotoToken) Type() TokenType {
	return GotoTokenType
}

// Eval changes the program counter to something based on the next token.
func (GotoToken) Eval(param EvalParam) (result EvalResult) {
	linenumberToProgramIndex := *param.LinenumberToProgramIndex

	// Eval the next element in the tokenlist.
	nextParam := param.Copy()
	nextParam.LineTokenPosition++
	res := ((*param.LineTokens)[nextParam.LineTokenPosition]).Eval(nextParam)
	if res.Error != nil {
		result.Error = res.Error
		return
	}

	*param.ProgramCounter = linenumberToProgramIndex[res.Integer] - 1
	return
}

func (GotoToken) String() string {
	return "GotoToken"
}
