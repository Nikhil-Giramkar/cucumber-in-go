package tests

import (
	"os"
	"testing"

	gosteps "github.com/Nikhil-Giramkar/cucumber-in-go/gosteps"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

var opts = godog.Options{Output: colors.Colored(os.Stdout)}

func TestData(t *testing.T) {
	o := opts
	o.TestingT = t
	o.Tags = "goimpl"
	o.Format = "pretty"
	o.Paths = []string{"Features\\Completed\\Mathematics"}
	status := godog.TestSuite{
		Name:    "MyMath",
		Options: &o,
		ScenarioInitializer: func(sc *godog.ScenarioContext) {
			InitializeScenario(t, sc)
		},
	}.Run()

	if status != 0 {
		t.Fatalf("0 status expected, %d received", status)
	}
}

func InitializeScenario(t *testing.T, ctx *godog.ScenarioContext) {
	var textContext = gosteps.NewTestContext(t)
	gosteps.InitializeScenario(ctx, textContext)
}
