package main

import "strings"

type tokenType int

const (
	lstring tokenType = iota
	lint
	lfloat
	function
	instruction
	label
	compound
)

func tokenizeLine(line string) ([]string, *interpreterError) {
	tokens := make([]string, 0)
	var token strings.Builder

	escaped := false
	quoted := false
	expression := 0

charLoop:
	for _, char := range line {
		if expression > 0 {
			token.WriteRune(char)

			if char == '[' {
				expression++
			} else if char == ']' {
				expression--

				if expression == 0 {
					tokens = append(tokens, token.String())
					token.Reset()
				}
			} else if char == ';' {
				return nil, &interpreterError{msg: "Comments are not permitted in variable expressions"}
			}
			continue
		}

		switch char {
		case ';':
			if !quoted {
				break charLoop
			}
			token.WriteRune(char)
		case '[':
			if !quoted {
				expression++
			}
			token.WriteRune(char)
		case '\\':
			if escaped {
				token.WriteRune(char)
			}
			escaped = !escaped
		case '"':
			token.WriteRune(char)
			if !escaped {
				if quoted {
					tokens = append(tokens, token.String())
					token.Reset()
				}
				quoted = !quoted
			}
		case ' ', '\t':
			if !quoted {
				if token.Len() > 0 {
					tokens = append(tokens, token.String())
				}
				token.Reset()
			} else {
				token.WriteRune(char)
			}
		case 'n':
			if escaped {
				token.WriteRune('\n')
				escaped = false
			} else {
				token.WriteRune(char)
			}
		case 't':
			if escaped {
				token.WriteRune('\t')
				escaped = false
			} else {
				token.WriteRune(char)
			}
		default:
			token.WriteRune(char)
		}
	}

	if escaped {
		token.WriteRune('\\')
	}

	if expression > 0 {
		return nil, &interpreterError{msg: "Unclosed variable expression"}
	} else if quoted {
		return nil, &interpreterError{msg: "Unclosed string"}
	}

	if token.Len() > 0 {
		tokens = append(tokens, token.String())
	}

	return tokens, nil
}

func 