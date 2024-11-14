//go:build acceptance

package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/stretchr/testify/assert"
)

// nolint: unused
type testContext struct {
	err      error
	cmdInput struct {
		parameters string
	}
	cmdResult struct {
		Output       string
		Err          error
		ProcessState *os.ProcessState
	}
}

// Errorf is used by the called assertion to report an error and is required to
// make testify assertions work
func (c *testContext) Errorf(format string, args ...interface{}) {
	c.err = fmt.Errorf(format, args...)
}

func (c *testContext) theAppRunsWithoutParameters() error {
	cmd := exec.Command("./testapp")
	output, err := cmd.CombinedOutput()
	c.cmdResult.Output = string(output)
	c.cmdResult.Err = err
	c.cmdResult.ProcessState = cmd.ProcessState

	return nil
}

func (c *testContext) theAppRunsWithParameters(args string) error {
	c.cmdInput.parameters = args
	cmdArgs := strings.Split(args, " ")
	cmd := exec.Command("./testapp", cmdArgs...)
	output, err := cmd.CombinedOutput()
	c.cmdResult.Output = string(output)
	c.cmdResult.Err = err
	c.cmdResult.ProcessState = cmd.ProcessState

	return nil
}

func (c *testContext) theAppExitsWithoutError() error {
	assert.NoError(c, c.cmdResult.Err)
	return c.err
}

func (c *testContext) theAppExitsWithAnError() error {
	assert.Error(c, c.cmdResult.Err)
	return c.err
}

func (c *testContext) theAppOutputContains(expected string) error {
	expected = strings.ReplaceAll(expected, "\\\"", "\"")
	assert.Contains(c, c.cmdResult.Output, expected)
	return c.err
}

func (c *testContext) theAppOutputDoesNotContain(unexpected string) error {
	unexpected = strings.ReplaceAll(unexpected, "\\\"", "\"")
	assert.NotContains(c, c.cmdResult.Output, unexpected)
	return c.err
}

func (c *testContext) theAppOutputHasLines(expected int) error {
	lines := strings.Split(c.cmdResult.Output, "\n")
	// Remove the empty line at the end
	lines = lines[:len(lines)-1]

	assert.Len(c, lines, expected)
	return c.err
}
