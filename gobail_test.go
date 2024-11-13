package gobail

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func noError() error { return nil }

func TestRunDoesNotExit(t *testing.T) {
	Run(noError()).ExitMsg("there is not error")
}

func noReturnError() (string, error) { return "some string", nil }

func TestReturnDoesNotExit(t *testing.T) {
	result := Return(noReturnError()).ExitMsg("there is not error")
	assert.Equal(t, result, "some string")
}
