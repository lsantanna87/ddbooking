package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CommandTestSuite struct {
	suite.Suite
	t *testing.T
}

func (c *CommandTestSuite) SetupSuite() {
	c.t = c.T()
}

func TestCommandSuite(t *testing.T) {
	suite.Run(t, new(CommandTestSuite))
}

func (c *CommandTestSuite) TestShouldReturnAValidCommandForImport() {
	command := CreateImportCMD()

	assert.NotNil(c.t, command)
	assert.Equal(c.t, command.Name, "import")
	assert.Equal(c.t, command.Usage, "Import Events")
	assert.NotNil(c.t, command.Action)
}

func (c *CommandTestSuite) TestShouldReturnAValidCommandForValidate() {
	validate := CreateValidateCMD()

	assert.NotNil(c.t, validate)
	assert.Equal(c.t, validate.Name, "validate")
	assert.Equal(c.t, validate.Usage, "Validate if events are valid")
	assert.NotNil(c.t, validate.Action)
}

func (c *CommandTestSuite) TestShouldReturnCommandsWhenCreateCommand() {
	commands := CreateCommands(CreateImportCMD, CreateValidateCMD)
	assert.Len(c.t, commands, 2)
}

func (c *CommandTestSuite) TestShouldReturnNoCommandWhenCommandsListIsEmpty() {
	commands := CreateCommands()
	assert.Len(c.t, commands, 0)
}
