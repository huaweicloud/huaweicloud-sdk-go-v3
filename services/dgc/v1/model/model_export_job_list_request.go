package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ExportJobListRequest struct {
	Body *ExportJobsReq `json:"body,omitempty"`
}

func (o ExportJobListRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExportJobListRequest struct{}"
	}

	return strings.Join([]string{"ExportJobListRequest", string(data)}, " ")
}
