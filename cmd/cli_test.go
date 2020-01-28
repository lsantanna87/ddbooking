package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CLITestSuite struct {
	suite.Suite
	t *testing.T
}

func (c *CLITestSuite) SetupSuite() {
	c.t = c.T()
}

func TestCLISuite(t *testing.T) {
	suite.Run(t, new(CLITestSuite))
}

func (c *CLITestSuite) TestShouldReturnErrorWhenContextIsNil() {
	err := validateCLI(nil)

	assert.Error(c.t, err)
}

func (c *CLITestSuite) TestShouldReturnErrorWhenMoreThanOneFlagPassedAsArgument() {
	context := CreateFakeContextWithTwoFlags("test", "test2", "--two")
	err := validateCLI(context)

	assert.Error(c.t, err)
}
