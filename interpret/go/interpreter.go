package main

import (
	"fmt"
	"os"
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

	return parseFile(pr, loc, file)
}

/*
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
*/
