//go:build acceptance

package main_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/cucumber/godog"
)

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format: "pretty",
			Paths:  []string{"features"},
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

// nolint: unused
func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {})
}

// nolint: unused
func InitializeScenario(ctx *godog.ScenarioContext) {
	tc := testContext{}
	ctx.Before(func(ctx context.Context, scn *godog.Scenario) (context.Context, error) {
		return ctx, nil
	})
	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		if err != nil {
			fmt.Printf(
				"Command line output for test\nUsing parameters: %s\n%s",
				tc.cmdInput.parameters,
				tc.cmdResult.Output,
			)
		}

		return ctx, nil
	})
	ctx.Step(`^the app runs with parameters "(.*)"$`, tc.theAppRunsWithParameters)
	ctx.Step(`^the app exits with an error$`, tc.theAppExitsWithAnError)
	ctx.Step(`^the app output contains "(.*)"$`, tc.theAppOutputContains)
	ctx.Step(`^the app output does not contain "(.*)"$`, tc.theAppOutputDoesNotContain)
}
