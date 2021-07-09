package steps

import (
	"code.vegaprotocol.io/data-node/integration/stubs"
)

func TimeIsUpdatedTo(timeService *stubs.TimeStub, newTime string) error {
	t, err := Time(newTime)
	panicW("date", err)
	timeService.SetTime(t)
	return nil
}
