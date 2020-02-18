package reporting

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/asobti/kube-monkey/chaos"
	"github.com/bouk/monkey"
	"github.com/stretchr/testify/assert"
)

type mockError struct {
	mockError string
}

func (m mockError) Error() string {
	return m.mockError
}

func TestEmptyResults(t *testing.T) {
	e := NewReport()

	assert.Empty(t, e.Results)
}

func TestAddResult(t *testing.T) {
	e := NewReport()
	c := chaos.NewMock(nil)
	r := c.NewResult(nil)

	e.AddEntry(time.Now(), r)
	assert.Equal(t, e.Len(), 1)
}

func TestEncodeJSON(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2018, 4, 16, 12, 0, 0, 0, time.UTC)
	})
	defer monkey.Unpatch(time.Now)
	e := NewReport()
	c1 := chaos.NewMock(nil)
	r1 := c1.NewResult(nil)
	c2 := chaos.NewMock(nil)
	r2 := c2.NewResult(mockError{"Fizzbut"})

	e.AddEntry(time.Now(), r1)
	e.AddEntry(time.Now(), r2)

	bytes, err := json.Marshal(e)
	assert.Nil(t, err)
	assert.Equal(t, `{"schedule_built":"2018-04-16T12:00:00Z","results":[{"ReportTime":"2018-04-16T12:00:00Z","Result":"OK","Kind":"Pod","Name":"name","Err":""},{"ReportTime":"2018-04-16T12:00:00Z","Result":"FAIL","Kind":"Pod","Name":"name","Err":"Fizzbut"}]}`, string(bytes))
}
