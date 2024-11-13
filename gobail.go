package gobail

import (
	"fmt"
	"os"
	"strings"
)

type runEvaluator struct {
	err error
}

func (e runEvaluator) ExitMsg(msg string) {
	if e.err != nil {
		if strings.Contains(msg, "%v") {
			fmt.Fprintf(os.Stderr, msg, e.err)
		} else {
			fmt.Fprint(os.Stderr, msg)
		}
		os.Exit(1)
	}
}

func Run(err error) runEvaluator {
	return runEvaluator{err}
}

type returnEvaluator[R any] struct {
	err error
	ret R
}

func (e returnEvaluator[R]) ExitMsg(msg string) R {
	if e.err != nil {
		if strings.Contains(msg, "%v") {
			fmt.Fprintf(os.Stderr, msg, e.err)
		} else {
			fmt.Fprint(os.Stderr, msg)
		}
		os.Exit(1)
	}

	return e.ret
}

func Return[R any](ret R, err error) returnEvaluator[R] {
	return returnEvaluator[R]{
		err: err,
		ret: ret,
	}
}
