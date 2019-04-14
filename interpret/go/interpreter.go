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

	if err = parse(pr, tokens); err != nil {
        return nil, false, 
    }

	return nil, false, nil
}
