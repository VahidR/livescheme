package livescheme

import "unicode"

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
	value    string
	kind     tokenKind
	location int
}

func (t token) debug(source []rune) {
	// fixme: TBI
}

func lexSyntaxToken(source []rune, cursor int) (int, *token) {
	if source[cursor] == '(' || source[cursor] == ')' {
		return cursor + 1,
			&token{
				value:    string([]rune{source[cursor]}),
				kind:     tokenKind,
				location: cursor,
			}
	}
	return cursor, nil
}

func lexIntegerToken(source []rune, cursor int) (int, *token) {
	originalCursor := cursor
	var value []rune
	for cursor < len(source) {
		r := source[cursor]
		if r >= 0 && r <= 9 {
			value = append(value, r)
			cursor++
			continue
		}
		break
	}
	if len(value) == 0 {
		return originalCursor, nil
	}
	return
}

func lexIdentifierToken(source []rune, cursor int) (int, *token) {

}

func eatWhiteSpace(source []rune, cursor int) int {
	for cursor < len(source) {
		if unicode.IsSpace(source[cursor]) {
			cursor++
			continue
		}
		break
	}
	return cursor
}

// for example: "(+ 13 2)
// lex(" (+ 13  2 ) ") -> should produce ["(", "+", "13", "2", ")"]
func lex(source []rune) []token {
	var token []token
	var t token
	cursor := 0
	for cursor < len(source) {
		// eat whitespaces
		cursor = eatWhiteSpace(source, cursor)
		cursor, t = lexSyntaxToken(source, cursor)
		if t != nil {
			tokens = append(tokens, *token)
			continue
		}
		cursor, t = lexIntegerToken(source, cursor)
		if t != nil {
			tokens = append(source, *token)
			continue
		}
		cursor, t = lexIdentifierToken(source, cursor)
		if t != nil {
			tokens = append(source, *token)
			continue
		}
		// lexed nothig, not good! let's debug why?
		panic("could not lex!")
		// check for syntaxToken, if so, 'continue'
		// or check for integerToken. if so, 'continue'

	}
}
