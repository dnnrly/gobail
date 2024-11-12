package gobail

import (
	"fmt"
	"os"
	"strings"
)

type evaluator struct {
	err error
}

func (e evaluator) ExitMsg(msg string) {
	if e.err != nil {
		if strings.Contains(msg, "%v") {
			fmt.Fprintf(os.Stderr, msg, e.err)
		} else {
			fmt.Fprint(os.Stderr, msg)
		}
		os.Exit(1)
	}
}

func Run(err error) evaluator {
	return evaluator{err}
}
