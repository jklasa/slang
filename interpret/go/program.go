package main

const (
	globalStart      = 1
	globalSize       = 200
	stackFramesStart = globalStart + globalSize
	stackFramesSize  = 5000
	stackStart       = stackFramesStart + stackFramesSize
	stackSize        = 300
	heapStart        = stackStart + stackSize
	initialHeapSize  = 6000
	memEnd           = heapStart + initialHeapSize
)

type program struct {
	files map[string]bool
}

func (pr *program) run() *runtimeError {
	return nil
}
