package main

import (
	"bufio"
	"io"
	"strings"
)

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

func nextNTokens(count int, loc *location, reader *bufio.Reader) ([]string, *interpreterError) {
	tokens := make([]string, count)

	for i := 0; i < count; i++ {
		token, err, eof := nextToken(loc, reader)
		if err != nil {
			err.loc = loc
			return nil, err
		}

		if eof == true {
			return nil, &interpreterError{
				msg: "Unexpectedly reached end of file",
				loc: loc,
			}
		}

		tokens[i] = token
	}

	return tokens, nil
}

func nextToken(loc *location, reader *bufio.Reader) (string, *interpreterError, bool) {
	var token strings.Builder

	comment := false
	escaped := false
	quoted := false
	variadic := false
	scope := 0
	expression := 0

	for {
		if char, _, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				switch {
				case quoted:
					return "", &interpreterError{msg: "Unclosed string"}, true
				case variadic:
					return "", &interpreterError{msg: "Unclosed variadic"}, true
				case scope > 0:
					return "", &interpreterError{msg: "Unclosed structure scope"}, true
				case expression > 0:
					return "", &interpreterError{msg: "Unclosed variable expression"}, true
				}

				if token.Len() > 0 {
					return token.String(), nil, true
				}

				return "", nil, true
			} else {
				return "", &interpreterError{
					msg: "Error reading file",
				}, false
			}

		} else {
			if char == '\n' {
				loc.line++
			}

			switch {
			case comment:
				if char == '\n' {
					comment = false
				}
			case quoted:
				switch char {
				case '\\':
					if escaped {
						token.WriteRune(char)
					}
					escaped = !escaped
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
				case '"':
					token.WriteRune(char)
					if !escaped {
						return token.String(), nil, false
					}
					escaped = false
				default:
					token.WriteRune(char)
				}

			case variadic:
				token.WriteRune(char)
				switch char {
				case '(':
					return "", &interpreterError{
						msg: "Variadics are not permitted in other variadics",
					}, false
				case ')':
					return token.String(), nil, false
				case ';':
					return "", &interpreterError{
						msg: "Comments are not permitted in variadics",
					}, false
				case '{', '}':
					return "", &interpreterError{
						msg: "Scopes are not permitted in variadics",
					}, false
				}

			case scope > 0:
				token.WriteRune(char)
				if char == '{' {
					scope++
				} else if char == '}' {
					scope--

					if scope == 0 {
						return token.String(), nil, false
					}
				}

			case expression > 0:
				token.WriteRune(char)
				switch char {
				case '[':
					expression++
				case ']':
					expression--
					if expression == 0 {
						return token.String(), nil, false
					}
				case ';':
					return "", &interpreterError{
						msg: "Comments are not permitted in variable expressions",
					}, false
				case '{', '}':
					return "", &interpreterError{
						msg: "Scopes are not permitted in variable expressions",
					}, false
				case '(', ')':
					return "", &interpreterError{
						msg: "Variadics are not permitted in variable expressions",
					}, false
				}

			default:
				switch char {
				case ';':
					comment = true
				case '[':
					token.WriteRune(char)
					expression++
				case '(':
					token.WriteRune(char)
					variadic = true
				case '{':
					token.WriteRune(char)
					scope++
				case '"':
					token.WriteRune(char)
					quoted = true
				case ' ', '\t', '\r', '\n':
					if token.Len() > 0 {
						return token.String(), nil, false
					}
				default:
					token.WriteRune(char)
				}
			}
		}
	}
}
