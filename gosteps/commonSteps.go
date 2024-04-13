package gosteps

import (
	"context"

	"github.com/rdumont/assistdog"
)

var (
	intVal int
)

func getAssistDogInstance(ctx context.Context) (context.Context, *assistdog.Assist) {
	var assist *assistdog.Assist
	assist, found := ctx.Value(ContextKeyAssistDogInstance).(*assistdog.Assist)
	if !found {
		assist := assistdog.NewDefault()
		//We have an integer value in table, so add its parser in assistDog instance
		assist.RegisterParser(intVal, intParser)
		//We also need to add a custom comparer, to compare values other than string
		assist.RegisterComparer(intVal, intComparer)
		//We can similarly add parsers and comparers for float, time.Time and all other types
		ctx = context.WithValue(ctx, ContextKeyAssistDogInstance, assist)
		return ctx, assist
	}
	return ctx, assist
}
