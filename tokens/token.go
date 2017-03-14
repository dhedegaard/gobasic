package tokens

// TokenType is an iota/Enum of the type of a given token.
type TokenType int

const (
	// IntTokenType is a token with an integer,
	IntTokenType = TokenType(iota)
	// FloatTokenType is a token with a floating point number.
	FloatTokenType
	// StringTokenType contains some text, as a string.
	StringTokenType
	// GotoTokenType is used for tokens that GOTO somewhere, ie changes the program counter.
	GotoTokenType
	// PrintTokenType prints something to stdout.
	PrintTokenType
)

// Token is the generic interface for all token types.
type Token interface {
	// Returns the type enum for the token type.
	Type() TokenType
	// Executes the given token with any or no param, returning anything or nothing, depending of the token type.
	Eval(param EvalParam) EvalResult
	// Repressents the token type and perhaps value in a readable way.
	String() string
}

// EvalParam contains everything supplied to Eval() calls on tokens during execution.
type EvalParam struct {
	LineTokens               *[]Token     // Tokens for the current line
	LineTokenPosition        int          // Current position in LikeTokens.
	ProgramCounter           *int         // Reference to the PC
	LinenumberToProgramIndex *map[int]int // Map of linenumber -> PC index
}

// Copy returns a shallow copy of the param struct.
func (param EvalParam) Copy() EvalParam {
	return EvalParam{
		LineTokens:               param.LineTokens,
		LineTokenPosition:        param.LineTokenPosition,
		ProgramCounter:           param.ProgramCounter,
		LinenumberToProgramIndex: param.LinenumberToProgramIndex,
	}
}

// EvalResult is return from calls to Token.Eval, remember to check for errors.
type EvalResult struct {
	String  string
	Integer int
	Error   error
}
