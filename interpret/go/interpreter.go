package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("To run a program, use\n\tslang <file>.sl")
		return
	}

	pr := &program{
		files: make(map[string]bool),
	}

	if err := interpretFile(pr, nil, os.Args[1]); err != nil {
		fmt.Println(err)
		return
	}
	pr.run()
}

func interpretFile(pr *program, loc *location, filename string) *interpreterError {
	loc = &location{
		parent:   loc,
		filename: filename,
	}

	file, err := os.Open(filename)
	if err != nil {
		return &interpreterError{
			msg: "Could not open file",
			loc: loc,
		}
	}
	defer file.Close()

	pr.files[filename] = true

	scanner := bufio.NewScanner(file)
	for loc.line = 1; scanner.Scan(); loc.line++ {
		importFile, guarded, err := interpretLine(pr, scanner.Text())
		if err != nil {
			err.loc = loc
			return err
		}

		if importFile != nil {
			if !guarded || (guarded && !pr.files[*importFile]) {
				if *importFile == filename {
					return &interpreterError{
						msg: "Recursive import of file",
						loc: loc,
					}
				}

				if err = interpretFile(pr, loc, *importFile); err != nil {
					return err
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return &interpreterError{
			msg: "Error reading file",
			loc: loc,
		}
	}

	return nil
}

func interpretLine(pr *program, line string) (*string, bool, *interpreterError) {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil, false, nil
	}

	tokens, err := tokenizeLine(line)
	if err != nil {
		return nil, false, err
	}

	for _, token := range tokens {
		fmt.Printf("'%s'", token)
	}
	fmt.Println()

	return nil, false, nil
}

func tokenizeLine(line string) ([]string, *interpreterError) {
	tokens := make([]string, 0)
	var token strings.Builder

	quoted := false
	escaped := false
	label := false
	compound := 0

	for _, char := range line {
		if compound > 0 {
			token.WriteRune(char)

			if char == '[' {
				compound++
			} else if char == ']' {
				compound--

				if compound == 0 {
					tokens = append(tokens, token.String())
					token.Reset()
				}
			} else if char == ';' {
				return nil, &interpreterError{msg: "Comments are not permitted in compound variables"}
			}
			continue
		}

		switch char {
		case '[':
			if !quoted {
				compound++
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
				if token.Len() != 0 {
					if label {
						// Labels
					} else {
						tokens = append(tokens, token.String())
					}
				}
				token.Reset()
			} else {
				token.WriteRune(char)
			}
		case ';':
			continue
		default:
			if escaped {
				if char == 'n' {
					token.WriteRune('\n')
				} else if char == 't' {
					token.WriteRune('\t')
				} else {
					token.WriteRune(char)
				}
				escaped = false
			} else {
				token.WriteRune(char)
			}
		}
	}

	if token.Len() > 0 {
		tokens = append(tokens, token.String())
	}

	return tokens, nil
}
