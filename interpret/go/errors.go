package main

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
	return err.msg
}

func (err *runtimeError) Error() string {
	return ""
}
