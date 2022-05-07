package component_tests

import "github.com/cucumber/godog"

func iEat(arg1 int) error {
	return nil
	// return godog.ErrPending
}

func thereAreGodogs(arg1 int) error {
	return nil
	// return godog.ErrPending
}

func thereShouldBeRemaining(arg1 int) error {
	return nil
	// return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I eat (\d+)$`, iEat)
	ctx.Step(`^there are (\d+) godogs$`, thereAreGodogs)
	ctx.Step(`^there should be (\d+) remaining$`, thereShouldBeRemaining)
}
