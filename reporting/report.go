package reporting

import (
	"time"

	"github.com/asobti/kube-monkey/chaos"
)

type ReportEntry struct {
	ReportTime time.Time
	Result     string
	Kind       string
	Name       string
	Err        string
}

type Report struct {
	ReportAt time.Time      `json:"schedule_built"`
	Results  []*ReportEntry `json:"results"`
}

func NewReport() *Report {
	r := Report{time.Now(), make([]*ReportEntry, 0, 10)}
	return &r
}

func (r *Report) AddEntry(resultTime time.Time, result *chaos.Result) {
	report := &ReportEntry{}
	report.ReportTime = resultTime
	if result.Error() != nil {
		report.Result = "FAIL"
		report.Err = result.Error().Error()
	} else {
		report.Result = "OK"
		report.Err = ""
	}
	report.Kind = result.Victim().Kind()
	report.Name = result.Victim().Name()

	r.Results = append(r.Results, report)
}

func (r *Report) Len() int {
	return len(r.Results)
}
