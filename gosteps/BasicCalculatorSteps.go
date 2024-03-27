package gosteps

import (
	"context"
	"strconv"

	backend "github.com/Nikhil-Giramkar/cucumber-in-go/backend"
	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
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
	assist := assistdog.NewDefault()

	//To remove the first row that contains heading of table
	table = &godog.Table{
		Rows: table.Rows[1:len(table.Rows)],
	}

	operation := ctx.Value(ContextKeyOperationSelected).(string)
	firstNum := ctx.Value(ContextKeyFirstNumber).(int)
	secondNum := ctx.Value(ContextKeySecondNumber).(int)

	backendResult := backend.Calculate(firstNum, secondNum, operation)
	actualRes := &ActualResult{
		Answer:  strconv.Itoa(backendResult.Answer),
		Message: backendResult.Message,
	}

	assist.CompareToInstance(actualRes, table)
}
