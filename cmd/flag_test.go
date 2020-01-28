package cmd

import (
	"io/ioutil"
	"testing"

	"github.com/lsantanna87/ddbooking/pkg/service"
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

func (c *FlagTestSuite) TestShouldReturnValidFlagForFile() {
	flag := CreateFileFlag()

	assert.NotNil(c.t, flag)
	assert.Equal(c.t, "file", flag.Names()[0])
}

func (c *FlagTestSuite) TestShouldReturnValidFlagForText() {
	flag := CreateTextFlag()

	assert.NotNil(c.t, flag)
	assert.Equal(c.t, "text", flag.Names()[0])
}

func (c *FlagTestSuite) TestShouldReturnAllFlagsCreated() {
	flags := CreateFlags(CreateTextFlag, CreateTextFlag)

	assert.NotNil(c.t, flags)
	assert.Len(c.t, flags, 2)
}

func (c *FlagTestSuite) TestShouldReturnEmptyWhenNoFlagProvided() {
	flags := CreateFlags()

	assert.Nil(c.t, flags)
	assert.Len(c.t, flags, 0)
}

func (c *FlagTestSuite) TestShouldReturnValidEventListWhenProcessingFile() {
	events, _ := processFile("./fixture/events.json")

	valid, err := service.EventService{}.IsEventsValid(events)

	assert.NoError(c.t, err)
	assert.True(c.t, valid)
}

func (c *FlagTestSuite) TestShouldReturnErrorWhenPassingNilToProcessFile() {
	events, _ := processFile("")

	valid, err := service.EventService{}.IsEventsValid(events)

	assert.Error(c.t, err)
	assert.False(c.t, valid)
}

func (c *FlagTestSuite) TestShouldReturnValidEventListWhenProcessingText() {
	dat, _ := ioutil.ReadFile("./fixture/events.json")
	events, _ := processText(string(dat))

	valid, err := service.EventService{}.IsEventsValid(events)

	assert.NoError(c.t, err)
	assert.True(c.t, valid)
}

func (c *FlagTestSuite) TestShouldReturnErrorWhenPassingNilToProcessText() {
	events, _ := processFile("")

	valid, err := service.EventService{}.IsEventsValid(events)

	assert.Error(c.t, err)
	assert.False(c.t, valid)
}

func (t *FlagTestSuite) TestShouldReturnErrorWhenCommandImportWithNonExistentFlag() {
	c := CreateFakeContextWithFlag("teste")

	err := commandImport(c)

	assert.Error(t.t, err)
}

func (t *FlagTestSuite) TestShouldNotReturnErrorWhenCommandImportWithValidFlag() {
	c := CreateFakeContextWithFlag("text")

	err := commandImport(c)

	assert.Error(t.t, err)
}
