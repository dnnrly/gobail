package gobail

import (
	"testing"
)

func noError() error { return nil }

func TestRunDoesNotExit(t *testing.T) {
	Run(noError()).ExitMsg("there is not error")
}
