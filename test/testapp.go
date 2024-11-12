package main

import (
	"errors"
	"os"

	"github.com/dnnrly/gobail"
)

func returnError() error {
	return errors.New("there was an error")
}

func main() {
	switch os.Args[1] {
	case "exit-no-return":
		gobail.Run(returnError()).ExitMsg("exited with an error: %v")
		break
	case "exit-no-return-without-error":
		gobail.Run(returnError()).ExitMsg("exited with an error")
		break
	}
}
