package reporting

import (
	"encoding/json"
	"testing"

	"github.com/asobti/kube-monkey/chaos"
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

	assert.Empty(t, e.Entries)
}

func TestAddResult(t *testing.T) {
	e := NewReport()
	c := chaos.NewMock(nil)
	r := c.NewResult(nil)

	e.AddEntry(r)
	assert.Equal(t, e.Len(), 1)
}

func TestEncodeJSON(t *testing.T) {
	e := NewReport()
	c1 := chaos.NewMock(nil)
	r1 := c1.NewResult(nil)
	c2 := chaos.NewMock(nil)
	r2 := c2.NewResult(mockError{"Fizzbut"})

	e.AddEntry(r1)
	e.AddEntry(r2)

	bytes, err := json.Marshal(e)
	assert.Nil(t, err)
	assert.Equal(t, string(bytes), `{"Entries":[{"Result":"OK","Kind":"Pod","Name":"name","Err":""},{"Result":"FAIL","Kind":"Pod","Name":"name","Err":"Fizzbut"}]}`)
}
