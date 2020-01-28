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

func (f *FlagTestSuite) SetupSuite() {
	f.t = f.T()
}

func TestFlagSuite(t *testing.T) {
	suite.Run(t, new(FlagTestSuite))
}

func (f *FlagTestSuite) TestShouldReturnValidFlagForFile() {
	flag := createFileFlag()

	assert.NotNil(f.t, flag)
	assert.Equal(f.t, "file", flag.Names()[0])
}

func (f *FlagTestSuite) TestShouldReturnValidFlagForText() {
	flag := createTextFlag()

	assert.NotNil(f.t, flag)
	assert.Equal(f.t, "text", flag.Names()[0])
}

func (f *FlagTestSuite) TestShouldReturnAllFlagsCreated() {
	flags := createFlags(createTextFlag, createTextFlag)

	assert.NotNil(f.t, flags)
	assert.Len(f.t, flags, 2)
}

func (f *FlagTestSuite) TestShouldReturnEmptyWhenNoFlagProvided() {
	flags := createFlags()

	assert.Nil(f.t, flags)
	assert.Len(f.t, flags, 0)
}

func (c *FlagTestSuite) TestShouldReturnValidEventListWhenProcessingFile() {
	events, _ := processFile("./fixture/events.json")

	valid, err := service.EventService{}.IsEventsValid(events)

	assert.NoError(c.t, err)
	assert.True(c.t, valid)
}

func (f *FlagTestSuite) TestShouldReturnErrorWhenPassingNilToProcessFile() {
	events, _ := processFile("")

	valid, err := service.EventService{}.IsEventsValid(events)

	assert.Error(f.t, err)
	assert.False(f.t, valid)
}

func (f *FlagTestSuite) TestShouldReturnValidEventListWhenProcessingText() {
	dat, _ := ioutil.ReadFile("./fixture/events.json")
	events, _ := processText(string(dat))

	valid, err := service.EventService{}.IsEventsValid(events)

	assert.NoError(f.t, err)
	assert.True(f.t, valid)
}

func (f *FlagTestSuite) TestShouldReturnErrorWhenPassingNilToProcessText() {
	events, _ := processText("")

	valid, err := service.EventService{}.IsEventsValid(events)

	assert.Error(f.t, err)
	assert.False(f.t, valid)
}
