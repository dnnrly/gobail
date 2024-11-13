package main

import (
	"errors"
	"os"

	"github.com/dnnrly/gobail"
)

var err = errors.New("there was an error")

func returnError() error {
	return err
}

func returnValAndError() (int, error) {
	return 0, err
}

func main() {
	switch os.Args[1] {
	case "exit-no-return":
		gobail.Run(returnError()).ExitMsg("exited with an error: %v")
	case "exit-no-return-without-error":
		gobail.Run(returnError()).ExitMsg("exited with an error")
	case "exit-return":
		gobail.Return(returnValAndError()).ExitMsg("exited with an error: %v")
	case "exit-return-without-error":
		gobail.Return(returnValAndError()).ExitMsg("exited with an error")
	}
}
