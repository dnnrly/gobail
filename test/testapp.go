package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/dnnrly/gobail"
)

var err = errors.New("there was an error")

func returnError() error                        { return err }
func returnValAndError() (int, error)           { return 0, err }
func return2ValsAndError() (int, string, error) { return 0, "str", err }

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No arguments provided")
		os.Exit(0)
	}
	fmt.Println("Test case is: " + os.Args[1])

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered panic:", r)
			os.Exit(1)
		}
	}()

	switch os.Args[1] {
	// Exit
	case "exit-no-return-no-msg":
		gobail.Run(returnError()).OrExit()
	case "exit-return-no-msg":
		gobail.Run(returnError()).OrExit()
	case "exit-return2-no-msg":
		gobail.Run(returnError()).OrExit()
	case "exit-no-return":
		gobail.Run(returnError()).OrExitMsg("exited with an error: %v")
	case "exit-no-return-without-error":
		gobail.Run(returnError()).OrExitMsg("exited with an error")
	case "exit-return":
		gobail.Return(returnValAndError()).OrExitMsg("exited with an error: %v")
	case "exit-return-without-error":
		gobail.Return(returnValAndError()).OrExitMsg("exited with an error")
	case "exit-return2":
		gobail.Return2(return2ValsAndError()).OrExitMsg("exited with an error: %v")
	case "exit-return2-without-error":
		gobail.Return2(return2ValsAndError()).OrExitMsg("exited with an error")

	// Panic
	case "panic-no-return-no-msg":
		gobail.Run(returnError()).OrPanic()
	case "panic-return-no-msg":
		gobail.Run(returnError()).OrPanic()
	case "panic-return2-no-msg":
		gobail.Run(returnError()).OrPanic()
	case "panic-no-return":
		gobail.Run(returnError()).OrPanicMsg("exited with an error: %v")
	case "panic-no-return-without-error":
		gobail.Run(returnError()).OrPanicMsg("exited with an error")
	case "panic-return":
		gobail.Return(returnValAndError()).OrPanicMsg("exited with an error: %v")
	case "panic-return-without-error":
		gobail.Return(returnValAndError()).OrPanicMsg("exited with an error")
	case "panic-return2":
		gobail.Return2(return2ValsAndError()).OrPanicMsg("exited with an error: %v")
	case "panic-return2-without-error":
		gobail.Return2(return2ValsAndError()).OrPanicMsg("exited with an error")
	}
}
