package reporting

import (
	"github.com/asobti/kube-monkey/chaos"
)

type ReportEntry struct {
	Result string
	Kind   string
	Name   string
	Err    string
}

type Report struct {
	Entries []*ReportEntry
}

func NewReport() *Report {
	r := Report{make([]*ReportEntry, 0, 10)}
	return &r
}

func (r *Report) AddEntry(result *chaos.Result) {
	report := &ReportEntry{}
	if result.Error() != nil {
		report.Result = "FAIL"
		report.Err = result.Error().Error()
	} else {
		report.Result = "OK"
		report.Err = ""
	}
	report.Kind = result.Victim().Kind()
	report.Name = result.Victim().Name()

	r.Entries = append(r.Entries, report)
}

func (r *Report) Len() int {
	return len(r.Entries)
}
