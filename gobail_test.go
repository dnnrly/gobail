package gobail

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func noError() error                       { return nil }
func noReturnError() (string, error)       { return "some string", nil }
func noReturn2Error() (string, int, error) { return "some string", 9, nil }

func TestRunDoesNotExit(t *testing.T) {
	Run(noError()).ExitMsg("there is not error")
}

func TestReturnDoesNotExit(t *testing.T) {
	result := Return(noReturnError()).ExitMsg("there is not error")
	assert.Equal(t, result, "some string")
}

func TestReturn2DoesNotExit(t *testing.T) {
	result1, result2 := Return2(noReturn2Error()).ExitMsg("there is not error")
	assert.Equal(t, result1, "some string")
	assert.Equal(t, result2, 9)
}
