package test_commands

import (
	"testing"

	"github.com/winstarshl/pip-services3-commons-go-vgo/commands"
	"github.com/winstarshl/pip-services3-commons-go-vgo/run"
	"github.com/stretchr/testify/assert"
)

type TestListener struct{}

func (c *TestListener) OnEvent(correlationId string, e commands.IEvent, value *run.Parameters) {
	if correlationId == "wrongId" {
		panic("Test error")
	}
}

func TestGetEventName(t *testing.T) {
	event := commands.NewEvent("name")

	assert.NotNil(t, event)
	assert.Equal(t, "name", event.Name())
}

func TestEventNotify(t *testing.T) {
	event := commands.NewEvent("name")

	listener := &TestListener{}
	event.AddListener(listener)
	assert.Equal(t, 1, len(event.Listeners()))

	event.Notify("", nil)

	event.RemoveListener(listener)
	assert.Equal(t, 0, len(event.Listeners()))
}
