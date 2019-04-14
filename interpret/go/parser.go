package main

import (
	"bufio"
	"fmt"
	"io"
)

func parseFile(pr *program, loc *location, file io.Reader) *interpreterError {
	loc.line = 1
	reader := bufio.NewReader(file)

	for {
		token, err, eof := nextToken(loc, reader)
		if err != nil {
			err.loc = loc
			return err
		} else if eof == true {
			return nil
		}

		fmt.Printf("'%s'\n", token)
	}
}
