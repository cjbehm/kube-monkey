package scheduleService

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/asobti/kube-monkey/chaos"
	"github.com/asobti/kube-monkey/schedule"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/stretchr/testify/assert"
)

func newSchedule() *schedule.Schedule {
	return &schedule.Schedule{}
}

func TestGetScheduleEmpty(t *testing.T) {
	var sched1 string
	s := newSchedule()

	svc := scheduleService{sync.RWMutex{}, nil}
	sched1, _ = svc.GetSchedule()
	assert.Equal(t, "{}", sched1)

	svc.ReplaceSchedule(s)
	sched1, _ = svc.GetSchedule()
	assert.Equal(t, "{\"victims\":[]}", sched1)
}

func TestGetSchedule(t *testing.T) {
	var sched1 string
	now := time.Now()
	s := newSchedule()
	e1 := chaos.NewMock(&now)
	s.Add(e1)

	svc := scheduleService{sync.RWMutex{}, s}
	sched1, _ = svc.GetSchedule()
	assert.Equal(
		t,
		fmt.Sprintf("{\"victims\":[{\"kind\":\"Pod\",\"namespace\":\"default\",\"name\":\"name\",\"killat\":\"%s\"}]}", now.Format(schedule.DateFormat)),
		sched1,
	)
}

func TestReplaceSchedule(t *testing.T) {
	s := newSchedule()
	e1 := chaos.NewMock(nil)
	s.Add(e1)

	svc := scheduleService{sync.RWMutex{}, nil}
	sched1, _ := svc.GetSchedule()
	assert.Equal(t, "{}", sched1)
	svc.ReplaceSchedule(s)

	sched2, _ := svc.GetSchedule()
	assert.NotEqual(t, sched1, sched2)
	time.Sleep(1 * time.Second)

	s2 := newSchedule()
	e2 := chaos.NewMock(nil)
	s2.Add(e2)
	svc.ReplaceSchedule(s2)
	sched3, _ := svc.GetSchedule()
	assert.NotEqual(t, sched2, sched3)
}

func TestEndpointWithNoSchedule(t *testing.T) {
	sched := newSchedule()
	svc := &scheduleService{sync.RWMutex{}, sched}
	eps := makeScheduleEndpoint(svc)
	mux := http.NewServeMux()
	mux.Handle("/schedule", httptransport.NewServer(
		eps,
		func(_ context.Context, r *http.Request) (interface{}, error) { return nil, nil },
		encodeResponse,
	))
	srv := httptest.NewServer(mux)
	defer srv.Close()

	{
		want := `{"schedule":{"victims":[]},"asof":`
		req, _ := http.NewRequest("GET", srv.URL+"/schedule", strings.NewReader(""))
		resp, _ := http.DefaultClient.Do(req)
		body, _ := ioutil.ReadAll(resp.Body)
		assert.Contains(t, string(body), want)
	}
}

func TestEndpointWithSchedule(t *testing.T) {
	sched := newSchedule()
	svc := &scheduleService{sync.RWMutex{}, sched}
	eps := makeScheduleEndpoint(svc)
	mux := http.NewServeMux()
	mux.Handle("/schedule", httptransport.NewServer(
		eps,
		func(_ context.Context, r *http.Request) (interface{}, error) { return nil, nil },
		encodeResponse,
	))
	srv := httptest.NewServer(mux)
	defer srv.Close()
	generatedTime := time.Now()
	e := chaos.NewMock(&generatedTime)
	sched.Add(e)
	{
		want := fmt.Sprintf(`{"schedule":{"victims":[{"kind":"Pod","namespace":"default","name":"name","killat":"%s"}]},"asof":`, generatedTime.Format(schedule.DateFormat))
		req, _ := http.NewRequest("GET", srv.URL+"/schedule", strings.NewReader(""))
		resp, _ := http.DefaultClient.Do(req)
		body, _ := ioutil.ReadAll(resp.Body)
		assert.Contains(t, string(body), want)
	}
}
