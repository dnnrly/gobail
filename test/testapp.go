package main

import (
	"errors"
	"fmt"
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

func return2ValAndError() (int, string, error) {
	return 0, "str", err
}

func main() {
	fmt.Println("Test case is: " + os.Args[1])
	switch os.Args[1] {
	case "exit-no-return":
		gobail.Run(returnError()).ExitMsg("exited with an error: %v")
	case "exit-no-return-without-error":
		gobail.Run(returnError()).ExitMsg("exited with an error")
	case "exit-return":
		gobail.Return(returnValAndError()).ExitMsg("exited with an error: %v")
	case "exit-return-without-error":
		gobail.Return(returnValAndError()).ExitMsg("exited with an error")
	case "exit-return2":
		gobail.Return2(return2ValAndError()).ExitMsg("exited with an error: %v")
	case "exit-return2-without-error":
		gobail.Return2(return2ValAndError()).ExitMsg("exited with an error")
	}
}
