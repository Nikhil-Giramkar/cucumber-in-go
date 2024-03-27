package gosteps

import (
	"context"
	"strconv"

	backend "github.com/Nikhil-Giramkar/cucumber-in-go/backend"
	"github.com/cucumber/godog"
	"gotest.tools/assert"
)

func (t TestContext) iClickOnDigit(ctx context.Context, firstNum int) (context.Context, error) {
	ctx = context.WithValue(ctx, ContextKeyFirstNumber, firstNum)
	return ctx, nil
}

func (t TestContext) iClickOnSecondDigit(ctx context.Context, secondNum int) (context.Context, error) {
	ctx = context.WithValue(ctx, ContextKeySecondNumber, secondNum)
	return ctx, nil
}

func (t TestContext) iClickOnOperation(ctx context.Context, operation string) (context.Context, error) {
	ctx = context.WithValue(ctx, ContextKeyOperationSelected, operation)
	return ctx, nil
}

func (t TestContext) calculationResultsAre(ctx context.Context, table *godog.Table) {
	operation := ctx.Value(ContextKeyOperationSelected).(string)
	firstNum := ctx.Value(ContextKeyFirstNumber).(int)
	secondNum := ctx.Value(ContextKeySecondNumber).(int)

	actualResult := backend.Calculate(firstNum, secondNum, operation)
	expectedResult, _ := strconv.ParseInt(table.Rows[1].Cells[1].Value, 10, 32)
	assert.Check(t.Testing, actualResult.Answer == int(expectedResult), "Error")
}
