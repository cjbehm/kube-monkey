package reportingService

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/asobti/kube-monkey/reporting"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type ReportService interface {
	AddReport(r *reporting.Report)
	GetReports() []*reporting.Report
}

func New() ReportService {
	svc := &reportService{reportLock: sync.RWMutex{}, reports: make([]*reporting.Report, 0, 10)}
	return svc
}

func ServeReports(svc ReportService) {
	reportHandler := httptransport.NewServer(
		makeReportEndpoint(svc),
		func(_ context.Context, r *http.Request) (interface{}, error) { return nil, nil },
		encodeResponse,
	)

	http.Handle("/report", reportHandler)
}

// --------------------------------------------

type reportService struct {
	reportLock sync.RWMutex
	reports    []*reporting.Report
}

type getReportResponse struct {
	Report json.RawMessage `json:"report"`
	Asof   time.Time       `json:"asof"`
	Err    string          `json:"err,omitempty"`
}

func makeReportEndpoint(svc ReportService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		// v, err := svc.GetSchedule()
		// if err != nil {
		// 	return getReportResponse{json.RawMessage(v), time.Now(), err.Error()}, nil
		// }
		//return getReportResponse{json.RawMessage(`{}`), time.Now(), ""}, nil
		bytes, _ := json.Marshal(svc.GetReports())
		return getReportResponse{bytes, time.Now(), ""}, nil
	}
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func (svc *reportService) AddReport(r *reporting.Report) {
	svc.reportLock.Lock()
	defer svc.reportLock.Unlock()
	svc.reports = append(svc.reports, r)
}

func (svc *reportService) GetReports() []*reporting.Report {
	return svc.reports
}
