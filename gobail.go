package gobail

import (
	"fmt"
	"os"
	"strings"
)

type baseEvaluator struct {
	err error
}

func (e baseEvaluator) doExit(msg string) {
	if e.err != nil {
		if strings.Contains(msg, "%v") {
			fmt.Fprintf(os.Stderr, msg, e.err)
		} else {
			fmt.Fprint(os.Stderr, msg)
		}
		os.Exit(1)
	}
}

func (e baseEvaluator) doPanic(msg string) {
	if e.err != nil {
		if strings.Contains(msg, "%v") {
			panic(fmt.Sprintf(msg, e.err))
		} else {
			panic(msg)
		}
	}
}

// RunEvaluator handles functions with no return value
type RunEvaluator struct {
	baseEvaluator
}

// OrExit exits the application without an error message if there is an error
func (e RunEvaluator) OrExit() {
	e.doExit("")
}

// OrExitMsg exits the application with an error message if there is an error
// You may use the %v placeholder to include the error in the message
func (e RunEvaluator) OrExitMsg(msg string) {
	e.doExit(msg)
}

// OrPanic panics without an error message if there is an error
func (e RunEvaluator) OrPanic() {
	e.doPanic("")
}

// OrPanicMsg exits the application with an error message if there is an error
// You may use the %v placeholder to include the error in the message
func (e RunEvaluator) OrPanicMsg(msg string) {
	e.doPanic(msg)
}

// Run collects the error for evaluation in the OrExit or OrPanic methods
func Run(err error) RunEvaluator {
	return RunEvaluator{
		baseEvaluator: baseEvaluator{err},
	}
}

// RunEvaluator handles functions with 1 return value
type ReturnEvaluator[R any] struct {
	baseEvaluator
	ret R
}

// OrExit exits the application without an error message if there is an error, returning the value of the function if there is no error
func (e ReturnEvaluator[R]) OrExit() R {
	e.doExit("")
	return e.ret
}

// OrExitMsg exits the application with an error message if there is an error, returning the value of the function if there is no error
// You may use the %v placeholder to include the error in the message
func (e ReturnEvaluator[R]) OrExitMsg(msg string) R {
	e.doExit(msg)
	return e.ret
}

// OrPanic panics without an error message if there is an error, returning the value of the function if there is no error
func (e ReturnEvaluator[R]) OrPanic() R {
	e.doPanic("")
	return e.ret
}

// OrPanicMsg exits the application with an error message if there is an error, returning the value of the function if there is no error
// You may use the %v placeholder to include the error in the message
func (e ReturnEvaluator[R]) OrPanicMsg(msg string) R {
	e.doPanic(msg)
	return e.ret
}

// Return collects the error and other return value for evaluation in the OrExit or OrPanic methods
func Return[R any](ret R, err error) ReturnEvaluator[R] {
	return ReturnEvaluator[R]{
		baseEvaluator: baseEvaluator{err},
		ret:           ret,
	}
}

// RunEvaluator handles functions with 2 return values
type Return2Evaluator[R, S any] struct {
	baseEvaluator
	ret1 R
	ret2 S
}

// OrExit exits the application without an error message if there is an error, returning the values of the function if there is no error
func (e Return2Evaluator[R, S]) OrExit() (R, S) {
	e.doExit("")
	return e.ret1, e.ret2
}

// OrExitMsg exits the application with an error message if there is an error, returning the values of the function if there is no error
// You may use the %v placeholder to include the error in the message
func (e Return2Evaluator[R, S]) OrExitMsg(msg string) (R, S) {
	e.doExit(msg)
	return e.ret1, e.ret2
}

// OrPanic panics without an error message if there is an error, returning the values of the function if there is no error
func (e Return2Evaluator[R, S]) OrPanic() (R, S) {
	e.doPanic("")
	return e.ret1, e.ret2
}

// OrPanicMsg exits the application with an error message if there is an error, returning the values of the function if there is no error
// You may use the %v placeholder to include the error in the message
func (e Return2Evaluator[R, S]) OrPanicMsg(msg string) (R, S) {
	e.doPanic(msg)
	return e.ret1, e.ret2
}

// Return2 collects the error and other return values for evaluation in the OrExit or OrPanic methods
func Return2[R, S any](ret1 R, ret2 S, err error) Return2Evaluator[R, S] {
	return Return2Evaluator[R, S]{
		baseEvaluator: baseEvaluator{err},
		ret1:          ret1,
		ret2:          ret2,
	}
}
