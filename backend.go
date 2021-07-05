package haproxy

import (
	"encoding/csv"
	"fmt"

	"github.com/gocarina/gocsv"
)

// Response from HAProxy "show backend" command.
type ShowBackendResponse struct {
	Name string `csv:"# name"`
}

// Equivalent to HAProxy "show backend" command.
func ShowBackend(h HAProxy) (showBackendResponses []*ShowBackendResponse, err error) {
	res, err := h.RunCommand("show backend")
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(res)
	err = gocsv.UnmarshalCSV(reader, &showBackendResponses)
	if err != nil {
		return nil, fmt.Errorf("error reading csv: %s", err)
	}

	return showBackendResponses, nil
}
