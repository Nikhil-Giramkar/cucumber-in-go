package gosteps

import (
	"testing"

	"github.com/cucumber/godog"
)

type TestContext struct {
	Testing *testing.T
}

func NewTestContext(t *testing.T) TestContext {
	return TestContext{
		Testing: t,
	}
}

func InitializeScenario(ctx *godog.ScenarioContext, context TestContext) {
	ctx.Step(`^I click on digit (\d+)$`, context.iClickOnDigit)
	ctx.Step(`^i click on second digit (\d+)$`, context.iClickOnSecondDigit)
	ctx.Step(`^I click on (\w+) button$`, context.iClickOnOperation)
	ctx.Step(`^calculation results are$`, context.calculationResultsAre)
}
