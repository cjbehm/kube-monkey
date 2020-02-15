package scheduleService

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/asobti/kube-monkey/schedule"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type ScheduleService interface {
	GetSchedule() (string, error)
	ReplaceSchedule(*schedule.Schedule)
}

func New() ScheduleService {
	return &scheduleService{sync.RWMutex{}, nil}
}

func ServeSchedule(s ScheduleService, port int) {
	scheduleHandler := httptransport.NewServer(
		makeUppercaseEndpoint(s),
		func(_ context.Context, r *http.Request) (interface{}, error) { return nil, nil },
		encodeResponse,
	)

	http.Handle("/schedule", scheduleHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

// --------------------------------------------

type getScheduleResponse struct {
	Schedule json.RawMessage `json:"schedule"`
	Err      string          `json:"err,omitempty"`
}

func makeUppercaseEndpoint(svc ScheduleService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		v, err := svc.GetSchedule()
		if err != nil {
			return getScheduleResponse{json.RawMessage(v), err.Error()}, nil
		}
		return getScheduleResponse{json.RawMessage(v), ""}, nil
	}
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type scheduleService struct {
	l sync.RWMutex
	s *schedule.Schedule
}

func (service *scheduleService) GetSchedule() (string, error) {
	service.l.RLock()
	defer service.l.RUnlock()
	if service.s == nil {
		return "{}", nil
	}

	bytes, err := service.s.MarshalJSON()

	return string(bytes), err
}

func (service *scheduleService) ReplaceSchedule(s *schedule.Schedule) {
	service.l.Lock()
	defer service.l.Unlock()
	service.s = s
}
