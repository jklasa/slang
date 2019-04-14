package main

import (
	"fmt"
	"strings"
)

type location struct {
	parent   *location
	filename string
	line     int
}

type interpreterError struct {
	msg string
	loc *location
}

type runtimeError struct {
	msg string
	loc *location
}

func (err *interpreterError) Error() string {
	msg := &strings.Builder{}

	msg.WriteString("InterpreterError: ")
	msg.WriteString(err.msg)
	addLocation(msg, err.loc)

	return msg.String()
}

func addLocation(msg *strings.Builder, loc *location) {
	if loc.parent != nil {
		addLocation(msg, loc.parent)
	}

	msg.WriteString(fmt.Sprintf("\n\t%s: %d", loc.filename, loc.line))
}

func (err *runtimeError) Error() string {
	return ""
}
