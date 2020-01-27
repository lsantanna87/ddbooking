package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FlagTestSuite struct {
	suite.Suite
	t *testing.T
}

func (c *FlagTestSuite) SetupSuite() {
	c.t = c.T()
}

func TestFlagSuite(t *testing.T) {
	suite.Run(t, new(FlagTestSuite))
}

func (c *FlagTestSuite) TestShouldReturnAValidFlagForFile() {
	flag := CreateFileFlag()

	assert.NotNil(c.t, flag)
	assert.Equal(c.t, "file", flag.Names()[0])
}

func (c *FlagTestSuite) TestShouldReturnAValidFlagForText() {
	flag := CreateTextFlag()

	assert.NotNil(c.t, flag)
	assert.Equal(c.t, "text", flag.Names()[0])
}

func (c *FlagTestSuite) TestShouldReturnAllFlagsCreated() {
	flags := CreateFlags(CreateTextFlag, CreateTextFlag)

	assert.NotNil(c.t, flags)
	assert.Len(c.t, flags, 2)
}

func (c *FlagTestSuite) TestShouldReturnEmptyWhenNoFlafProvided() {
	flags := CreateFlags()

	assert.Nil(c.t, flags)
	assert.Len(c.t, flags, 0)
}
