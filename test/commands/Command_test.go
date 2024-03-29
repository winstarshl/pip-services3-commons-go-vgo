package test_commands

import (
	"testing"

	"github.com/winstarshl/pip-services3-commons-go-vgo/commands"
	"github.com/winstarshl/pip-services3-commons-go-vgo/run"
	"github.com/stretchr/testify/assert"
)

func commandExec(correlationId string, args *run.Parameters) (interface{}, error) {
	if correlationId == "wrongId" {
		panic("Test error")
	}

	return nil, nil
}

func TestGetCommandName(t *testing.T) {
	command := commands.NewCommand("name", nil, commandExec)

	// Check match by individual fields
	assert.NotNil(t, command)
	assert.Equal(t, "name", command.Name())
}

func TestExecuteCommand(t *testing.T) {
	command := commands.NewCommand("name", nil, commandExec)

	_, err := command.Execute("", nil)
	assert.Nil(t, err)

	_, err = command.Execute("wrongId", nil)
	assert.NotNil(t, err)
}
