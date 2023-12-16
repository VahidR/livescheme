package livescheme

type tokenKind uint

const (
	// e.g. "(", ")"
	syntaxToken tokenKind = iota
	// e.g. "1", "13"
	integerToken
	// e.g. "+", "define"
	identifierToken
)

type token struct {
	value string
	kind  tokenKind
}

// for example: "(+ 13 2)
// lex(" (+ 13  2 ) ") -> should produce ["(", "+", "13", "2", ")"]
func lex(source []rune) []token {
	var token []token
	for {
		// eat whitespaces
		// check for syntaxToken
		// or check for integerToken

	}
}
