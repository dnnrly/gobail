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

type runEvaluator struct {
	baseEvaluator
}

func (e runEvaluator) OrExit() {
	e.doExit("")
}

func (e runEvaluator) OrExitMsg(msg string) {
	e.doExit(msg)
}

func (e runEvaluator) OrPanic() {
	e.doPanic("")
}

func (e runEvaluator) OrPanicMsg(msg string) {
	e.doPanic(msg)
}

func Run(err error) runEvaluator {
	return runEvaluator{
		baseEvaluator: baseEvaluator{err},
	}
}

type returnEvaluator[R any] struct {
	baseEvaluator
	ret R
}

func (e returnEvaluator[R]) OrExit() R {
	e.doExit("")
	return e.ret
}

func (e returnEvaluator[R]) OrExitMsg(msg string) R {
	e.doExit(msg)
	return e.ret
}

func (e returnEvaluator[R]) OrPanic() R {
	e.doPanic("")
	return e.ret
}

func (e returnEvaluator[R]) OrPanicMsg(msg string) R {
	e.doPanic(msg)
	return e.ret
}

func Return[R any](ret R, err error) returnEvaluator[R] {
	return returnEvaluator[R]{
		baseEvaluator: baseEvaluator{err},
		ret:           ret,
	}
}

type return2Evaluator[R, S any] struct {
	baseEvaluator
	ret1 R
	ret2 S
}

func (e return2Evaluator[R, S]) OrExit() (R, S) {
	e.doExit("")
	return e.ret1, e.ret2
}

func (e return2Evaluator[R, S]) OrExitMsg(msg string) (R, S) {
	e.doExit(msg)
	return e.ret1, e.ret2
}

func (e return2Evaluator[R, S]) OrPanic() (R, S) {
	e.doPanic("")
	return e.ret1, e.ret2
}

func (e return2Evaluator[R, S]) OrPanicMsg(msg string) (R, S) {
	e.doPanic(msg)
	return e.ret1, e.ret2
}

func Return2[R, S any](ret1 R, ret2 S, err error) return2Evaluator[R, S] {
	return return2Evaluator[R, S]{
		baseEvaluator: baseEvaluator{err},
		ret1:          ret1,
		ret2:          ret2,
	}
}
