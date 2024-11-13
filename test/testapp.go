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
	// Exit
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

	// Panic
	case "panic-no-return":
		gobail.Run(returnError()).PanicMsg("exited with an error: %v")
	case "panic-no-return-without-error":
		gobail.Run(returnError()).PanicMsg("exited with an error")
	case "panic-return":
		gobail.Return(returnValAndError()).PanicMsg("exited with an error: %v")
	case "panic-return-without-error":
		gobail.Return(returnValAndError()).PanicMsg("exited with an error")
	case "panic-return2":
		gobail.Return2(return2ValAndError()).PanicMsg("exited with an error: %v")
	case "panic-return2-without-error":
		gobail.Return2(return2ValAndError()).PanicMsg("exited with an error")
	}
}
