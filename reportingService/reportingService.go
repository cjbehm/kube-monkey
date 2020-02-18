package reportingService

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/asobti/kube-monkey/reporting"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type ReportService interface {
	AddReport(r *reporting.Report)
}

func ServeReport(r ReportService) {
	reportHandler := httptransport.NewServer(
		makeReportEndpoint(r),
		func(_ context.Context, r *http.Request) (interface{}, error) { return nil, nil },
		encodeResponse,
	)

	http.Handle("/report", reportHandler)
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

// --------------------------------------------

type getReportResponse struct {
	Schedule json.RawMessage `json:"report"`
	Asof     time.Time       `json:"asof"`
	Err      string          `json:"err,omitempty"`
}

func makeReportEndpoint(r ReportService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		// v, err := svc.GetSchedule()
		// if err != nil {
		// 	return getReportResponse{json.RawMessage(v), time.Now(), err.Error()}, nil
		// }
		return getReportResponse{json.RawMessage(`{}`), time.Now(), ""}, nil
	}
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
